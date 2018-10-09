package with_interface

// public interface
type DoSimple interface {
	GetFirstname() string
	GetLastname() string
}

// implementation
type doSimple struct {
	firstname, lastname string
}

// NOTES: we need stub the receiver method, assume that this method have http request, kafka, redis
// we don't want do with real connect.
func (d *doSimple) GetFirstname() string {
	return d.firstname
}

func (d *doSimple) GetLastname() string {
	return d.lastname
}

// implement interface
func FetchDoSimple(d DoSimple) string {
	// return d.GetFirstname(), d.GetLastname()
	return d.GetFirstname()
}

func NewDoSimple(f, l string) DoSimple {
	return &doSimple{
		firstname: f,
		lastname:  l,
	}
}

// caller
// func main() {
// 	// new instance object
// 	// new := &doSimple{
// 	// 	firstname: "Teerapong",
// 	// 	lastname:  "Singthong",
// 	// }

// 	// f, l := FetchDoSimple(new)
// 	// fmt.Println(f)
// 	// fmt.Println(l)

// 	new := NewDoSimple("Teerapong", "Singthong")
// 	fmt.Println(new.GetFirstname())
// 	fmt.Println(new.GetLastname())
// }
