
## 03 Structs und Pointer

### Structs
Typdefinition:
```go
type Person struct {
	Name  string
	Given string
	Age   int
}
```

Verwendung:
```go
person := Person{
	Name:  "Mancke",
	Given: "Sebastian",
	Age:   42,
}
fmt.Println(person)

person.Given = "Felix"
fmt.Println(person)
```

### Pointer
* Parameterübergabe immer Call-by-Value (auch bei struct)
* Typsichere Pointer können verwendet werden
```go
a := 41

var b *int
b = &a

*b++

fmt.Println(a) // 42
```

### Pointer und Structs
Referenzen auf structs:
```go
// copy by value
person1 := Person{
	Name: "Mancke",
}
    
person2 := person1
person2.Name = "Meyer"

fmt.Println(person1.Name) // Mancke

// copy by reference
person3 := &person1
person3.Name = "Meyer"

fmt.Println(person1.Name) // Meyer
```

### Pointer und Slices/Maps
Mit `make()` erstellte Datentypen sind immer Pointer-Typen.

```go
colors1 := []string{"red", "blue"}

colors2 := colors1
colors2[0] = "black"
colors2[1] = "white"
    
fmt.Println(colors1) // black, white
```

## Übung: Key-Value Store
Schreibe einen kleines Programm `kv`, mit dem Du Schlüssel-Werte
Paare in einer Datei speichern und abfragen kannst.

Setzen von Werten:
```shell
kv name=Mancke vorname=Sebastian alter=42
```

Abfragen bestimmter Werte:
```shell
kv name vorname
> name=Mancke
> vorname=Sebastian
```

Abfragen aller Werte:
```shell
kv
> name=Mancke
> vorname=Sebastian
> alter=42
```

