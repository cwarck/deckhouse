# [<en>] Section for bootstrapping the Deckhouse cluster.
# [<en>] https://deckhouse.io/documentation/v1/installing/configuration.html#initconfiguration
# [<ru>] Секция первичной инициализации кластера Deckhouse.
# [<ru>] https://deckhouse.ru/documentation/v1/installing/configuration.html#initconfiguration
apiVersion: deckhouse.io/v1
kind: InitConfiguration
deckhouse:
  releaseChannel: Stable
  # [<en>] The Minimal bundle is used when installing Deckhouse in an existing cluster.
  # [<ru>] При установке Deckhouse в существующий кластер используется вариант поставки — Minimal.
  bundle: Minimal
  configOverrides:
    global:
      modules:
        # [<en>] Template that will be used for system apps domains within the cluster.
        # [<en>] E.g., Grafana for %s.example.com will be available as 'grafana.example.com'.
        # [<en>] You can change it to your own or follow the steps in the guide and change it after installation.
        # [<ru>] Шаблон, который будет использоваться для составления адресов системных приложений в кластере.
        # [<ru>] Например, Grafana для %s.example.com будет доступна на домене 'grafana.example.com'.
        # [<ru>] Можете изменить на свой сразу, либо следовать шагам руководства и сменить его после установки.
        publicDomainTemplate: "%s.example.com"
        # [<ru>] If necessary, specify in the customTolerationKeys array
        # [<ru>] all the tains to which Deckhouse should have toleration.
        # [<ru>] The following is an example for the case if you need Deckhouse and its components to be able
        # [<ru>] to run on nodes that have taint SystemLoad.
        # [<ru>] При необходимости, укажите в массиве customTolerationKeys все taint'ы
        # [<ru>] к которым Deckhouse должен иметь toleration.
        # [<ru>] Далее приведен пример для случая, если нужно чтобы Deckhouse и его компоненты смогли запускаться
        # [<ru>] на узлах, имеющих taint SystemLoad
        # [<en>] You might consider changing this.
        # [<ru>] Возможно, захотите изменить.
        placement:
          customTolerationKeys:
          - SystemLoad
    certManagerEnabled: true
    # [<en>] Use for Deckhouse 1.46 or newer.
    # [<ru>] для Deckhouse 1.46 и новее.
    # documentationEnabled: true
    # [<en>] Use for Deckhouse 1.45 or older.
    # [<ru>] Для Deckhouse 1.45 и старее.
    deckhouseWebEnabled: true
