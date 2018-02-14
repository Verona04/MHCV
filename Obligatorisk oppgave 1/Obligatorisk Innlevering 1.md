
<h1>Obligatorisk Innlevering 1<br>
9. februar 2018</h1>
<h2>Gruppe MCV <br>
Maria, Caroline, Veronika</h2>

<h2>Oppgave 1:</h2>

| Binære tall          | Hexadesimaltall | Desimaltall |
| -------------        |:-------------:  | -----:      |
| 1101                 | 0xD             | 13          |
| 110111101010         | 0xDEA           | 3562        |
| 1010111100110100     | 0xAF34          | 44852       |
| 1111111111111111     | 0xFFFF          | 65535       |
| 00010001011110001010 | 0x1178A         | 71562       |


<h2>Oppgave 1A:</h2>
<p>Fra binære tall til hexadesimale tall: Del binærtall opp i sekvenser på 4 bits, legg evt på 1-3 0er foran dersom det ikke går opp på 4. 
4 bits blir 1 hexadesimalsiffer. <br> 
Første siffer til venstre i 4-bit sekvensen ganges med 8, nummer to ganges med 4, nummer tre ganges med 2 og det siste ganges med 1. 
Summen av de 4 tallene man sitter igjen med blir hexadesimalsifferet (0-F).</p>
<p>Fra hexadesimale tall tl binær: 
Konverter hvert hexadesimalsiffer til 4 binærsifre.
</p>
<h2>Oppgave 1B:</h2>
<p>Fra hexadesimale tall til desimaltall: 
Start ned sifferet til høyre og gang dette med 16^0. Neste siffer (mot venstre) ganges med 16^1, osv. (Skriv ned, ikke regn ut)
Konverter alfabetiske siffer til desimalsiffer (fra A=10 til F=15), Når alt står med desimaltall, regn ut hver for seg og summer.</p> 
<p>Fra desimaltall tall til hexadesimale: Del desimaltallet på 16. 
Ignorer alle siffer etter kommategnet på kvotienten, så du sitter igjen med et heltall. 
Gang dette tallet med 16, og trekk det fra det opprinnelige desimaltallet. Differansen (rest) konverteres til hexadesimaltall. Dette er siste hexadesimalsiffer i hexadesimaltallet.
Gjenta ved å dele kvotienten fra første divisjon (heltallet) på 16, og følg samme prosedyre. Neste rest er nest siste tall i hexadesimaltallet, osv. 
Fortsett til tallet du skal dele på 16 er lavere enn 16, konverter dette til hexadesimaltall. Dette er første hexadesimalsiffer i hexadesiamltallet.


<h2>Oppgave 2C:</h2>

Test: 9 total, 9 passed

| Navn                          | Antall            | Resultat          |
| -------------                 |:-------------:    | ---------:        |
| BenchmarkBSortModified100     | 30000	            | 59599 ns/op       |
| BenchmarkBSortModified1000    | 500	            | 3248392 ns/op     |
| BenchmarkBSortModified10000   | 3                 | 448335263800 ns/op|
| BenchmarkBSort100             | 50000	            | 37288 ns/op       |
| BenchmarkBSort1000            | 1000	            | 2156661 ns/op     |
| BenchmarkBSort10000           | 5                 | 2156661 ns/op     |
| BenchmarkQSort100             | 300000            | 8554 ns/op        |
| BenchmarkQSort1000            | 20000	            | 73936 ns/op       |
| BenchmarkQSort10000           | 2000	            | 828352 ns/op      |