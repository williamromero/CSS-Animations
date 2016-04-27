#### Para cambiar tu PEM-key a PKK-key.

* Instalar **putty** en el MacOSX:
```
  sudo port install putty
```

* Esto también instalará **puttygen**. Para obtener puttygen para obtener la salida del archivo .PEM:
```
  puttygen webapp.pem -O private -o webapp.pkk
```

##### Referencias
1. http://stackoverflow.com/questions/3475069/use-ppk-file-in-mac-terminal-to-connect-to-remote-connection-over-ssh/3932540#3932540
2. http://unix.stackexchange.com/questions/116303/convert-amazon-pem-key-to-putty-ppk-key-linux
