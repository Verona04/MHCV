<h1>Obligatorisk Innlevering 4<br>
11. mai 2018</h1>
<h2>Gruppe MCV <br>
Maria Gina Nessler, Caroline Andersen, Veronika Andersen</h2>

<h2>Spesifikasjon</h2>
<p>
Vi har lagd en oversikt over alle parkeringsplasser i Norge. Oversikten
viser navn, type parkeringsplass, adresse, antall avgiftsfrie plasser,
antall avgiftsbelagte plasser, antall ladeplasser, antall parkeringsplasser
for forflytningshemmede, og informasjon for forflytningshemmede.</p>
<p>
Man kan søke, enten på poststed eller navn på parkeringsplassen, og
resultat vises både som markører på kart og, i en liste med all
informasjon.</p><p>
Resultatlisten har også link til kartet, slik at kartet sentreres på
koordinatene til valgt parkeringsplass når man trykker på linken.
Trykker man på en parkeringsplass på kartet, scrolles resultatlista
til riktig sted, og highlighter den valgte parkeringsplassen.</p>
<p>
Vi har også en søkefunksjon som søker basert på området som vises på
kartet. Vi tar utgangspunkt i koordinatene midt på kartet, og søker i
en radius på 10 km. Resultatlista under vil da sorteres på avstand
(i luftlinje) fra senter av kartet.
Søket inneholder også to filter: «Vis kun parkeringsanlegg tilpasset
forflytningshemmede», og «Vis kun parkeringsanlegg med ladestasjon for
elbil». Dette for at man enkelt skal kunne finne parkering tilpasset
spesielle behov.</p>
<p>Søket er begrenset til 150 treff.</p>

<h2>Arkitektur</h2>