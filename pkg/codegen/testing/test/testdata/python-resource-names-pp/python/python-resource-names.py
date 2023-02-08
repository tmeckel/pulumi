import pulumi
import pulumi_azure_native as azure_native

cluster = azure_native.containerservice.ManagedCluster("cluster",
    agent_pool_profiles=[azure_native.containerservice.ManagedClusterAgentPoolProfileArgs(
        count=2,
        enable_fips=False,
        kubelet_disk_type="OS",
        max_pods=110,
        mode="System",
        name="type1",
        orchestrator_version="1.21.9",
        os_disk_size_gb=128,
        os_disk_type="Managed",
        os_sku="Ubuntu",
        os_type="Linux",
        type="VirtualMachineScaleSets",
        vm_size="Standard_B2ms",
        vnet_subnet_id="/subscriptions/0282681f-7a9e-424b-80b2-96babd57a8a1/resourceGroups/test-rga2bd359a/providers/Microsoft.Network/virtualNetworks/test-vnet4b80e99b/subnets/test-subnet",
    )],
    dns_prefix="dns",
    enable_rbac=True,
    kubernetes_version="1.21.9",
    location="eastus",
    network_profile={
        "dnsServiceIP": "10.10.0.10",
        "dockerBridgeCidr": "172.17.0.1/16",
        "loadBalancerProfile": {
            "effectiveOutboundIPs": [azure_native.containerservice.ResourceReferenceArgs(
                id="/subscriptions/0282681f-7a9e-424b-80b2-96babd57a8a1/resourceGroups/MC_test-rga2bd359a_test-aks5fb1e730_eastus/providers/Microsoft.Network/publicIPAddresses/2a2610b5-67f3-4aec-a277-a032b2364d70",
            )],
            "managedOutboundIPs": {
                "count": 1,
            },
        },
        "loadBalancerSku": "Standard",
        "networkPlugin": "azure",
        "outboundType": "loadBalancer",
        "serviceCidr": "10.10.0.0/16",
    },
    node_resource_group="MC_test-rga2bd359a_test-aks5fb1e730_eastus",
    resource_group_name="test-rga2bd359a",
    resource_name_="test-aks5fb1e730",
    service_principal_profile={
        "clientId": "64e3783f-3214-4ba7-bb52-12ad85412527",
    },
    sku={
        "name": "Basic",
        "tier": "Free",
    },
    opts=pulumi.ResourceOptions(protect=True))
