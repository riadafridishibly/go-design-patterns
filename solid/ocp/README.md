# Open Close Principle

Your types should be open for extension, closed for modification. If you need to filter something rather than using different method use one method but make it so that it accepts different spec.

Common approach in sorting in std library. You can pass a function to the sorting algorithm. Otherwise you have to write the whole algorithm for each single criterion.