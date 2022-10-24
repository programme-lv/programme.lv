# deikstra
Moderna programmēšanas izglītības platforma.

## Mērķis

Projekta “Deikstra” mērķis ir dot iespēju ikkatram patstāvīgi apgūt algoritmizācijas un programmēšanas pamatus. Projekts sniedz lietotājam atgriezenisko saiti, testējot iesūtītos risinājumus ar noteiktiem laika, atmiņas ierobežojumiem.

### Iesūtīšana

Standarta iesūtījumam jāatbilst https://github.com/imachug/problem-xml-specs norādītajai struktūrai.

### Ierobežojumi

Izpildot lietotāja iesūtījumu uz katra testa jāierobežo:

- *CPU laiks* - laiks, ko procesors ir pavadījis iesūtījuma izpildei summāri pa visiem kodoliem.
- *īstais laiks* - pulksteņa laiks, kas nepieciešams iesūtījuma izpildei.
- *fiziskā atmiņa* - fiziskā atmiņa ( [fiziskā vs virtuālā](https://stackoverflow.com/questions/14347206/what-are-the-differences-between-virtual-memory-and-physical-memory) ) atvēlēta procesam.
- *neaktīvais laiks* = *īstais laiks - CPU laiks.*

Dati par procesa izmantoto resursu daudzumu ir, protams, jāsaglabā.

Jasaglabā ir arī procesa *exit status, stderr.*

### Atgriezeniskā saite

Iespējamie atsevišķa **testa** statusi:

- *TLE* - pārsniegts izpildes laiks (Time Limit Exceeded)
- *MLE* - pārsniegts atmiņas limits (Memory Limit Exceeded)
- *OK* - atbilde ir pareiza (Accepted)
- *PT* - daļēji pareiza atbilde (Partial solution)
    - *PT* [punkti], ja jāprecizē iegūtie punkti
    - punkti ir reāls skaitlis, kas pieder intervālam [0;1e5]
- *WA* - atbilde nav pareiza (Wrong Answer)
- *RE* - izpildes kļūda (Runtime Error)
- *PE* - prezentācijas kļūda (Presentation Error)
- *ILE* - atbilde netika sagaidīta (Idleness limit exceeded)
- *IG* - noignorēts, jo tā rezultāts neko nemainītu (Ignored)
- *SV* - process veica potenciāli ļaunprātīgus pieprasījumus (Security violation)
- *CF* - kaut kas nogāja greizi (Check Failed)

Iespējamie **iesūtījuma** *statusi*:

- *IQS* - iesūtījums gaida rindā (In Queue State)
- *ICS* - programma tiek kompilēta (In Compilation State)
- *ITS* - programma tiek testēta (In Testing State)
- *CE* - neveiksmīga kompilācija (Compilation Error)
- *TLE* - pārsniegts izpildes laiks (Time Limit Exceeded)
- *MLE* - pārsniegts atmiņas limits (Memory Limit Exceeded)
- *OK* - risinājums ir pareizs (Accepted)
- *PT* - daļēji pareiza atbilde (Partial solution)
- *WA* - atbilde nav pareiza (Wrong Answer)
- *RE* - izpildes kļūda (Runtime Error)
- *PE* - prezentācijas kļūda (Presentation Error)
- *ILE* - pārsniegt neaktīvais laiks (Idleness Limit Exceeded)
- *CF* - kaut kas nogāja greizi (Check Failed)
- *SV* - process veica potenciāli ļaunprātīgus pieprasījumus (Security violation)
- *RJ* - iesūtījums tika noraidīts pirms izpildes (Rejected)
- *DQ* - iesūtījums tika diskvalificēts pēc izpildes (Disqualified)

## Uzdevumu gatavošana 

[https://quangloc99.github.io/2022/03/08/polygon-codeforces-tutorial.html](https://quangloc99.github.io/2022/03/08/polygon-codeforces-tutorial.html)

[https://github.com/ioi-2017/tps/tree/master/docs](https://github.com/ioi-2017/tps/tree/master/docs)