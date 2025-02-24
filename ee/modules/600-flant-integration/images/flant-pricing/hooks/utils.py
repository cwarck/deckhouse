# Copyright 2023 Flant JSC
# Licensed under the Deckhouse Platform Enterprise Edition (EE) license.
# See https://github.com/deckhouse/deckhouse/blob/main/ee/LICENSE
#
# util functions for hook.

import ssl
import json
import os
from typing import Dict, List, Any
from urllib3 import disable_warnings
from requests.adapters import Retry, PoolManager, ResponseError

disable_warnings()

BACKOFF_LIMIT = 3

PROMETHEUS_URL = "https://prometheus.d8-monitoring:9090/api/v1/query"


try:
    with open("/var/run/secrets/kubernetes.io/serviceaccount/token", encoding="utf-8") as f:
        SERVICE_ACCOUNT_TOKEN = f.read()
except FileNotFoundError:
    SERVICE_ACCOUNT_TOKEN = os.getenv("SERVICE_ACCOUNT_TOKEN")


def prometheus_query(query: str) -> List[Dict[str, Any]]:
    """query prometheus from query and get api result response"""

    response = do_get_request(
        url=PROMETHEUS_URL,
        params={"query": query},
        headers={"Authorization": f"Bearer {SERVICE_ACCOUNT_TOKEN}"},
        decode_json=True,
    )
    # response:
    # {"status": "success", "data": {"resultType": "vector", "result": [{"metric": {"__name__": "name"}, "value": ["timestamp", "value"]}]}}
    if response["status"] != "success":
        raise ResponseError(f"error quering prometheus with query '{query}'")

    return response.get("data", {}).get("result", [0])


def prometheus_query_value(query: str) -> float:
    """query prometheus from query and get only value result"""

    query_result = prometheus_query(query)
    if len(query_result) > 0:
        return float(query_result[0].get("value", [0, 0])[1])
    return 0.0


def do_get_request(
        url: str,
        params: Dict[str, str] = None,
        headers: Dict[str, str] = None,
        decode_json: bool = False,
) -> Any:
    """do http request with retries"""

    retries = Retry(
        total=BACKOFF_LIMIT,
        backoff_factor=0.1,
        status_forcelist=[500, 502, 503, 504],
    )
    pool = PoolManager(retries=retries, cert_reqs=ssl.CERT_NONE)
    response = pool.request("GET", url, fields=params, headers=headers)

    if response.status != 200:
        raise ResponseError(
            f"{response.status} status code: {response.data.decode('utf-8')}"
        )

    if decode_json:
        return json.loads(response.data)

    return response.data
