#Configurar un servidor AWS EC2 para Rails

1. Seleccionar la microinstancia que ofrece AWS. 
![alt text](https://platzi.com/blog/wp-content/uploads/2015/07/Create_instanse1.png "Create Instance")

2. Configurar las reglas de seguridad de tu instancia.
![alt text](https://platzi.com/blog/wp-content/uploads/2015/07/SecurityGroup.png "Firewall Instance")
![alt text](https://github.com/williamromero/Code-Snippets/blob/master/RoR/Screen%20Shot%202016-04-25%20at%2011.02.59%20AM.png?raw=true "Firewall Instance")

3. Crear o descargar la llave de seguridad privada para poder acceder a tu servidor.
![alt text](https://platzi.com/blog/wp-content/uploads/2015/07/PEM-Key.png "Firewall Instance")

4. Una vez descargada, debemos de instalar la llave en la carpeta ~/.ssh/ de nuestro ordenador MacOSX, por lo tanto vamos a dirigirnos a la carpeta con el siguiente comando:

```
  cd /                                                            # Nos dirige al root del ordenador.
  cd ~/.ssh/                                                      # Nos dirige a la carpeta SSH en donde se guardan las llaves de seguridad de los servicios.
  mkdir webapp                                                    # Creamos una carpeta, con el nombre de la llave de nuestro servicio.
  cd webapp                                                       # Ingresamos a la carpeta que hemos creado.
  cp /Users/tu_nombre_de_usuario/Downloads/webapp.pem webapp.pem  # Copiamos la llave de seguridad de la carpeta en la que se descargó, hasta la carpeta "webapp" que se encuentra en el folder ~/.ssh/.
```
...
5. Antes de conectarnos a la instancia tendremos que configurar los permisos de la llave que ya descargamos para poder usarla en modo lectura únicamente. Para alterar los permisos de la llave usamos el comando:
```
  chmod 0400 webapp.pem
```
...
6. Ahora, para acceder por medio de SSH, nos dirigimos a la ventana de comandos y colocamos:
```
  $ ssh -i "webapp.pem" ubuntu@ec2-00-00-000-00.us-west-2.compute-1.amazonaws.com
           *Key Name*                     *Public DNS*
```
...
7. Ya dentro del servidor, debemos de actualizar todas las dependencias:
```
  sudo apt-get update && sudo apt-get -y upgrade
```
...
8. Luego, instalar CURL & git:
```
  sudo apt-get install git CURL
```
...
9. Luego de instalar estas dependencias, instalamos el RVM & Ruby:
```
  \curl -L https://get.rvm.io | bash -s stable --ruby
```

* Si en algún caso, cometieramos un error en seguir los pasos o cualquiera que sea ingresando la llave privada y el registro de esta quedase grabado
dentro de la memoria de las rutas ssh, podemos borrarlo con el siguiente comando:

```
  ssh-keygen -R ec2-00-00-000-00.us-west-2.compute-1.amazonaws.com
```

###Artículo en producción

***
Referencias:
* http://www.sitepoint.com/deploy-your-rails-app-to-aws/
* https://platzi.com/blog/instalar-ghost-amazon-ec2/
* http://docs.aws.amazon.com/gettingstarted/latest/wah-linux/getting-started-deploy-app.html
* https://manuais.iessanclemente.net/index.php/Servidor_Virtual_VPS_con_Amazon_EC2_-_Instalaci%C3%B3n_y_configuraci%C3%B3n
* http://superuser.com/questions/30087/remove-key-from-known-hosts
* https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet#lists
* http://railsapps.github.io/installrubyonrails-ubuntu.html


### Deploy
***
* http://www.sitepoint.com/deploy-your-rails-app-to-aws/
