# Sistemas Operativos 2: USAC  - Full Stack App on Kubernetes


# Setear gcloud y kubectl

 - Instalar gcloud, correr el siguiente comando en Windows Powerll y ejecutar instalador

 ```bash

# Correr comando para iniciar configuración
> gcloud init

#Se mostraraá un mensaje como el siguiente, darle "Y" para aceptar y loggearnos en GCP

Network diagnostic detects and fixes local network connection issues.
Checking network connection...done.
Reachability Check passed.
Network diagnostic passed (1/1 checks passed).

You must log in to continue. Would you like to log in (Y/n)? Y

# Se mostrará una lista de proyectos, escogemos el proyecto o creamos uno nuevo.

You are logged in as: [<your-account-email>].

Pick cloud project to use:
 [1] sopes2-proyecto-329600
Please enter numeric choice or text value (must exactly match list item): 2

# Configuramos una región y zona predeterminada, ejemplo us-central1-a

> gcloud config set compute/zone us-central1-a

# Verificamos que haya sido configurada correctamente

> gcloud config list compute/zone

# Instalamos kubectl

> gcloud components install kubectl

# Creamos cluster kubernetes

> gcloud container clusters create squidgames --num-nodes=3 --no-enable-ip-alias

# Recuperando credenciales para Kubectl
> gcloud container clusters get-credentials k8s-demo --zone=us-central1-c

# Creamos reglas de firewall para los puertos
> gcloud compute firewall-rules create fwrule-kubernetes --allow tcp:30000-32767 


 ```


 # Comandos kubectl

 ```bash

# Levantar recursos Frontend
>  kubectl apply -f frontend-deployment.yaml -f frontend-service.yaml

# Borrar recursos Frontend
>  kubectl delete -f frontend-deployment.yaml -f frontend-service.yaml

# Levantar recursos Backend
>  kubectl apply -f backend-deployment.yaml -f backend-service.yaml

# Borrar recursos Backend
>  kubectl delete -f backend-deployment.yaml -f backend-service.yaml

 ```
## Seteando nueva imagen en deployment ejecutandose.

```bash
> kubectl set image deployments/<deployment name> <label name>=<new image registry name:tage>

# Ejemplo
> kubectl set image deployments/frontend-deployment frontend-deployment=oscarllamas6/frontend-changed:v1
```
## Verificar estado del rolling update

```bash
> kubectl rollout status deployment/frontend-deployment
```

## Retroceder con rollout updates

```bash
> kubectl rollout undo deployments/frontend-deployment
```



