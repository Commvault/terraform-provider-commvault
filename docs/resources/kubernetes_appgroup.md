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

Each `commvault_kubernetes_*` data source requires a live cluster ID, so the cluster resource
must be created first. Data sources resolve Kubernetes object GUIDs by name at plan time.

```hcl
# Access node (MediaAgent) that communicates with the cluster API server
data "commvault_client" "access_node1" {
  name = "client1"
}

# Backup plan that defines the RPO and retention for this application group
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

# Resolves the GUID of a specific pod by name and namespace
data "commvault_kubernetes_applications" "kubernetes_applications" {
  name      = "my-pod"
  clusterid =  commvault_kubernetes_cluster.kubernetes_cluster1.id
  namespace = "default"
}

# Resolves the GUID of a Kubernetes label selector
data "commvault_kubernetes_labels" "kubernetes_labels" {
  name      = "modifierAt=12655"
  clusterid = commvault_kubernetes_cluster.kubernetes_cluster1.id
  namespace = "default"
}

# Resolves the GUID of a namespace by name
data "commvault_kubernetes_namespaces" "kubernetes_namespaces" {
  name      = "1sts-volctemplate"
  clusterid = commvault_kubernetes_cluster.kubernetes_cluster1.id
}

# Resolves the GUID of a StorageClass by name
data "commvault_kubernetes_storageclasses" "kubernetes_storageclasses" {
  name      = "rook-ceph-block"
  clusterid = commvault_kubernetes_cluster.kubernetes_cluster1.id
}

# Resolves the GUID of a PersistentVolumeClaim by name and namespace
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

This example shows optional fields: `filters`, `activitycontrol`, `timezone`, and `options`.
Data sources are scoped to the cluster created in the same config block.

```hcl
# Access node (MediaAgent) for the cluster
data "commvault_client" "access_node1" {
  name = "bdcsrvtest05"
}

# Plan used for etcd protection on the cluster
data "commvault_plan" "plan1" {
  name = "AWS-Test-Plan"
}

# Region to associate with the cluster
data "commvault_region"   "region1" {
  name = "Australia"
}

# Backup plan for the application group
data "commvault_plan" "plan2" {
  name = "Demo Plan"
}

# Timezone used to interpret jobstarttime (seconds from midnight)
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

# Resolves the GUID of a Kubernetes label selector
data "commvault_kubernetes_labels" "kubernetes_labels" {
  name      = "modifierAt=12655"
  clusterid = commvault_kubernetes_cluster.kubernetes_cluster2.id
  namespace = "default"
}

# Resolves the GUID of a namespace by name
data "commvault_kubernetes_namespaces" "kubernetes_namespaces" {
  name      = "1sts-volctemplate"
  clusterid = commvault_kubernetes_cluster.kubernetes_cluster2.id
}

# Resolves the GUID of a StorageClass by name
data "commvault_kubernetes_storageclasses" "kubernetes_storageclasses" {
  name      = "rook-ceph-block"
  clusterid = commvault_kubernetes_cluster.kubernetes_cluster2.id
}

# Resolves the GUID of a PersistentVolumeClaim by name and namespace
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
- `options` (Block List) Appgroup-level operational settings including schedule start time, worker configuration, and snapshot behaviour. (see [below for nested schema](#nestedblock--options))
- `tags` (Block Set) Commvault entity tags (key-value metadata) on the application group resource in CommCell. Use content.labelselectors for Kubernetes label-based content selection. (see [below for nested schema](#nestedblock--tags))
- `timezone` (Block List) Timezone for the application group schedule. Affects when jobstarttime is evaluated. Use the commvault_timezone data source to look up the ID. (see [below for nested schema](#nestedblock--timezone))

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

`content` controls which Kubernetes resources are included in the backup. Three modes are supported:

- **`applications` only** — explicit selection by GUID. Use `commvault_kubernetes_namespaces`,
  `commvault_kubernetes_applications`, or `commvault_kubernetes_volumes` data sources to obtain the GUID,
  or construct it manually as `` namespace`Kind`name`<k8s-uid> ``.
- **`labelselectors` only** — dynamic selection. Any Kubernetes resource matching the given labels
  at backup time is included. Useful for namespace-level or workload-level coverage without listing GUIDs.
- **Both combined** — the union of all matched resources is protected.

The `filters` block is the exclusion counterpart to `content`: resources matching `filters` are removed
from the final backup scope regardless of what `content` selects.

Optional:

- `applications` (Block Set) List of applications to be added as content (see [below for nested schema](#nestedblock--content--applications))
- `labelselectors` (Block Set) List of label selectors to be added as content (see [below for nested schema](#nestedblock--content--labelselectors))

<a id="nestedblock--content--applications"></a>
### Nested Schema for `content.applications`

Required:

- `guid` (String) GUID of the Kubernetes resource as tracked by CommCell. Retrieve via the `commvault_kubernetes_namespaces`, `commvault_kubernetes_applications`, or `commvault_kubernetes_volumes` data source. For unsupported types, construct manually: `` namespace`Kind`name`<k8s-uid> ``.
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

- `guid` (String) GUID of the Kubernetes resource as tracked by CommCell. Retrieve via the `commvault_kubernetes_namespaces`, `commvault_kubernetes_applications`, or `commvault_kubernetes_volumes` data source. For unsupported types, construct manually: `` namespace`Kind`name`<k8s-uid> ``.
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
- `scheduleworkertoconfignamespace` (String) When true, schedules worker Pods into the Commvault config namespace (confignamespace). Enable for CSI snapshot-based backups so the worker can access VolumeSnapshot CRDs. See also: options.workernamespace, cluster options.confignamespace.
- `jobstarttime` (Number) Offset from midnight in seconds at which the backup job starts each day (e.g. 66540 = 18:29:00). Use with the timezone field.
- `snapfallbacktolivevolumebackup` (String) Define setting to enable fallback to live volume backup in case of snap failure
- `workerresources` (Block List) (see [below for nested schema](#nestedblock--options--workerresources))

<a id="nestedblock--options--workerresources"></a>
### Nested Schema for `options.workerresources`

Optional:

- `cpulimits` (String) Maximum CPU for the worker Pod. Kubernetes quantity string, e.g. `"500m"`, `"1"`.
- `cpurequests` (String) Minimum CPU for the worker Pod. Kubernetes quantity string, e.g. `"100m"`, `"0.5"`.
- `memorylimits` (String) Maximum memory for the worker Pod. Kubernetes quantity string, e.g. `"512Mi"`, `"2Gi"`.
- `memoryrequests` (String) Minimum memory for the worker Pod. Kubernetes quantity string, e.g. `"256Mi"`, `"1Gi"`.



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


