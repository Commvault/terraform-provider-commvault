---
page_title: " Commvault : commvault_kubernetes_appgroup Resource"
subcategory: "Kubernetes"
description: |-
    Use the commvault_kubernetes_appgroup resource type to create or delete kubernetes appgroup in the CommCell environment.

---

# commvault_kubernetes_appgroup (Resource)

Use the commvault_kubernetes_appgroup resource type to create or delete kubernetes appgroup in the CommCell environment.

## Example Usage

**Configure commvault kubernetes appgroup with required fields**
```hcl
data "commvault_client" "access_node1" {
  name = "client1"
}

data "commvault_plan" "plan1" {
  name = "AWS-Test-Plan"
}

resource "commvault_kubernetes_cluster" "kubernetes_cluster1" {
  name           = "SP32-Terraform-Test-Kubernetes"
  apiserver      = "https://1.2.3.4:6443"
  serviceaccount = "cvadmin"
  servicetoken   = "##############"
  accessnodes {
    id   = data.commvault_client.access_node1.id
    type = 3
  }
}

data "commvault_kubernetes_applications" "kubernetes_applications" {
  name      = "my-pod"
  clusterid =  commvault_kubernetes_cluster.kubernetes_cluster1.id
  namespace = "default"
}

data "commvault_kubernetes_labels" "kubernetes_labels" {
  name      = "modifierAt=12655"
  clusterid = commvault_kubernetes_cluster.kubernetes_cluster1.id
  namespace = "default"
}

data "commvault_kubernetes_namespaces" "kubernetes_namespaces" {
  name      = "1sts-volctemplate"
  clusterid = commvault_kubernetes_cluster.kubernetes_cluster1.id
}

data "commvault_kubernetes_storageclasses" "kubernetes_storageclasses" {
  name      = "rook-ceph-block"
  clusterid = commvault_kubernetes_cluster.kubernetes_cluster1.id
}

data "commvault_kubernetes_volumes" "kubernetes_volumes" {
  name      = "mysql-pvc"
  clusterid = commvault_kubernetes_cluster.kubernetes_cluster1.id
  namespace = "default"
}

resource "commvault_kubernetes_appgroup" "kubernetes_appgroup1" {
  name = "SP32-Terraform-Test-Kubernetes-APPGROUP"
  cluster {
    id = commvault_kubernetes_cluster.kubernetes_cluster1.id
  }
  plan {
    id = data.commvault_plan.plan1.id
  }
  content {
    labelselectors {
      selectorlevel = "Application"
      selectorvalue = "Test1=Value1"
    }
    labelselectors {
      selectorlevel = "Volumes"
      selectorvalue = "Test2=Value2"
    }
    labelselectors {
      selectorlevel = "Namespace"
      selectorvalue = "Test3=Value3"
    }
    applications {
      guid = data.commvault_kubernetes_namespaces.kubernetes_namespaces.id
      name = "1sts-volctemplate"
      type = "NAMESPACE"
    }
    applications {
      guid = data.commvault_kubernetes_applications.kubernetes_applications.id
      name = "my-pod"
      type = "APPLICATION"
    }
    applications {
      guid = "default`PersistentVolumeClaim`mysql-pvc`f5e7a010-cc52-4e4c-82aa-6656de0118ee"
      name = "mysql-pvc"
      type = "PVC"
    }
    applications {
      guid = "default`Label`modifierAt=12655"
      name = "modifierAt=12655"
      type = "APPLICATION"
    }
  }
}
```

**Configure commvault kubernetes appgroup with custom fields**
```hcl
data "commvault_client" "access_node1" {
  name = "bdcsrvtest05"
}

data "commvault_plan" "plan1" {
  name = "AWS-Test-Plan"
}

data "commvault_region"   "region1" {
  name = "Australia"
}

data "commvault_plan" "plan2" {
  name = "Demo Plan"
}

data "commvault_timezone" "timezone2" {
  name = "Singapore Standard Time"
}

