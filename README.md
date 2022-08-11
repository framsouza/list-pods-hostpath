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
CLUSTER NAME           POD NAME                                                         NAMESPACE            VOLUME TYPE                VOLUMES PATH
maritests-gke          test-76665f4d68-f4gcz                                            default              Directory                  /                                                        
maritests-gke          event-exporter-gke-5479fd58c8-p84kc                              kube-system                                     /etc/ssl/certs                                           
maritests-gke          fluentbit-gke-4pcm8                                              kube-system                                     /var/run/google-fluentbit/pos-files                      
maritests-gke          fluentbit-gke-4pcm8                                              kube-system                                     /var/log                                                 
maritests-gke          fluentbit-gke-4pcm8                                              kube-system                                     /var/lib/kubelet/pods                                    
maritests-gke          fluentbit-gke-4pcm8                                              kube-system                                     /var/lib/docker/containers                               
maritests-gke          fluentbit-gke-szwzv                                              kube-system                                     /var/run/google-fluentbit/pos-files                      
maritests-gke          fluentbit-gke-szwzv                                              kube-system                                     /var/log                                                 
maritests-gke          fluentbit-gke-szwzv                                              kube-system                                     /var/lib/kubelet/pods                                    
maritests-gke          fluentbit-gke-szwzv                                              kube-system                                     /var/lib/docker/containers                               
maritests-gke          fluentbit-gke-tr879                                              kube-system                                     /var/run/google-fluentbit/pos-files                      
maritests-gke          fluentbit-gke-tr879                                              kube-system                                     /var/log                                                 
maritests-gke          fluentbit-gke-tr879                                              kube-system                                     /var/lib/kubelet/pods                                    
maritests-gke          fluentbit-gke-tr879                                              kube-system                                     /var/lib/docker/containers                               
maritests-gke          gke-metrics-agent-94mhr                                          kube-system          Directory                  /etc/ssl/certs                                           
maritests-gke          gke-metrics-agent-sgg9w                                          kube-system          Directory                  /etc/ssl/certs                                           
maritests-gke          gke-metrics-agent-zpmwm                                          kube-system          Directory                  /etc/ssl/certs                                           
maritests-gke          kube-proxy-gke-maritests-gke-default-pool-8608c62b-xfzh          kube-system                                     /usr/share/ca-certificates                               
maritests-gke          kube-proxy-gke-maritests-gke-default-pool-8608c62b-xfzh          kube-system                                     /etc/ssl/certs                                           
maritests-gke          kube-proxy-gke-maritests-gke-default-pool-8608c62b-xfzh          kube-system          FileOrCreate               /var/lib/kube-proxy/kubeconfig                           
maritests-gke          kube-proxy-gke-maritests-gke-default-pool-8608c62b-xfzh          kube-system                                     /var/log                                                 
maritests-gke          kube-proxy-gke-maritests-gke-default-pool-8608c62b-xfzh          kube-system          FileOrCreate               /run/xtables.lock                                        
maritests-gke          kube-proxy-gke-maritests-gke-default-pool-8608c62b-xfzh          kube-system                                     /lib/modules                                             
maritests-gke          kube-proxy-gke-maritests-gke-default-pool-af2b35ac-wr63          kube-system                                     /usr/share/ca-certificates                               
maritests-gke          kube-proxy-gke-maritests-gke-default-pool-af2b35ac-wr63          kube-system                                     /etc/ssl/certs                                           
maritests-gke          kube-proxy-gke-maritests-gke-default-pool-af2b35ac-wr63          kube-system          FileOrCreate               /var/lib/kube-proxy/kubeconfig                           
maritests-gke          kube-proxy-gke-maritests-gke-default-pool-af2b35ac-wr63          kube-system                                     /var/log                                                 
maritests-gke          kube-proxy-gke-maritests-gke-default-pool-af2b35ac-wr63          kube-system          FileOrCreate               /run/xtables.lock                                        
maritests-gke          kube-proxy-gke-maritests-gke-default-pool-af2b35ac-wr63          kube-system                                     /lib/modules                                             
maritests-gke          kube-proxy-gke-maritests-gke-default-pool-cebcda79-lpz3          kube-system                                     /usr/share/ca-certificates                               
maritests-gke          kube-proxy-gke-maritests-gke-default-pool-cebcda79-lpz3          kube-system                                     /etc/ssl/certs                                           
maritests-gke          kube-proxy-gke-maritests-gke-default-pool-cebcda79-lpz3          kube-system          FileOrCreate               /var/lib/kube-proxy/kubeconfig                           
maritests-gke          kube-proxy-gke-maritests-gke-default-pool-cebcda79-lpz3          kube-system                                     /var/log                                                 
maritests-gke          kube-proxy-gke-maritests-gke-default-pool-cebcda79-lpz3          kube-system          FileOrCreate               /run/xtables.lock                                        
maritests-gke          kube-proxy-gke-maritests-gke-default-pool-cebcda79-lpz3          kube-system                                     /lib/modules                                             
maritests-gke          metricbeat-bw7gl                                                 kube-system                                     /proc                                                    
maritests-gke          metricbeat-bw7gl                                                 kube-system                                     /sys/fs/cgroup                                           
maritests-gke          metricbeat-bw7gl                                                 kube-system          DirectoryOrCreate          /var/lib/metricbeat-data                                 
maritests-gke          metricbeat-mjr9c                                                 kube-system                                     /proc                                                    
maritests-gke          metricbeat-mjr9c                                                 kube-system                                     /sys/fs/cgroup                                           
maritests-gke          metricbeat-mjr9c                                                 kube-system          DirectoryOrCreate          /var/lib/metricbeat-data                                 
maritests-gke          metricbeat-tlz88                                                 kube-system                                     /proc                                                    
maritests-gke          metricbeat-tlz88                                                 kube-system                                     /sys/fs/cgroup                                           
maritests-gke          metricbeat-tlz88                                                 kube-system          DirectoryOrCreate          /var/lib/metricbeat-data                                 
maritests-gke          pdcsi-node-bpjhj                                                 kube-system          Directory                  /var/lib/kubelet/plugins_registry/                       
maritests-gke          pdcsi-node-bpjhj                                                 kube-system          Directory                  /var/lib/kubelet                                         
maritests-gke          pdcsi-node-bpjhj                                                 kube-system          DirectoryOrCreate          /var/lib/kubelet/plugins/pd.csi.storage.gke.io/          
maritests-gke          pdcsi-node-bpjhj                                                 kube-system          Directory                  /dev                                                     
maritests-gke          pdcsi-node-bpjhj                                                 kube-system          Directory                  /etc/udev                                                
maritests-gke          pdcsi-node-bpjhj                                                 kube-system          Directory                  /lib/udev                                                
maritests-gke          pdcsi-node-bpjhj                                                 kube-system          Directory                  /run/udev                                                
maritests-gke          pdcsi-node-bpjhj                                                 kube-system          Directory                  /sys                                                     
maritests-gke          pdcsi-node-c4qkw                                                 kube-system          Directory                  /var/lib/kubelet/plugins_registry/                       
maritests-gke          pdcsi-node-c4qkw                                                 kube-system          Directory                  /var/lib/kubelet                                         
maritests-gke          pdcsi-node-c4qkw                                                 kube-system          DirectoryOrCreate          /var/lib/kubelet/plugins/pd.csi.storage.gke.io/          
maritests-gke          pdcsi-node-c4qkw                                                 kube-system          Directory                  /dev                                                     
maritests-gke          pdcsi-node-c4qkw                                                 kube-system          Directory                  /etc/udev                                                
maritests-gke          pdcsi-node-c4qkw                                                 kube-system          Directory                  /lib/udev                                                
maritests-gke          pdcsi-node-c4qkw                                                 kube-system          Directory                  /run/udev                                                
maritests-gke          pdcsi-node-c4qkw                                                 kube-system          Directory                  /sys                                                     
maritests-gke          pdcsi-node-qcmz4                                                 kube-system          Directory                  /var/lib/kubelet/plugins_registry/                       
maritests-gke          pdcsi-node-qcmz4                                                 kube-system          Directory                  /var/lib/kubelet                                         
maritests-gke          pdcsi-node-qcmz4                                                 kube-system          DirectoryOrCreate          /var/lib/kubelet/plugins/pd.csi.storage.gke.io/          
maritests-gke          pdcsi-node-qcmz4                                                 kube-system          Directory                  /dev                                                     
maritests-gke          pdcsi-node-qcmz4                                                 kube-system          Directory                  /etc/udev                                                
maritests-gke          pdcsi-node-qcmz4                                                 kube-system          Directory                  /lib/udev                                                
maritests-gke          pdcsi-node-qcmz4                                                 kube-system          Directory                  /run/udev                                                
maritests-gke          pdcsi-node-qcmz4                                                 kube-system          Directory                  /sys                                           
```

If the `VOLUME TYPE` column is empty, it means the manifest contains an empty `type`. The `VOLUME PATH` represents the host path that has being mounted inside the pod. The pods that are **NOT** using `hostPath` won't be in the list.