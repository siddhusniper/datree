
# Custom Policies To Validate Kubernetes Mainfests Using Datree CLI

Kubernetes labels enable engineers to perform in-cluster object searches, apply bulk configuration changes, and more. Labels can help simplify and solve many day-to-day challenges encountered in Kubernetes environments if they are set correctly.

__This custom policies helps to enforce the following labels best practices for use cases:__
* [Ensure strategy has pre-defined labels ](#ensure-strategy-has-pre-defined-labels)
* [Ensure RollingUpdate strategy has maxSurge and maxUnavailable labels](#ensure-RollingUpdate-strategy-has-maxSurge-and-maxUnavailable-labels)
* [Ensure pre-defined DnsPolicy labels are used for pods](#ensure-pre-defined-DnsPolicy-labels-are-used-for-pods)
* [Ensure custom nodeselector has pre-defined label](#ensure-custom-nodeselector-has-pre-defined-label)


## Ensure strategy has pre-defined labels 
Kubernetes deployment strategies are use to replace existing pods with new ones. This rule will also ensure that only pre-approved `strategy` label values are used:
* `Recreate`
* `RollingUpdate`
### When this rule is failing?
If the `strategy` key is missing from the labels section:  
```
kind: Deployment
spec:
  replicas: 2
  template:
```

__OR__ a different `strategy.type` value is used:
```
kind: Deployment
spec:
  strategy:
    type: Ab
```


## Ensure RollingUpdate strategy has maxSurge and maxUnavailable labels
RollingUpdate strategy replace the old ReplicaSets by new one using rolling update i.e gradually scale down the old ReplicaSets and scale up the new one.
Rolling update config params. Present only if DeploymentStrategyType = RollingUpdate. This rule will ensure that `rollingupdate` labels are numeric
* `maxSurge`
* `maxUnavailable`  
params control the desired behavior of rolling update.
### When this rule is failing?
while `DeploymentStrategyType` = `RollingUpdate` but `maxSurge` and `maxUnavailable`  are not defined
```
kind: Deployment
spec:
  strategy:
    type: RollingUpdate
    rollingUpdate:
```
__OR__ assigned non-numeric labels

```
kind: Deployment
spec:
  strategy:
    type: RollingUpdate
    rollingUpdate:
     maxSurge: ab
     maxUnavailable: aa
```