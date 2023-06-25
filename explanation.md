### Auswertung der Ergebnisse:

#### *Ziel der Aufgabe:*
- Ermitteln Sie die Anzahl der sich schneidenden Stecken
- Messen Sie die pro Datei aufgewendete Zeit

#### *Erläuterung des Ansatzes:*
Um die Anzahl der sich schneidenden Strecken herauszufinden, wird zunächst die Funktion
`amountOfInterceptingGraphs(graphs []Graph) int {}` aufgerufen. Der Parameter vom Typ `[] Graph` besteht aus zwei 
Punkten, welche eine Strecke aufstellen. In dieser Funktion werden zwei verschachtelte 
Schleifen durchgelaufen. Die äußere Schleife läuft von 0 bis zur Länge der Liste der 
Graphen. Die innere Schleife läuft von der aktuellen äußeren Schleifenvariabele bis zur 
Länge der Liste der Graphen. Mit diesen Schleifen wird sichergestellt, dass jedes 
Graphenpaar nur einmal überprüft wird und nicht in umgekehrter Reihenfolge erneut überprüft 
wird. Die Bedingung `i != j` stellt sicher, dass der Graph nicht mit sich selbst überprüft wird.

In der inneren Schleife wird die Funktion `areIntercepting` aufgerufen und die beiden Graphen als 
Parameter übergeben. In dieser Funktion werden zunächst die Punkte p1 und p2 bestimmt, welche den 
ersten Graphen repräsentieren. Und für den zweiten Graphen die Punkte q1 und q2. Um zu überprüfen, 
ob sich beide Graphen schneiden, wird zunächst überprüft, ob beiden Graphen kollinear sind. Dies 
wird mithilfe der Hilfsfunktion `ccw` erreicht. Gibt diese Funktion 0 zurück, liegen die Punkte 
p1, p2 und q1 auf einer Geraden und die Punkte p1, p2 und q2 auf einer Geraden. Mit Hilfe der 
Hilfsfunktion `isPointOnLine` wird überprüft, ob entweder der Punkt q1 oder q2 auf der Strecke 
p1p2 liegt. In diesem Fall gibt die Funktion `true` zurück, da sich beide Graphen schneiden. 
Sind die Graphen kollinear, wird `ccw` erneut aufgerufen, um die Orientierung von p1, p2 und q1 
(bzw. p1, p2, q2) zu bestimmen. Wenn die beiden Orientierungen unterschiedliche Vorzeichen haben 
und die Orientierung von q1, q2 und p1 sowie die Orientierung von q1, q2 und p2 ebenfalls 
unterschiedliche Vorzeichen haben, dann schneiden sich die Graphen. In diesem Fall wird erneut 
ein `true` zurückgegeben. Wenn keins der obigen Bedingungen erfüllt ist, wird ein `false` 
zurückgegeben, da sich die Strecken nicht schneiden bzw. überlappen.

Die Hilfsfunktion `isPointOnLine` überprüft, ob ein Punkt q auf einer Linie liegt. Hierfür 
wird überprüft, ob die x-Koordinate von q zwischen den x-Koordinaten von p1 und p2 liegt. 
Liegt der Punkt auf der Strecke p1p2 wird `true` zurückgegeben, ansonsten `false`.

In der Hilfsfunktion `ccw` wird die Orientierung der drei Punkte p1, p2 und p3 mittels des Kreuzprodukts berechnet. Das Vorzeichen des Ergebnisses des Kreuzprodukts gibt die relative Position der Punkte p1, p2 und p3 an.
-	Wenn das Ergebnis negativ ist, liegen die Punkte p1, p2 und p3 *im Uhrzeigersinn*
-	Wenn das Ergebnis positiv ist, liegen die Punkte p1, p2 und p3 *gegen den Uhrzeigersinn*
-	Wenn das Ergebnis gleich null ist, sind die Punkte p1, p2 und p3 *kollinear* (auf einer Geraden)

#### *Fazit:*
Unsere Ergebnisse für alle drei Datensätze sind in der folgenden Tabelle dargestellt:

| Datensatz | Streckenanzahl | Schnittpunkte | Zeitaufwand |
|-----------|----------------|------------|-------------|
| 1         | 1001           | 11         | 1.643 ms    |
| 2         | 10001       | 733        | 173.304 ms  |
| 3         | 100001       | 77138     | 16.476 s    |

