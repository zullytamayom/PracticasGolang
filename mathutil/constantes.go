package mathutil

const (
	Pi = 3.14
)

var DiasSemana = []string{
	"Lunes", "Martes", "Miercoles", "Jueves", "Viernes",
}

func Multiplicar(valor1 int) float64 {
	resultado := float64(valor1) * Pi
	return resultado
}
