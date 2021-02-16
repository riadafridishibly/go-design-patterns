package main

// Document ...
type Document struct {
}

// Machine ...
type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

// MultiFunctionPrinter ...
// ok if you need a multifunction device
type MultiFunctionPrinter struct {
	// ...
}

// Print ...
func (m MultiFunctionPrinter) Print(d Document) {

}

// Fax ...
func (m MultiFunctionPrinter) Fax(d Document) {

}

// Scan ...
func (m MultiFunctionPrinter) Scan(d Document) {

}

// OldFashionedPrinter ...
type OldFashionedPrinter struct {
	// ...
}

// Print ...
func (o OldFashionedPrinter) Print(d Document) {
	// ok
}

// Fax ...
func (o OldFashionedPrinter) Fax(d Document) {
	panic("operation not supported")
}

// Scan ...
// Deprecated: ...
func (o OldFashionedPrinter) Scan(d Document) {
	panic("operation not supported")
}

// Printer ...
// better approach: split into several interfaces
type Printer interface {
	Print(d Document)
}

// Scanner ...
type Scanner interface {
	Scan(d Document)
}

// MyPrinter ...
// printer only
type MyPrinter struct {
	// ...
}

// Print ...
func (m MyPrinter) Print(d Document) {
	// ...
}

// Photocopier ...
// combine interfaces
type Photocopier struct{}

// Scan ...
func (p Photocopier) Scan(d Document) {
	//
}

// Print ...
func (p Photocopier) Print(d Document) {
	//
}

// MultiFunctionDevice ...
type MultiFunctionDevice interface {
	Printer
	Scanner
}

// MultiFunctionMachine ...
// interface combination + decorator
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

// Print ...
func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

// Scan ...
func (m MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}

func main() {

}
