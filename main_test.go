package main

import (
	"os"
	"testing"
)

type Item struct {
	Pos   int `json:"pos"`
	Nr    int
	Title string `json:"title"`
	Year  int    `json:"year"`
}

func Test_prepareTweet(t *testing.T) {
	t.Parallel()
	tests := []struct {
		act  Item
		want string
	}{
		{act: Item{
			Pos:   100,
			Title: "Rozporządzenie @MF_gov_PLN z dnia 31 grudnia 2019 r. w sprawie postępowania kwalifikacyjnego w stosunku do kandydatów ubiegających się o przyjęcie do służby w Służbie Celno-Skarbowej",
			Year:  2020},
			want: "Dz.U. 2020 poz. 💯\nRozporządzenie @MF_gov_PLN z dnia 31 grudnia 2019 r. w sprawie postępowania kwalifikacyjnego w stosunku do kandydatów ubiegających się o przyjęcie do służby w Służbie Celno-Skarbowej\nhttps://dziennikustaw.gov.pl/D2020000010001.pdf",
		},
		{act: Item{
			Pos:   2,
			Title: "Rozporządzenie @MF_gov_PLN z dnia 31 grudnia 2019 r. w sprawie postępowania kwalifikacyjnego w stosunku do kandydatów ubiegających się o przyjęcie do służby w Służbie Celno-Skarbowej",
			Year:  2020},
			want: "Dz.U. 2020 poz. 2\nRozporządzenie @MF_gov_PLN z dnia 31 grudnia 2019 r. w sprawie postępowania kwalifikacyjnego w stosunku do kandydatów ubiegających się o przyjęcie do służby w Służbie Celno-Skarbowej\nhttps://dziennikustaw.gov.pl/D2020000000201.pdf",
		},
		{act: Item{
			Pos:   2,
			Title: "Oświadczenie Rządowe z dnia 18 grudnia 2019 r. w sprawie mocy obowiązującej w relacjach między Rzecząpospolitą Polską a Republiką Islandii Konwencji wielostronnej implementującej środki traktatowego prawa podatkowego mające na celu zapobieganie erozji podstawy opodatkowania i przenoszeniu zysku, sporządzonej w Paryżu dnia 24 listopada 2016 r., oraz jej zastosowania w realizacji postanowień Umowy między Rządem Rzeczypospolitej Polskiej a Rządem Republiki Islandii w sprawie unikania podwójnego opodatkowania i zapobiegania uchylaniu się od opodatkowania w zakresie podatków od dochodu i majątku, sporządzonej w Reykjaviku dnia 19 czerwca 1998 r., oraz w realizacji postanowień Protokołu między Rządem Rzeczypospolitej Polskiej a Rządem Republiki Islandii o zmianie Umowy między Rządem Rzeczypospolitej Polskiej a Rządem Republiki Islandii w sprawie unikania podwójnego opodatkowania i zapobiegania uchylaniu się od opodatkowania w zakresie podatków od dochodu i majątku, sporządzonej w Reykjaviku dnia 19 czerwca 1998 r., podpisanego w Reykjaviku dnia 16 maja 2012 r.",
			Year:  2020},
			want: "Dz.U. 2020 poz. 2\nOświadczenie Rządowe z dnia 18 grudnia 2019 r. w sprawie mocy obowiązującej w relacjach między Rzecząpospolitą Polską a Republiką Islandii Konwencji wielostronnej implementującej środki traktatowego prawa podatkowego …\nhttps://dziennikustaw.gov.pl/D2020000000201.pdf",
		},
		{act: Item{
			Pos:   2146,
			Title: "Rozporządzenie Ministra Edukacji i Nauki z dnia 1 grudnia 2020 r. zmieniające rozporządzenie w sprawie pomocy de minimis w ramach programu „Wsparcie dla czasopism naukowych”", Year: 2020},
			want: "Dz.U. 2020 poz. 2146\nRozporządzenie @MEIN_gov_PL z dnia 1 grudnia 2020 r. zmieniające rozporządzenie w sprawie pomocy de minimis w ramach programu „Wsparcie dla czasopism naukowych”\nhttps://dziennikustaw.gov.pl/D2020000214601.pdf",
		},
		{act: Item{
			Pos:   241, Nr: 41,
			Title: "Protokół w sprawie zmiany Umowy o rozliczeniach wielostronnych w rublach transferowych i o utworzeniu Międzynarodowego Banku Współpracy Gospodarczej oraz Statutu tego Banku, sporządzony w Moskwie dnia 18 grudnia 1970 r.", Year: 1973},
			want: "Dz.U. 1973 poz. 241\nProtokół w sprawie zmiany Umowy o rozliczeniach wielostronnych w rublach transferowych i o utworzeniu Międzynarodowego Banku Współpracy Gospodarczej oraz Statutu tego Banku, sporządzony w Moskwie dnia 18 grudnia 1970 r.\nhttps://dziennikustaw.gov.pl/D1973041024101.pdf",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.act.Title, func(t *testing.T) {
			t.Parallel()
			if got := prepareTweet(tt.act.Year,  tt.act.Nr, tt.act.Pos, tt.act.Title); got != tt.want {
				t.Errorf("prepareTweet() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}

func Test_trimTweet(t *testing.T) {
	t.Parallel()
	tests := []struct {
		title string
		want  string
	}{
		{
			title: "Rozporządzenie @MF_gov_PLN z dnia 31 grudnia 2019 r. w sprawie postępowania kwalifikacyjnego w stosunku do kandydatów ubiegających się o przyjęcie do służby w Służbie Celno-Skarbowej",
			want:  "Rozporządzenie @MF_gov_PLN z dnia 31 grudnia 2019 r. w sprawie postępowania kwalifikacyjnego w stosunku do kandydatów ubiegających się o przyjęcie do służby w Służbie Celno-Skarbowej",
		},
		{
			title: "Oświadczenie Rządowe z dnia 18 grudnia 2019 r. w sprawie mocy obowiązującej w relacjach między Rzecząpospolitą Polską a Republiką Islandii Konwencji wielostronnej implementującej środki traktatowego prawa podatkowego mające na celu zapobieganie erozji podstawy opodatkowania i przenoszeniu zysku, sporządzonej w Paryżu dnia 24 listopada 2016 r., oraz jej zastosowania w realizacji postanowień Umowy między Rządem Rzeczypospolitej Polskiej a Rządem Republiki Islandii w sprawie unikania podwójnego opodatkowania i zapobiegania uchylaniu się od opodatkowania w zakresie podatków od dochodu i majątku, sporządzonej w Reykjaviku dnia 19 czerwca 1998 r., oraz w realizacji postanowień Protokołu między Rządem Rzeczypospolitej Polskiej a Rządem Republiki Islandii o zmianie Umowy między Rządem Rzeczypospolitej Polskiej a Rządem Republiki Islandii w sprawie unikania podwójnego opodatkowania i zapobiegania uchylaniu się od opodatkowania w zakresie podatków od dochodu i majątku, sporządzonej w Reykjaviku dnia 19 czerwca 1998 r., podpisanego w Reykjaviku dnia 16 maja 2012 r.",
			want:  "Oświadczenie Rządowe z dnia 18 grudnia 2019 r. w sprawie mocy obowiązującej w relacjach między Rzecząpospolitą Polską a Republiką Islandii Konwencji wielostronnej implementującej środki traktatowego prawa podatkowego …",
		},
		{
			title: "Obwieszczenie Ministra Zdrowia z dnia 21 maja 2020 r. w sprawie ogłoszenia jednolitego tekstu rozporządzenia Ministra Zdrowia w sprawie grzybów dopuszczonych do obrotu lub produkcji przetworów grzybowych, środków spożywczych zawierających grzyby oraz uprawnień klasyfikatora grzybów i grzyboznawcy",
			want:  "📢Obwieszczenie @MZ_gov_PL z dnia 21 maja 2020 r. w sprawie ogłoszenia jednolitego tekstu rozporządzenia @MZ_gov_PL w sprawie grzybów dopuszczonych do obrotu lub produkcji przetworów grzybowych, środków spożywczych …",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()
			if got := trimTitle(tt.title); got != tt.want {
				t.Errorf("prepareTweet() =\n%v, want\n%v", got, tt.want)
			}
		})
	}
}

func Test_extractActFromTweet(t *testing.T) {
	t.Parallel()
	tests := []struct {
		title string
		year  int
		nr    int
		pos   int
	}{
		{
			title: "Rozporządzenie @MF_gov_PLN z dnia 31 grudnia 2019 r. w sprawie postępowania kwalifikacyjnego w stosunku do kandydatów ubiegających się o przyjęcie do służby w Służbie Celno-Skarbowej",
			year:  0, pos: 0,
		}, {
			title: "@Grzegorz_64 @LBalcerowicz Art. 14. Wprowadzenie podatku nie może stanowić podstawy do zmiany warunków świadczenia usług finansowych i ubezpieczeniowych wykonywanych na podstawie umów zawartych przed dniem wejścia w życie ustawy. (Dz.U. 2016 poz. 68) 😉",
			year:  2016, pos: 68,
		},
		{
			title: "@p_lichtarowicz @GminaPolice @police24pl \"Samorząd województwa, z własnej inicjatywy ... może występować o dofinansowanie realizacji programów rozwoju, regionalnego programu operacyjnego ... środkami budżetu państwa i środkami pochodzącymi z budżetu Unii Europejskiej...\" (art. 11 p. 4 - Dz.U.2020.1668 t.j.).",
			year:  2020, pos: 1668,
		}, {
			title: "@gucio_70 @pan04509420 @Pan__Robert @Tezakrim @copone1dak @ZiobroPL Źródło:  Dz.U.2019.2393    ",
			year:  2019, pos: 2393,
		}, {
			title: "Dz.U.2018.0.1799 t.j. - Ustawa z dnia 9 maja 1996 r. o wykonywaniu mandatu posła i senatora",
			year:  2018, pos: 1799,
		}, {
			title: "Źródło: Dz.U. z 2012 r.  poz. 318.",
			year:  2012, pos: 318,
		}, {
			title: "Dz.U.2020.0.360 t.j. - Ustawa z dnia 6 kwietnia 1990 r. o Policji",
			year:  2020, pos: 360,
		}, {
			title: "[1/2] Absurd z Dz.U. 2020 Poz. 2132",
			year:  2020, pos: 2132,
		}, {
			title: "Dziennik Ustaw Dz.U.2019.1347 t.j. dla ułatwienia Dział II - wystarczy przeczytać i sprawa będzie jasna\n",
			year:  2019, pos: 1347,
		}, {
			title: "„Art.46bb. Nieprzestrzeganie obowiązku,o którym mowa w art. 46b pkt 13,stanowi uzasadnioną przyczynę\nodmowy sprzedaży,o której mowa w art.135 ustawy z dnia 20 maja 1971 r.–Kodeks wykroczeń (Dz.U. z 2019 r. poz. 821,z późn. zm.2)”",
			year:  2019, pos: 821,
		}, {
			title: " Prawo o zgromadzeniach (Dz. U. z 2019 r. poz. 631),z wyłączeniem zgromadzeń organizowanych na podstawie zawiadomienia, o którym mowa ",
			year:  2019, pos: 631,
		}, {
			title: "Może państwo \"prawa\" zrobi coś zgodnie z \"prawem\" Dz.U.2019.821 t.j. | Akt obowiązujący",
			year:  2019, pos: 821,
		}, {
			title: "Dz.U. z 2012 r.  poz. 318. Dz.U. z 2012 r.  poz. 319.",
			year:  2012, pos: 318,
		}, {
			title: "Kolejna ważna zmiana tym razem dla 19 latków /Dz.U. poz.2157/. Wprowadzono w związku z #COVID19 zapis, że coroczną kwalifikację wojskową można realizować w kilku okresach w roku. Dotychczas kwalifikacja trwała od stycznia do kwietnia. https://t.co/qaf5ipyqJR",
			year:  0, pos: 2157,
		}, {
			title: "Źródło: Dz.U. z 1998 r. Nr 51, poz. 318.",
			year:  1998, nr: 51, pos: 318,
		}, {
			title: "Wbij sobie to do swojej pustej bani !\nDz. U. 1997.78.483 - Konstytucja Rzeczypospolitej Polskiej z dnia 2 kwietnia 1997 r. Nikt nie może być poddany eksperymentom naukowym, w tym medycznym, bez dobrowolnie wyrażonej zgody.",
			year:  1997, nr: 78, pos: 483,
		}, {
			title: "Dz.U.1997.106.681 z późn. zm.",
			year:  1997, nr: 106, pos: 681,
		}, {
			title: "w związku z Art. 47 w związku z Art. 51 w związku z Art. 52  Konstytucji Rzeczypospolitej Polskiej (Dz.U. 1997 nr 78 poz. 483) !!!",
			year:  1997, nr: 78, pos: 483,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()
			y, n, p := extractActFromTweet(tt.title)
			if y != tt.year {
				t.Errorf("extractActFromTweet() =\n%v, want\n%v", y, tt.year)
			}
			if p != tt.pos {
				t.Errorf("extractActFromTweet() =\n%v, want\n%v", p, tt.pos)
			}
			if n != tt.nr {
				t.Errorf("extractActFromTweet() =\n%v, want\n%v", n, tt.nr)
			}
		})
	}
}

func Test_getIdFromTweet(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in string
		id int
		y  int
	}{
		{in: "Dz.U. 2020 poz.1", y: 0, id: 0},
		{in: "Dz.U. abcd poz. xyz", y: 0, id: 0},
		{in: "Dz.U. 2020 poz. xyz", y: 0, id: 0},
		{in: "Dz.U. abcd poz. 1", y: 0, id: 0},
		{in: "Dz.U. 2020 poz. 1", y: 2020, id: 1},
		{in: "Dz.U. 2020 poz. 999", y: 2020, id: 999},
		{in: "Dz.U. 2020 poz. 2\nRozporządzenie @MF_gov_PLN z dnia 31 grudnia 2019 r. w sprawie postępowania kwalifikacyjnego w stosunku do kandydatów ubiegających się o przyjęcie do służby w Służbie Celno-Skarbowej\nhttp://api.sejm.gov.pl/eli/acts/DU/2020/2/text.pdf", y: 2020, id: 2},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()
			y, id := getIdFromTweet(tt.in)
			if y != tt.y {
				t.Errorf("getIdFromTweet() = %v, want %v", y, tt.y)
			}
			if id != tt.id {
				t.Errorf("getIdFromTweet() = %v, want %v", id, tt.id)
			}
		})
	}
}

func Test_getTitleFromPage(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		want string
	}{
		{name: "sample.html", want: "Rozporządzenie Ministra Edukacji i Nauki z dnia 1 grudnia 2020 r. zmieniające rozporządzenie w sprawie pomocy de minimis w ramach programu „Wsparcie dla czasopism naukowych”"},
		{name: "404.html", want: ""},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			file, _ := os.Open("testdata/" + tt.name)
			if got := getTitleFromPage(file); got != tt.want {
				t.Errorf("getTitleFromPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertPDFToPng(t *testing.T) {
	t.Skip("There is an intergration test for that")
	t.Parallel()
	file, _ := os.Open("testdata/D2020000000101.pdf")
	out, err := convertPDFToJpgs(file)
	if err != nil {
		t.Errorf("Got %v", err)
	}
	if len(out) != 2 {
		t.Errorf("Got %v", out)
	}
}

func TestIntegrationGetTweetText(t *testing.T) {
	t.Parallel()
	text := getTweetText(1997,78,483 )
	e := "Dz.U. 1997 poz. 483\nKonstytucja Rzeczypospolitej Polskiej z dnia 2 kwietnia 1997 r. uchwalona przez Zgromadzenie Narodowe w dniu 2 kwietnia 1997 r., przyjęta przez Naród w referendum konstytucyjnym w dniu 25 maja 1997 r., podpisana przez …\nhttps://dziennikustaw.gov.pl/D1997078048301.pdf"
	if text != e {
		t.Errorf("Expected %s got %s", e, text)
	}
}

func TestIntegrationgetPDF(t *testing.T) {
	t.Parallel()
	r, err := getPDF(2020,0,2267 )
	if err != nil {
		t.Fatal(err)
	}
	pages, err := convertPDFToJpgs(r.Body)
	if len(pages) != 1 {
		t.Errorf("expected pages %d but got %d", 1, len(pages))
	}
	if err != nil {
		t.Error(err)
	}
}

