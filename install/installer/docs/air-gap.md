# Installing Nxpod in an air-gap network with the Nxpod Installer

## Mirror Nxpod Images

You need a registry that is reachable in your network. Add the domain of your registry to the Nxpod config `gitpod.config.yaml` like this:
```yaml
repository: your-registry.example.com
```

The command `gitpod-installer mirror list` gives you a list of all images needed by Nxpod. You can run the following code to pull the needed images and push them to your registry:

```
for row in $(gitpod-installer mirror list --config gitpod.config.yaml | jq -c '.[]'); do
    original=$(echo $row | jq -r '.original')
    target=$(echo $row | jq -r '.target')

    docker pull $original
    docker tag $original $target
    docker push $target
done
```

## Install Nxpod in Air-Gap Mode

To install Nxpod in an air-gap network, you need to configure the repository of the images needed by Nxpod (see previous step). Add this to your Nxpod config:

```yaml
repository: your-registry.example.com
```

That's it. Run the following commands as usual and Nxpod fetches the images from your registry and does not need internet access to operate:

```
gitpod-installer render --config gitpod.config.yaml > gitpod.yaml
kubectl apply -f gitpod.yaml
```
