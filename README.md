# Guía de Aprendizaje: Go desde Cero

Colección de 10 proyectos progresivos para aprender los fundamentos del lenguaje Go.

---

## Requisitos

- [Go](https://go.dev/dl/) instalado (versión 1.18 o superior)
- Ejecutar cualquier ejemplo con: `go run <carpeta>/main.go`

---

## Estructura del Proyecto

```
pruebas_Go/
├── 1-primeros_pasos/
├── 2-variables/
├── 3-condicionales/
├── 4-bucles/
├── 5-funciones/
├── 6-slices/
├── 7-mapas/
├── 8-structs/
├── 9-goroutines/
└── 10-errores/
```

---

## 1. Primeros Pasos — Hola Mundo

**Carpeta:** `1-primeros_pasos/`

El programa más básico en Go. Todo programa Go pertenece a un `package` y el punto de entrada es siempre la función `main`.

**Conceptos:** `package main`, `import`, `func main()`, `fmt.Println`

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hola Mundo")
}
```

**Salida:**
```
Hola Mundo
```

---

## 2. Variables y Tipos de Datos

**Carpeta:** `2-variables/`

Go es un lenguaje de tipado estático. Las variables pueden declararse explícitamente con `var` o usando el operador de inferencia de tipo `:=`.

**Conceptos:** `var`, `:=`, `string`, `int`, `bool`, `float64`

```go
package main

import "fmt"

func main() {
    var nombre string = "Go"
    var version int = 1
    activo := true
    precio := 3.14

    fmt.Println("Lenguaje:", nombre)
    fmt.Println("Version:", version)
    fmt.Println("Activo:", activo)
    fmt.Println("Precio:", precio)
}
```

**Salida:**
```
Lenguaje: Go
Version: 1
Activo: true
Precio: 3.14
```

**Notas:**
- `var nombre string = "Go"` — declaración explícita con tipo.
- `activo := true` — declaración corta, Go infiere el tipo automáticamente.
- Los tipos básicos son: `string`, `int`, `float64`, `bool`.

---

## 3. Condicionales

**Carpeta:** `3-condicionales/`

Go soporta `if/else` y `switch`. A diferencia de otros lenguajes, las condiciones no requieren paréntesis.

**Conceptos:** `if`, `else`, `switch`, `case`, `default`

```go
package main

import "fmt"

func main() {
    edad := 20

    if edad >= 18 {
        fmt.Println("Mayor de edad")
    } else {
        fmt.Println("Menor de edad")
    }

    dia := "lunes"
    switch dia {
    case "lunes":
        fmt.Println("Inicio de semana")
    case "viernes":
        fmt.Println("Fin de semana laboral")
    default:
        fmt.Println("Dia normal")
    }
}
```

**Salida:**
```
Mayor de edad
Inicio de semana
```

**Notas:**
- El `switch` en Go no necesita `break`, cada `case` termina automáticamente.
- El bloque `default` se ejecuta si ningún `case` coincide.

---

## 4. Bucles

**Carpeta:** `4-bucles/`

Go tiene un único tipo de bucle: `for`. Con él se pueden expresar tanto el bucle clásico como el bucle tipo `while`.

**Conceptos:** `for`, incremento, condición de parada

```go
package main

import "fmt"

func main() {
    // Bucle clásico
    for i := 1; i <= 5; i++ {
        fmt.Println("Iteracion:", i)
    }

    // Bucle tipo while
    n := 10
    for n > 0 {
        fmt.Println("Contando:", n)
        n -= 3
    }
}
```

**Salida:**
```
Iteracion: 1
Iteracion: 2
Iteracion: 3
Iteracion: 4
Iteracion: 5
Contando: 10
Contando: 7
Contando: 4
Contando: 1
```

**Notas:**
- `for i := 0; i < n; i++` — forma clásica con inicialización, condición e incremento.
- `for condicion {}` — equivalente al `while` de otros lenguajes.
- `for {}` — bucle infinito (sin condición).

---

## 5. Funciones

**Carpeta:** `5-funciones/`

Las funciones en Go pueden retornar múltiples valores, lo cual es muy utilizado para retornar un resultado junto con un posible error.

**Conceptos:** `func`, parámetros tipados, retorno múltiple, `_` para ignorar valores

```go
package main

