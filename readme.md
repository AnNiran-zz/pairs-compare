Generators A and B:
take previous value produced, multiply it by a factor
and keep the remainder of the dividing the resulting product by 2147483647

the final remainder is the value used for the next cycle

Factors:
Generator A: 16807
Generator B: 48271

Starting values:
Generator A: 65
Generator B: 8921

The program count the number of pairs that match among the first 40 million
pairs
Matching is positiove if the least significant 16 bits of both values match