resource "commvault_kubernetes_cluster" "kubernetes_cluster2" {
  name           = "SP32-Terraform-Test-Kubernetes-Custom-up"
  apiserver      = "https://1.2.3.4:6443"
  serviceaccount = "cvadmin"
  servicetoken   = "#########"
  accessnodes {
    id   = data.commvault_client.access_node1.id
    type = 3
  }
  accessnodes {
    id   = 3986
    type = 3
  }
  servicetype = "ONPREM"
  etcdprotection {
    plan {
      id = data.commvault_plan.plan1.id
    }
    enabled = "true"
  }
  activitycontrol {
    enablebackup             = "true"
    enablerestore            = "true"
  }
  region {
    id   = data.commvault_region.region1.id
  }
  tags {
    name  = "testK8s"
    value = "K8s"
  }
  tags {
    name  = "testK8s-2"
    value = "K8s-2"
  }
}

data "commvault_kubernetes_applications" "kubernetes_applications" {
  name      = "my-pod"
  clusterid =  commvault_kubernetes_cluster.kubernetes_cluster2.id
  namespace = "default"
}

data "commvault_kubernetes_labels" "kubernetes_labels" {
  name      = "modifierAt=12655"
  clusterid = commvault_kubernetes_cluster.kubernetes_cluster2.id
  namespace = "default"
}

data "commvault_kubernetes_namespaces" "kubernetes_namespaces" {
  name      = "1sts-volctemplate"
  clusterid = commvault_kubernetes_cluster.kubernetes_cluster2.id
}

data "commvault_kubernetes_storageclasses" "kubernetes_storageclasses" {
  name      = "rook-ceph-block"
  clusterid = commvault_kubernetes_cluster.kubernetes_cluster2.id
}

data "commvault_kubernetes_volumes" "kubernetes_volumes" {
  name      = "mysql-pvc"
  clusterid = commvault_kubernetes_cluster.kubernetes_cluster2.id
  namespace = "default"
}

