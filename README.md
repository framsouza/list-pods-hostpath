# List pods that contains HostPath

This code will connect to every GKE cluster and list all the pods that is mount volume as `HostPath`, for more information about `hostPath` see, [Kubernetes HostPath] (https://kubernetes.io/docs/concepts/storage/volumes/#hostpath)

## Usage

First, make sure you have the GOOGLE_APPLICATION_CREDENTIALS environment variable set as per the google docs, https://cloud.google.com/docs/authentication/production

It will require one argument which is `project`, to connect to the specify GKE project to list the pods. By default the zone is set to `-` which means all GKE zones.

You can choose to build the binary by running
```
go build -o list-pods
```

and run it as
```
./list-pods -project=<PROJECTNAME>
```

or without building the binary
```
go run main.go -project=<PROJECTNAME>
```

## Output expected
You will have an output like the following:

```
CLUSTER NAME           POD NAME                                                                        NAMESPACE            VOLUME TYPE                VOLUMES PATH
test-gke          test-76665f4d68-f4gcz                                                                default              Directory                  /                                                        
test-gke          event-exporter-gke-5479fd58c8-p84kc                                                  kube-system                                     /etc/ssl/certs                                           
test-gke          fluentbit-gke-4pcm8                                                                  kube-system                                     /var/run/google-fluentbit/pos-files                      
test-gke          fluentbit-gke-4pcm8                                                                  kube-system                                     /var/log                                                 
test-gke          fluentbit-gke-4pcm8                                                                  kube-system                                     /var/lib/kubelet/pods                                    
test-gke          fluentbit-gke-4pcm8                                                                  kube-system                                     /var/lib/docker/containers                               
test-gke          fluentbit-gke-szwzv                                                                  kube-system                                     /var/run/google-fluentbit/pos-files                      
test-gke          fluentbit-gke-szwzv                                                                  kube-system                                     /var/log                                                 
test-gke          fluentbit-gke-szwzv                                                                  kube-system                                     /var/lib/kubelet/pods                                    
test-gke          fluentbit-gke-szwzv                                                                  kube-system                                     /var/lib/docker/containers                               
test-gke          fluentbit-gke-tr879                                                                  kube-system                                     /var/run/google-fluentbit/pos-files                      
test-gke          fluentbit-gke-tr879                                                                  kube-system                                     /var/log                                                 
test-gke          fluentbit-gke-tr879                                                                  kube-system                                     /var/lib/kubelet/pods                                    
test-gke          fluentbit-gke-tr879                                                                  kube-system                                     /var/lib/docker/containers                               
test-gke          gke-metrics-agent-94mhr                                                              kube-system          Directory                  /etc/ssl/certs                                           
test-gke          gke-metrics-agent-sgg9w                                                              kube-system          Directory                  /etc/ssl/certs                          
```

If the `VOLUME TYPE` column is empty, it means the manifest contains an empty `type`. The `VOLUME PATH` represents the host path that has being mounted inside the pod.