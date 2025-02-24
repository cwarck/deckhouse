#!/bin/bash
{{- /*
# Copyright 2021 Flant JSC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
*/}}
shopt -s extglob

function ip_in_subnet(){
  python3 -c "import ipaddress; exit(0) if ipaddress.ip_address('$1') in ipaddress.ip_network('$2') else exit(1)"
  return $?
}

if ! metadata="$(curl -sH Metadata-Flavor:Google 169.254.169.254/computeMetadata/v1/instance/?recursive=true 2>/dev/null)"; then
  echo "Can't get network cidr from metadata"
  exit 1
fi

network_cidr=$(echo "$metadata" | jq -er '.attributes."node-network-cidr"')
if [ -z "$network_cidr" ]; then
  echo "network cidr is empty"
  exit 1
fi

ip_addr_show_output=$(ip -json addr show)
primary_mac="$(grep -m 1 -Po '(?<=macaddress: ).+' /etc/netplan/50-cloud-init.yaml)"
primary_ifname="$(echo "$ip_addr_show_output" | jq -re --arg mac "$primary_mac" '.[] | select(.address == $mac) | .ifname')"
for i in /sys/class/net/!($primary_ifname); do
  if ! udevadm info "$i" 2>/dev/null | grep -Po '(?<=E: ID_NET_DRIVER=)virtio_net.*' 1>/dev/null 2>&1; then
    continue
  fi

  ifname=$(basename "$i")
  mac="$(echo "$ip_addr_show_output" | jq -re --arg ifname "$ifname" '.[] | select(.ifname == $ifname) | .address')"

  ip="$(echo "$metadata" | jq --arg m "$mac" -er '.networkInterfaces[] | select(.mac == $m) | .ip')"
  route_settings=""
  if ip_in_subnet "$ip" "$network_cidr"; then
    read -r -d '' route_settings <<ROUTE_EOF
      routes:
      - to: $network_cidr
        scope: link
ROUTE_EOF
  fi

{{- /* # Configure the internal interface to route all vpc to all vm */}}
  cat > /etc/netplan/999-cim-"$ifname".yaml <<BOOTSTRAP_NETWORK_EOF
network:
  version: 2
  ethernets:
    $ifname:
      dhcp4: true
      dhcp4-overrides:
        use-hostname: false
        use-routes: false
      match:
        macaddress: $mac
      $route_settings
BOOTSTRAP_NETWORK_EOF
done
netplan generate
netplan apply
shopt -u extglob
