
# Dependency Management

## `vendor` Verzeichnis
*Effekt*:
- Packages in einem Unterverzeichnis `/vendor` werden bei go builds priorisert angezogen.
- Files in `/vendor` können mit ins git repository eingecheckt werden, müssen aber nicht.
- Die Abhängigkeiten können mit Tools z.B. `godep` oder `glide` gemanaged werden.
- Ein offizielles Tool ist gerade in Arbeit: [Dep](https://github.com/golang/dep)

*Vorsicht:*
Die selbe Library in unterschieldichen vendor-Verzeichnissen macht Probleme.
Siehe auch: https://github.com/mattfarina/golang-broken-vendor

## glide
Glide ist ein package manager mit semantic versioning.

```shell
go get github.com/masterminds/glide
```

*Benutzung:*
https://github.com/Masterminds/glide

```shell
# Erstellen einer glide.yaml Konfigurationsdatei
glide init

# Update aller Dependencies ensprechend der glide.yaml
# Schreibt die konkreten Versionen in die glide.lock
glide update

# Installieren der Versionen aus glide.lock in /vendor
glide install
```

## Alternativen

### gopkg.in
Öffentlicher Proxy für den Import von Packages.

```shell
go get gopkg.in/yaml.v1
```

### Github forks
Eine einfache und sichere Variaten ist das erstellen von forks für genutzte Dependencies. Entweder auf github, oder in einem internen Repository.


