package main

import "fmt"

// combination of OCP and Repository demo

// Color represents unique identifier
type Color int

const (
	red Color = iota
	green
	blue
)

// Size represents unique size identifier
type Size int

const (
	small Size = iota
	medium
	large
)

// Product ...
type Product struct {
	name  string
	color Color
	size  Size
}

// Filter ...
type Filter struct {
}

func (f *Filter) filterByColor(
	products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

func (f *Filter) filterBySize(
	products []Product, size Size) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}

	return result
}

func (f *Filter) filterBySizeAndColor(
	products []Product, size Size,
	color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

// filterBySize, filterBySizeAndColor

// Specification ...
type Specification interface {
	IsSatisfied(p *Product) bool
}

// ColorSpecification ...
type ColorSpecification struct {
	color Color
}

// IsSatisfied ...
func (spec ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == spec.color
}

// SizeSpecification ...
type SizeSpecification struct {
	size Size
}

// IsSatisfied ...
func (spec SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == spec.size
}

// AndSpecification ...
type AndSpecification struct {
	first, second Specification
	rest          *AndSpecification
}

// IsSatisfied ...
func (spec AndSpecification) IsSatisfied(p *Product) bool {
	lastSpecSatisfied := spec.rest == nil || spec.rest.IsSatisfied(p)
	return spec.first.IsSatisfied(p) &&
		spec.second.IsSatisfied(p) && lastSpecSatisfied
}

// BetterFilter ...
type BetterFilter struct{}

// Filter ...
func (f *BetterFilter) Filter(
	products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}

	fmt.Print("Green products (old):\n")
	f := Filter{}
	for _, v := range f.filterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}
	// ^^^ BEFORE

	// vvv AFTER
	fmt.Print("Green products (new):\n")
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	largeSpec := SizeSpecification{large}

	largeGreenSpec := AndSpecification{largeSpec, greenSpec, nil}
	fmt.Print("Large blue items:\n")
	for _, v := range bf.Filter(products, largeGreenSpec) {
		fmt.Printf(" - %s is large and green\n", v.name)
	}
}