import "fmt"

func sumar(a, b int) int {
    return a + b
}

func dividir(a, b float64) (float64, string) {
    if b == 0 {
        return 0, "error: division por cero"
    }
    return a / b, "ok"
}

func main() {
    fmt.Println("Suma:", sumar(5, 3))

    resultado, estado := dividir(10, 2)
    fmt.Println("Division:", resultado, "-", estado)

    _, err := dividir(10, 0)
    fmt.Println("Error:", err)
}
```

**Salida:**
```
Suma: 8
Division: 5 - ok
Error: error: division por cero
```

**Notas:**
- Cuando dos parámetros son del mismo tipo, se puede escribir `func sumar(a, b int)` en lugar de `func sumar(a int, b int)`.
- El identificador `_` descarta un valor de retorno que no se necesita.

---

## 6. Arrays y Slices

**Carpeta:** `6-slices/`

Los arrays tienen tamaño fijo. Los slices son la versión dinámica y son los más usados en Go.

**Conceptos:** arrays, slices, `append`, `range`

```go
package main

import "fmt"

func main() {
    // Array fijo
    colores := [3]string{"rojo", "verde", "azul"}
    fmt.Println("Array:", colores)

    // Slice dinámico
    numeros := []int{1, 2, 3}
    numeros = append(numeros, 4, 5)
    fmt.Println("Slice:", numeros)

    // Recorrer con range
    for i, v := range numeros {
        fmt.Printf("numeros[%d] = %d\n", i, v)
    }
}
```

**Salida:**
```
Array: [rojo verde azul]
Slice: [1 2 3 4 5]
numeros[0] = 1
numeros[1] = 2
numeros[2] = 3
numeros[3] = 4
numeros[4] = 5
```

**Notas:**
- `[3]string{}` — array de tamaño fijo 3.
- `[]int{}` — slice sin tamaño fijo.
- `append(slice, valor)` — agrega elementos al slice y retorna el nuevo slice.
- `range` retorna el índice y el valor en cada iteración.

---

## 7. Mapas

**Carpeta:** `7-mapas/`

Los mapas son colecciones de pares clave-valor. Son equivalentes a los diccionarios en Python o los objetos en JavaScript.

**Conceptos:** `map`, acceso por clave, `delete`, iteración con `range`

```go
package main

import "fmt"

func main() {
    capitales := map[string]string{
        "Mexico":    "Ciudad de Mexico",
        "Argentina": "Buenos Aires",
        "España":    "Madrid",
    }

    fmt.Println("Capitales:", capitales)
    fmt.Println("Capital de Mexico:", capitales["Mexico"])

    capitales["Colombia"] = "Bogota"
    delete(capitales, "España")

    for pais, capital := range capitales {
        fmt.Printf("%s -> %s\n", pais, capital)
    }
}
```

**Salida:**
```
Capitales: map[Argentina:Buenos Aires Mexico:Ciudad de Mexico España:Madrid]
Capital de Mexico: Ciudad de Mexico
Mexico -> Ciudad de Mexico
Argentina -> Buenos Aires
Colombia -> Bogota
```

**Notas:**
- `map[string]string` — mapa con claves y valores de tipo `string`.
- `delete(mapa, clave)` — elimina una entrada del mapa.
- El orden de iteración de un mapa no está garantizado en Go.

---

## 8. Structs y Métodos

**Carpeta:** `8-structs/`

Los structs permiten agrupar campos relacionados. Se pueden asociar métodos a un struct usando un receptor.

**Conceptos:** `type`, `struct`, métodos con receptor, `fmt.Sprintf`

```go
package main

import "fmt"

type Persona struct {
    Nombre string
    Edad   int
}

func (p Persona) Saludar() string {
    return fmt.Sprintf("Hola, soy %s y tengo %d años", p.Nombre, p.Edad)
}