resource "commvault_kubernetes_appgroup" "kubernetes_appgroup2" {
  name = "SP32-Terraform-Test-Kubernetes-APPGROUP-Custom-up"
  cluster {
    id = commvault_kubernetes_cluster.kubernetes_cluster2.id
  }
  plan {
    id = data.commvault_plan.plan2.id
  }
  content {
    labelselectors {
      selectorlevel = "Application"
      selectorvalue = "Test1=Value1"
    }
    applications {
      guid = data.commvault_kubernetes_namespaces.kubernetes_namespaces.id
      name = "1sts-volctemplate"
      type = "NAMESPACE"
    }
    applications {
      guid = data.commvault_kubernetes_applications.kubernetes_applications.id
      name = "my-pod"
      type = "APPLICATION"
    }
  }
  filters {
    skipstatelessapps = "false"
    labelselectors {
      selectorlevel = "Namespace"
      selectorvalue = "Test3=Value3"
    }
    applications {
      guid = "default`Label`modifierAt=12655"
      name = "modifierAt=12655"
      type = "APPLICATION"
    }
    applications {
      guid = "non-namespaced`Namespace`1sts-volctemplate`9f2497f6-f6bb-46d8-b8e2-7a015763f4e3"
      name = "1sts-volctemplate"
      type = "NAMESPACE"
    }
  }
  activitycontrol {
    enablebackup = "false"
  }
  timezone {
    id = data.commvault_timezone.timezone2.id
  }
  options {
    backupstreams = 60
    jobstarttime = 66540
  }
  tags {
    name = "testK8sGP"
    value = "K8sGP"
  }
  tags {
    name = "testK8sGP-2"
    value = "K8sGP-2"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required
- `name` (String) Specify new name to rename an Application Group
- `cluster` (Block List) (see [below for nested schema](#nestedblock--cluster))
- `plan` (Block List) (see [below for nested schema](#nestedblock--plan))
- `content` (Block List) Item describing the content for Application Group (see [below for nested schema](#nestedblock--content))

### Optional
- `activitycontrol` (Block List) (see [below for nested schema](#nestedblock--activitycontrol))
- `filters` (Block List) (see [below for nested schema](#nestedblock--filters))
- `options` (Block List) (see [below for nested schema](#nestedblock--options))
- `tags` (Block Set) (see [below for nested schema](#nestedblock--tags))
- `timezone` (Block List) (see [below for nested schema](#nestedblock--timezone))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--activitycontrol"></a>
### Nested Schema for `activitycontrol`

Optional:

- `enablebackup` (String)


<a id="nestedblock--cluster"></a>
### Nested Schema for `cluster`

Optional:

- `name` (String)

Read-Only:

- `id` (Number) The ID of this resource.


<a id="nestedblock--content"></a>
### Nested Schema for `content`

Optional:

- `applications` (Block Set) List of applications to be added as content (see [below for nested schema](#nestedblock--content--applications))
- `labelselectors` (Block Set) List of label selectors to be added as content (see [below for nested schema](#nestedblock--content--labelselectors))

<a id="nestedblock--content--applications"></a>
### Nested Schema for `content.applications`

Required:

- `guid` (String) GUID value of the Kubernetes Application to be associated as content
- `type` (String) Type of the Kubernetes application [NAMESPACE, APPLICATION, PVC, LABELS]

Optional:

- `name` (String) Name of the application


<a id="nestedblock--content--labelselectors"></a>
### Nested Schema for `content.labelselectors`

Required:

- `selectorlevel` (String) Selector level of the label selector [Application, Volumes, Namespace]
- `selectorvalue` (String) Value of the label selector in key=value format



<a id="nestedblock--filters"></a>
### Nested Schema for `filters`

Optional:

- `applications` (Block Set) List of applications to be added as content (see [below for nested schema](#nestedblock--filters--applications))
- `labelselectors` (Block Set) List of label selectors to be added as content (see [below for nested schema](#nestedblock--filters--labelselectors))
- `skipstatelessapps` (String) Specify whether to skip backup of stateless applications

<a id="nestedblock--filters--applications"></a>
### Nested Schema for `filters.applications`

Required:

- `guid` (String) GUID value of the Kubernetes Application to be associated as content
- `type` (String) Type of the Kubernetes application [NAMESPACE, APPLICATION, PVC, LABELS]

Optional:

- `name` (String) Name of the application


<a id="nestedblock--filters--labelselectors"></a>
### Nested Schema for `filters.labelselectors`

Required:

- `selectorlevel` (String) Selector level of the label selector [Application, Volumes, Namespace]
- `selectorvalue` (String) Value of the label selector in key=value format



<a id="nestedblock--options"></a>
### Nested Schema for `options`

Optional:

- `backupstreams` (Number) Define number of parallel data readers
- `cvnamespacescheduling` (String) Define setting to enable scheduling worker Pods to CV Namespace for CSI-Snapshot enabled backups
- `jobstarttime` (Number) Define the backup job start time in epochs
- `snapfallbacktolivevolumebackup` (String) Define setting to enable fallback to live volume backup in case of snap failure
- `workerresources` (Block List) (see [below for nested schema](#nestedblock--options--workerresources))

<a id="nestedblock--options--workerresources"></a>
### Nested Schema for `options.workerresources`

Optional:

- `cpulimits` (String) Define limits.cpu to set on the worker Pod
- `cpurequests` (String) Define requests.cpu to set on the worker Pod
- `memorylimits` (String) Define limits.memory to set on the worker Pod
- `memoryrequests` (String) Define requests.memory to set on the worker Pod



<a id="nestedblock--plan"></a>
### Nested Schema for `plan`

Optional:

- `name` (String)

Read-Only:

- `id` (Number) The ID of this resource.


<a id="nestedblock--tags"></a>
### Nested Schema for `tags`

Optional:

- `name` (String)
- `value` (String)


<a id="nestedblock--timezone"></a>
### Nested Schema for `timezone`

Optional:

- `name` (String)

Read-Only:

- `id` (Number) The ID of this resource.


