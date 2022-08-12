package Section3

import (
	"testing"
)

/*
go test
go test -v
go test -v -run TestNamaFunction
go test -run TestNamaFunction/NamaSubTest
go test run /NamaSubTest

go test -v ./...
*/
/*
func TestHello(t *testing.T) {
	result := HelloWorld("Eko")
	if result != "Hello Eko" {
		// t.Fail()
		t.Error("Result must be 'Hello Eko'")
	}
	fmt.Println("Test Hello Success")
}

func TestHelloName(t *testing.T) {
	result := HelloWorld("Khannedy")
	if result != "Hello Khannedy" {
		// t.FailNow()
		t.Fatal("Result must be 'Hello Khannedy'") // sama kyk fatal now
	}
	fmt.Println("Test Hello Name Success")
}

func TestHelloNameAssert(t *testing.T) {
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result, "Result Must be Hello Eko") //assert manggil fail
	fmt.Println("Dieksekusi")
}

func TestHelloNameRequire(t *testing.T) {
	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "Result Must be Hello Eko") //require manggil failNow
	fmt.Println("Dieksekusi")
}

func TestHelloNameSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Unit test tidak bisa jalan di MAC")
	}
	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result)
}

func TestMain(m *testing.M) {
	fmt.Println("before unit test")
	m.Run()
	fmt.Println("After unit test")
}

func TestSubTest(t *testing.T) {
	t.Run("Eko", func(t *testing.T) {
		result := HelloWorld("Eko")
		require.Equal(t, "Hello Eko", result, "Result Must Be 'Hello Eko'")
	})

	t.Run("Kurniawan", func(t *testing.T) {
		result := HelloWorld("Kurniawan")
		require.Equal(t, "Hello Kurniawan", result, "Result Must Be 'Hello Kurniawan'")
	})
}
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Eko)",
			request:  "Eko",
			expected: "Hello Eko",
		},

		{
			name:     "HelloWorld(Kurniawan)",
			request:  "Kurniawan",
			expected: "Hello Kurniawan",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}
*/
/*
go test -v -bench=.
go test -v -run=NotMathUnitTest -bench=.
go test -v -bench=. ./...
go test -v -bench=BenchMarkNama/SubBench
*/

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Eko")
	}
}

func BenchmarkHelloWorldKurniawan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Kurniawan")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Eko", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Eko")
		}
	})
	b.Run("Kurniawan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Kurniawan")
		}
	})
}
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Eko",
			request: "Eko",
		},

		{
			name:    "Kurniawan",
			request: "Kurniawan",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}