func main() {
    p1 := Persona{Nombre: "Ana", Edad: 30}
    p2 := Persona{"Luis", 25}

    fmt.Println(p1.Saludar())
    fmt.Println(p2.Saludar())
    fmt.Println("Nombre:", p1.Nombre, "| Edad:", p1.Edad)
}
```

**Salida:**
```
Hola, soy Ana y tengo 30 años
Hola, soy Luis y tengo 25 años
Nombre: Ana | Edad: 30
```

**Notas:**
- `func (p Persona) Saludar()` — `p` es el receptor, similar a `self` en Python o `this` en otros lenguajes.
- Los campos del struct se acceden con notación de punto: `p1.Nombre`.
- `fmt.Sprintf` formatea un string sin imprimirlo, retornándolo como valor.

---

## 9. Goroutines y Concurrencia

**Carpeta:** `9-goroutines/`

Las goroutines son funciones que se ejecutan de forma concurrente. Son muy ligeras comparadas con los hilos del sistema operativo. `sync.WaitGroup` permite esperar a que todas terminen.

**Conceptos:** `go`, `sync.WaitGroup`, `wg.Add`, `wg.Done`, `wg.Wait`, `defer`

```go
package main

import (
    "fmt"
    "sync"
)

func saludar(nombre string, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Println("Hola desde goroutine:", nombre)
}

func main() {
    var wg sync.WaitGroup

    nombres := []string{"Alice", "Bob", "Carlos"}

    for _, nombre := range nombres {
        wg.Add(1)
        go saludar(nombre, &wg)
    }

    wg.Wait()
    fmt.Println("Todas las goroutines terminaron")
}
```

**Salida (el orden puede variar):**
```
Hola desde goroutine: Carlos
Hola desde goroutine: Alice
Hola desde goroutine: Bob
Todas las goroutines terminaron
```

**Notas:**
- `go func()` — lanza la función como goroutine en segundo plano.
- `wg.Add(1)` — registra que hay una goroutine más por esperar.
- `defer wg.Done()` — marca la goroutine como terminada al salir de la función.
- `wg.Wait()` — bloquea hasta que todas las goroutines llamen a `Done()`.
- El orden de ejecución de las goroutines no es determinista.

---

## 10. Manejo de Errores

**Carpeta:** `10-errores/`

Go no usa excepciones. El patrón estándar es retornar un valor de tipo `error` junto al resultado, y verificarlo con `if err != nil`.

**Conceptos:** `error`, `errors.New`, `nil`, patrón `if err != nil`

```go
package main

import (
    "errors"
    "fmt"
)

func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("no se puede dividir por cero")
    }
    return a / b, nil
}

func main() {
    resultado, err := dividir(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Resultado:", resultado)
    }

    _, err = dividir(5, 0)
    if err != nil {
        fmt.Println("Error:", err)
    }
}
```

**Salida:**
```
Resultado: 5
Error: no se puede dividir por cero
```

**Notas:**
- `error` es una interfaz nativa de Go.
- `errors.New("mensaje")` crea un error simple con un mensaje.
- `nil` indica ausencia de error (equivalente a `null` en otros lenguajes).
- Siempre verificar el error antes de usar el resultado.

---

## Resumen de Conceptos

| # | Proyecto | Conceptos Clave |
|---|---|---|
| 1 | Primeros Pasos | `package`, `import`, `func main`, `fmt.Println` |
| 2 | Variables | `var`, `:=`, tipos básicos |
| 3 | Condicionales | `if/else`, `switch/case` |
| 4 | Bucles | `for` clásico, `for` como while |
| 5 | Funciones | parámetros, retorno múltiple, `_` |
| 6 | Arrays y Slices | arrays fijos, slices, `append`, `range` |
| 7 | Mapas | `map`, `delete`, iteración |
| 8 | Structs | `type`, `struct`, métodos con receptor |
| 9 | Goroutines | `go`, `sync.WaitGroup`, concurrencia |
| 10 | Errores | `error`, `errors.New`, `nil`, `if err != nil` |

---

## Recursos Adicionales

- [Tour oficial de Go](https://go.dev/tour/)
- [Documentación oficial](https://go.dev/doc/)
- [Go by Example](https://gobyexample.com/)
- [Playground de Go](https://go.dev/play/)
