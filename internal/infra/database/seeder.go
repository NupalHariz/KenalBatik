package database

import (
	"kenalbatik-be/internal/domain"

	"gorm.io/gorm"
)

func SeedData(db *gorm.DB) {
	var totalProvinces int64

	if err := db.Model(domain.Province{}).Count(&totalProvinces).Error; err != nil {
		panic(err)
	}

	if totalProvinces == 0 {
		generateProvince(db)
	}

	var totalIslands int64

	if err := db.Model(domain.Island{}).Count(&totalIslands).Error; err != nil {
		panic(err)
	}

	if totalIslands == 0 {
		generateIsland(db)
	}

	var totalBatiks int64

	if err := db.Model(domain.Batik{}).Count(&totalBatiks).Error; err != nil {
		panic(err)
	}

	if totalBatiks == 0 {
		generateBatik(db)
	}

	var totalQuizzes int64

	if err := db.Model(domain.Quiz{}).Count(&totalQuizzes).Error; err != nil {
		panic(err)
	}

	if totalQuizzes == 0 {
		generateQuiz(db)
	}

}

func generateBatik(db *gorm.DB) {
	batikSeeder := &[]domain.Batik{
		//JAWA
		{
			Name:        "Batik Mega Mendung",
			Description: "Batik Mega Mendung adalah salah satu motif batik yang ikonik dari Cirebon, Jawa Barat. Motifnya terinspirasi dari bentuk awan mendung dengan garis-garis yang melengkung dan warna-warna yang kontras, seperti biru, merah, dan ungu. Filosofi dari Batik Mega Mendung adalah tentang ketenangan, kesabaran, dan pengendalian diri dalam menghadapi berbagai situasi hidup. Seperti awan yang tenang sebelum datangnya hujan, manusia diharapkan bisa bersikap tenang dalam menghadapi emosi dan cobaan hidup. Selain itu, motif Mega Mendung juga mencerminkan akulturasi budaya antara masyarakat Cirebon dan budaya Tiongkok, yang masuk melalui perdagangan pada masa lalu. Awan melambangkan kebebasan dan harapan, sedangkan warna-warna cerah menunjukkan keceriaan dan semangat hidup. Dengan filosofi yang kaya ini, Batik Mega Mendung menjadi simbol penting dalam budaya dan sejarah masyarakat Cirebon, serta salah satu warisan budaya Indonesia yang diakui dunia.",
			City:        "Cirebon",
			ProvinceID:  12, // Jawa Barat
			IslandID:    1,  // Jawa
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Jawa/MegaMendung.png?t=2024-10-14T08%3A25%3A03.717Z",
		},
		{
			Name:        "Batik Pekalongan",
			Description: "Batik Pekalongan dikenal dengan motif-motif yang kaya akan warna dan lebih bebas dalam bentuknya, dibandingkan dengan batik dari daerah lain di Jawa. Batik ini sering kali menampilkan motif flora, fauna, dan kehidupan sehari-hari yang dipengaruhi oleh budaya Tiongkok, Eropa, Arab, dan Melayu. Batik Pekalongan adalah contoh dari kemampuan masyarakat untuk beradaptasi dengan pengaruh luar tanpa meninggalkan identitas lokal. Filosofinya adalah keterbukaan terhadap perubahan dan keberagaman, namun tetap berpegang pada akar tradisional. Keterampilan tangan yang diturunkan dari generasi ke generasi juga menjadi simbol penting dari penghargaan terhadap tradisi.",
			City:        "Pekalongan",
			ProvinceID:  14, // Jawa Tengah
			IslandID:    1,  // Jawa
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Jawa/Pekalongan.webp?t=2024-10-14T07%3A54%3A55.019Z",
		},
		{
			Name:        "Batik Sogan",
			Description: "Batik Sogan merupakan batik tradisional yang berasal dari Yogyakarta dan Surakarta. Batik ini biasanya berwarna coklat dengan pola-pola klasik seperti motif parang, kawung, dan lainnya. Warna coklat sogan yang khas menciptakan kesan elegan dan mewah. Filosofi dari Batik Sogan adalah tentang kebijaksanaan, ketenangan, serta penghormatan terhadap nilai-nilai luhur dalam kehidupan. Pola dan motifnya sering dikaitkan dengan simbol kebangsawanan, karena dahulu batik jenis ini hanya dikenakan oleh kalangan keraton. Saat ini, Batik Sogan tidak hanya menjadi kebanggaan masyarakat Jawa, namun juga menjadi salah satu simbol budaya Indonesia yang dikenal hingga mancanegara.",
			City:        "Yogyakarta",
			ProvinceID:  15, // DI Yogyakarta
			IslandID:    1,  // Jawa
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Jawa/Sogan.jpg?t=2024-10-14T08%3A24%3A36.097Z",
		},
		{
			Name:        "Batik Gentongan",
			Description: "Batik Gentongan berasal dari Madura dan dikenal dengan proses pembuatannya yang unik menggunakan gentong atau tong sebagai tempat merendam kain dalam pewarnaan. Motif Batik Gentongan biasanya mencerminkan kehidupan sehari-hari masyarakat Madura, seperti flora, fauna, dan simbol-simbol tradisional lainnya. Warna-warna cerah seperti merah, biru, dan hijau mendominasi batik ini, memberikan kesan ceria dan penuh semangat. Filosofi dari Batik Gentongan adalah tentang ketekunan dan kesabaran dalam menjalani kehidupan, mencerminkan proses pembuatannya yang membutuhkan waktu lama dan ketelitian tinggi. Batik ini juga menunjukkan hubungan harmonis antara manusia dan alam yang telah menjadi bagian dari budaya Madura sejak lama.",
			City:        "Bangkalan",
			ProvinceID:  16, // Jawa Timur
			IslandID:    1,  // Jawa
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Jawa/Gentongan.jpeg?t=2024-10-14T08%3A24%3A18.842Z",
		},
		{
			Name:        "Batik Kraton",
			Description: "Batik Kraton adalah salah satu batik klasik yang berasal dari lingkungan keraton di Yogyakarta dan Surakarta. Motif-motifnya sangat kaya akan makna filosofis dan simbol-simbol yang berkaitan dengan kebijaksanaan, keberanian, dan kesucian. Dulu, batik ini hanya dikenakan oleh kalangan keraton atau bangsawan sebagai simbol status sosial. Motif-motif seperti parang, kawung, dan meru merupakan yang paling terkenal dari Batik Kraton. Filosofi dari Batik Kraton menekankan pentingnya keselarasan antara manusia dengan alam dan nilai-nilai spiritual. Setiap motif yang dihasilkan memiliki makna mendalam, seperti harapan akan kehidupan yang seimbang, harmoni, serta ketentraman batin. Batik ini menjadi warisan budaya yang sangat dihargai dan dilestarikan hingga sekarang.",
			City:        "Yogyakarta",
			ProvinceID:  15, // DI Yogyakarta
			IslandID:    1,  // Jawa
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Jawa/Kraton.png?t=2024-10-14T08%3A24%3A03.927Z",
		},
		{
			Name:        "Batik Simbut",
			Description: "Batik Simbut berasal dari suku Baduy di Banten dan memiliki motif yang sederhana namun penuh makna. Motif Simbut menggambarkan bentuk dedaunan atau sulur-suluran yang mencerminkan kesederhanaan hidup masyarakat Baduy yang menyatu dengan alam. Batik ini dibuat dengan pewarna alami, menggunakan teknik yang diwariskan dari generasi ke generasi. Filosofi dari Batik Simbut adalah tentang kesederhanaan, ketulusan, serta kebersamaan dengan alam. Masyarakat Baduy percaya bahwa hidup yang selaras dengan alam akan membawa kedamaian dan kesejahteraan. Batik Simbut menjadi simbol kekayaan budaya yang menjaga kearifan lokal dan identitas komunitas adat Baduy.",
			City:        "Lebak",
			ProvinceID:  13, // Banten
			IslandID:    1,  // Jawa
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Jawa/Simbut.jpeg?t=2024-10-14T08%3A23%3A41.210Z",
		},
		{
			Name:        "Batik Parang",
			Description: "Batik Parang merupakan salah satu motif batik tertua di Indonesia, dengan motif menyerupai ombak yang berulang dan diagonal. Asal-usulnya dari Keraton Yogyakarta dan Surakarta, dan motif ini menggambarkan semangat yang tidak pernah padam, kekuatan, serta perjuangan yang terus-menerus. Filosofinya adalah semangat pantang menyerah dalam menghadapi kehidupan, seperti ombak yang terus bergulir tanpa henti. Batik Parang juga melambangkan hubungan yang erat antara manusia dengan Sang Pencipta. Dulu, motif ini hanya boleh dipakai oleh keluarga kerajaan sebagai simbol kebijaksanaan dan kekuatan. Kini, Batik Parang menjadi salah satu simbol kebanggaan budaya Indonesia yang dikenal luas.",
			City:        "Surakarta",
			ProvinceID:  14, // Jawa Tengah
			IslandID:    1,  // Jawa
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Jawa/Parang.jpeg?t=2024-10-14T08%3A23%3A32.148Z",
		},
		{
			Name:        "Batik Kawung",
			Description: "Batik Kawung adalah motif batik yang terkenal dari Yogyakarta dan Jawa Tengah, dengan pola geometris yang menyerupai irisan buah kawung atau kolang-kaling. Motif ini dikenal dengan keteraturannya yang mencerminkan nilai-nilai kesederhanaan, kejujuran, serta keikhlasan. Batik Kawung sering dikaitkan dengan makna spiritualitas dan kebijaksanaan, yang menggambarkan kesucian dan kemurnian hati. Dulu, motif Kawung sering digunakan oleh raja atau pejabat tinggi sebagai simbol kebangsawanan dan keagungan. Filosofi dari Batik Kawung mengajarkan tentang keseimbangan antara dunia fisik dan spiritual, serta harapan untuk hidup yang sejahtera dan harmonis. Batik ini tetap menjadi salah satu motif yang digemari karena keindahan dan makna yang terkandung di dalamnya.",
			City:        "Yogyakarta",
			ProvinceID:  15, // DI Yogyakarta
			IslandID:    1,  // Jawa
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Jawa/Kawung.jpg?t=2024-10-14T08%3A23%3A01.443Z",
		},
		{
			Name:        "Batik Geblek Renteng",
			Description: "Batik Geblek Renteng adalah batik khas dari Kulon Progo, DI Yogyakarta. Motifnya terinspirasi dari makanan tradisional setempat yang disebut geblek, yakni sejenis camilan berbahan dasar singkong. Motif ini memiliki pola berbentuk lingkaran yang saling terhubung, menyerupai bentuk geblek yang berjejer. Filosofi dari Batik Geblek Renteng adalah tentang kebersamaan, gotong royong, dan kekompakan dalam masyarakat. Motif ini juga menggambarkan nilai-nilai tradisional yang masih dijaga oleh masyarakat Kulon Progo, serta semangat untuk melestarikan budaya lokal. Warna-warna batik ini umumnya lembut dan natural, mencerminkan kesederhanaan serta kehangatan dari kehidupan masyarakat Kulon Progo.",
			City:        "Kulon Progo",
			ProvinceID:  15, // DI Yogyakarta
			IslandID:    1,  // Jawa
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Jawa/Geblek%20Renteng.jpg?t=2024-10-14T08%3A22%3A47.556Z",
		},
		//Sumatra
		{
			Name:        "Batik Tanah Liek",
			Description: "Batik Tanah Liek adalah batik khas dari Sumatera Barat yang menggunakan tanah liat sebagai bahan pewarna alami. Proses pembuatannya yang unik ini menciptakan warna coklat kemerahan yang khas dan alami. Batik ini sering kali menampilkan motif-motif khas Minangkabau, seperti rumah gadang, motif flora, serta ukiran tradisional. Filosofi Batik Tanah Liek mencerminkan hubungan erat antara manusia dan alam, serta keberanian untuk mempertahankan tradisi di tengah modernisasi. Motifnya sering kali mengandung makna kehidupan masyarakat Minang yang penuh kearifan lokal, serta nilai-nilai adat yang dijunjung tinggi, seperti kebersamaan dan gotong royong. Batik ini adalah simbol kekayaan budaya Sumatera Barat yang terus dilestarikan oleh masyarakat setempat.",
			City:        "Padang",
			ProvinceID:  3, // Sumatera Barat
			IslandID:    2, // Sumatra
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Sumatra/Batik%20Tanah%20Liek.jpg?t=2024-10-14T14%3A12%3A30.015Z",
		},
		{
			Name:        "Batik Kapal Sanggat",
			Description: "Batik Kapal Sanggat berasal dari Kalimantan Barat, dengan motif utama berupa kapal yang melambangkan peran penting sungai dan laut dalam kehidupan masyarakat pesisir Kalimantan. Motif ini menggambarkan perjalanan hidup yang penuh tantangan, namun tetap berlayar dengan semangat yang tinggi. Warna-warna yang digunakan biasanya gelap dengan kombinasi warna merah dan hitam, memberikan kesan yang kuat dan berani. Filosofi Batik Kapal Sanggat mengajarkan tentang keteguhan dalam menghadapi arus kehidupan dan pentingnya menjalin hubungan baik dengan alam. Batik ini menjadi salah satu warisan budaya yang menggambarkan identitas masyarakat pesisir Kalimantan Barat dan mencerminkan semangat mereka dalam menghadapi tantangan hidup di sepanjang perairan sungai dan laut.",
			City:        "Pontianak",
			ProvinceID:  6, // Sumatera Selatan
			IslandID:    2, // Sumatra
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Sumatra/Kapal%20Sanggat.jpg?t=2024-10-14T14%3A12%3A39.588Z",
		},
		{
			Name:        "Batik Keluak Daun Pakis",
			Description: "Batik Keluak Daun Pakis merupakan batik khas dari Papua yang mengambil inspirasi dari flora khas Papua, yaitu daun pakis. Motif ini menggambarkan daun pakis yang tumbuh subur di alam Papua, dengan sentuhan warna alami yang didapat dari bahan pewarna tradisional seperti keluak. Batik ini mencerminkan harmoni antara masyarakat Papua dan alam sekitarnya, serta nilai-nilai tradisi yang diwariskan secara turun temurun. Filosofinya adalah tentang kehidupan yang tumbuh dan berkembang dalam lingkungan yang penuh tantangan namun tetap mempertahankan keaslian. Batik Keluak Daun Pakis juga mencerminkan rasa syukur masyarakat Papua terhadap alam yang memberikan banyak sumber kehidupan. Motif ini menjadi simbol kebanggaan dan identitas lokal yang unik.",
			City:        "Jayapura",
			ProvinceID:  3, // Sumatera Barat
			IslandID:    2, // Sumatra
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Sumatra/Keluak%20Daun%20Pakis.jpeg?t=2024-10-14T14%3A12%3A49.577Z",
		},
		{
			Name:        "Batik Rangkiang",
			Description: "Batik Rangkiang adalah batik khas dari Sumatera Barat yang terinspirasi dari rangkiang, yaitu lumbung padi tradisional Minangkabau. Motif ini melambangkan kemakmuran, kesuburan, dan kerja keras masyarakat dalam mengolah hasil bumi. Batik ini sering kali menampilkan bentuk geometris yang menyerupai rangkiang dengan pola-pola yang rapi dan simetris. Filosofi Batik Rangkiang adalah tentang pentingnya menyimpan hasil bumi sebagai persiapan untuk masa depan, serta kebersamaan dalam bekerja untuk kesejahteraan bersama. Warna-warna batik ini cenderung hangat dengan dominasi warna coklat, merah, dan hitam yang menunjukkan karakter kuat masyarakat Minangkabau. Batik Rangkiang menjadi salah satu simbol budaya yang mencerminkan identitas lokal Sumatera Barat yang kaya.",
			City:        "Bukittinggi",
			ProvinceID:  3, // Sumatera Barat
			IslandID:    2, // Sumatra
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Sumatra/Rangkiang.jpeg?t=2024-10-14T14%3A12%3A57.191Z",
		},
		{
			Name:        "Batik Pucuak Rabuang",
			Description: "Batik Pucuak Rabuang berasal dari Sumatera Barat, dengan motif utama berupa pucuak rabuang atau pucuk rebung. Motif ini menggambarkan pertumbuhan dan harapan, di mana pucuk rebung melambangkan regenerasi dan kelanjutan kehidupan. Batik ini mengajarkan tentang pentingnya menjaga nilai-nilai tradisi sambil tetap terbuka terhadap perkembangan zaman. Filosofinya adalah semangat untuk terus tumbuh dan berkembang seperti pucuk rebung yang selalu menuju ke atas. Batik Pucuak Rabuang sering kali menggunakan warna-warna natural yang mencerminkan alam Sumatera Barat, seperti hijau, coklat, dan merah. Batik ini adalah salah satu simbol budaya Minangkabau yang kaya akan nilai-nilai filosofis dan estetika.",
			City:        "Padang",
			ProvinceID:  3, // Sumatera Barat
			IslandID:    2, // Sumatra
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Sumatra/Pucuak%20Rabuang.png?t=2024-10-14T14%3A13%3A10.940Z",
		},
		{
			Name:        "Batik Cual",
			Description: "Batik Cual adalah batik tradisional dari Bangka Belitung yang memiliki keunikan dalam penggunaan benang emas dan motif-motif yang rumit. Batik ini sering kali menampilkan pola-pola yang terinspirasi dari alam laut dan kehidupan masyarakat pesisir, seperti motif ikan, bunga, dan ombak. Penggunaan benang emas memberikan kesan mewah dan elegan pada batik ini, menjadikannya simbol status dan kemewahan pada zaman dahulu. Filosofi dari Batik Cual adalah tentang keberanian dalam mengejar keindahan dan keagungan, serta apresiasi terhadap kekayaan alam laut yang menjadi bagian penting dari kehidupan masyarakat Bangka Belitung. Batik ini kini menjadi salah satu produk unggulan yang mengangkat budaya lokal ke kancah nasional.",
			City:        "Pangkal Pinang",
			ProvinceID:  8, // Bangka Belitung
			IslandID:    2, // Sumatra
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Sumatra/Cual.jpg?t=2024-10-14T14%3A13%3A30.044Z",
		},
		{
			Name:        "Batik Simeol-meol",
			Description: "Batik Simeol-meol merupakan batik tradisional yang berasal dari Kepulauan Riau, dengan motif-motif yang menggambarkan kehidupan laut dan pesisir. Batik ini menampilkan pola-pola seperti ombak, ikan, dan karang yang menggambarkan kekayaan alam bawah laut di wilayah Kepulauan Riau. Filosofi dari Batik Simeol-meol adalah tentang kehidupan yang dinamis dan penuh perubahan, seperti ombak yang tidak pernah berhenti. Motif ini mengajarkan tentang pentingnya adaptasi dan keberanian dalam menghadapi perubahan hidup. Warna-warna yang digunakan biasanya terinspirasi dari laut, seperti biru, hijau, dan putih. Batik ini menjadi simbol kekayaan budaya dan alam maritim yang sangat erat dengan kehidupan masyarakat Kepulauan Riau.",
			City:        "Tanjung Pinang",
			ProvinceID:  5, // Kepulauan Riau
			IslandID:    2, // Sumatra
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Sumatra/Simeol-meol.jpeg?t=2024-10-14T14%3A13%3A38.677Z",
		},
		{
			Name:        "Batik Pangeret-ngeret",
			Description: "Batik Pangeret-ngeret berasal dari Kalimantan Tengah dan dikenal dengan motif-motif yang terinspirasi dari alam dan kehidupan sehari-hari masyarakat Dayak. Motifnya sering kali menampilkan unsur-unsur seperti burung enggang, tanaman, dan pola geometris yang khas. Filosofi dari Batik Pangeret-ngeret adalah tentang keseimbangan antara manusia dan alam, serta keberanian dalam menjaga dan melestarikan lingkungan. Batik ini mengajarkan tentang pentingnya hidup selaras dengan alam, serta menghormati warisan nenek moyang yang penuh dengan kearifan lokal. Warna-warna yang digunakan biasanya berasal dari pewarna alami, yang mencerminkan kesederhanaan dan keindahan alam Kalimantan Tengah.",
			City:        "Palangka Raya",
			ProvinceID:  3, // Sumatera Barat
			IslandID:    2, // Sumatra
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Sumatra/Pangeret-ngeret.jpg?t=2024-10-14T14%3A13%3A47.359Z",
		},
		{
			Name:        "Batik Tapak Raja Sulaiman",
			Description: "Batik Tapak Raja Sulaiman adalah batik khas dari Maluku yang memiliki motif berbentuk telapak kaki yang melambangkan jejak sejarah dan keberanian. Batik ini sering digunakan dalam acara-acara adat dan upacara penting sebagai simbol penghormatan terhadap leluhur dan raja-raja terdahulu. Warna-warna yang digunakan cenderung cerah dan kontras, seperti merah, kuning, dan hitam, yang menggambarkan semangat yang kuat. Filosofi Batik Tapak Raja Sulaiman adalah tentang keberanian dan tekad dalam menghadapi tantangan, serta pentingnya menghormati sejarah dan warisan budaya. Batik ini menjadi simbol kekayaan tradisi Maluku yang penuh makna, dan merupakan salah satu bentuk penghargaan terhadap budaya lokal yang lestari.",
			City:        "Ambon",
			ProvinceID:  1, // Aceh
			IslandID:    2, // Sumatra
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Sumatra/Tapak%20Raja%20Sulaiman.jpg?t=2024-10-14T14%3A13%3A54.714Z",
		},
		{
			Name:        "Batik Marsihutuan",
			Description: "Batik Marsihutuan adalah batik tradisional dari Sumatera Utara yang terinspirasi dari kehidupan masyarakat Batak dan alam pegunungan di wilayah Tapanuli. Motifnya sering kali menampilkan pola-pola seperti daun, bunga, dan simbol-simbol yang berkaitan dengan kehidupan spiritual dan adat Batak. Filosofi dari Batik Marsihutuan adalah tentang pentingnya menjaga hubungan yang harmonis antara manusia, alam, dan roh nenek moyang. Batik ini mengajarkan tentang kearifan lokal yang telah ada sejak lama dan bagaimana nilai-nilai tersebut terus diwariskan dari generasi ke generasi. Warna-warna yang digunakan cenderung natural dan gelap, mencerminkan keaslian dan kekayaan alam Sumatera Utara. Batik ini menjadi salah satu simbol kebanggaan masyarakat Batak dan warisan budaya yang tak ternilai.",
			City:        "Medan",
			ProvinceID:  2, // Sumatera Utara
			IslandID:    2, // Sumatra
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Sumatra/Marsihutuan.jpg?t=2024-10-14T14%3A14%3A03.047Z",
		},
		//Kalimantan
		{
			Name:        "Batik Bayam Raja",
			Description: "Batik Bayam Raja adalah batik khas dari Kalimantan Selatan yang terkenal dengan motif bayam raja yang menggambarkan tumbuhan bayam sebagai simbol kekuatan dan kesuburan. Motif ini melambangkan keseimbangan alam dan kehidupan yang selaras dengan hutan Kalimantan yang luas. Batik ini mengajarkan filosofi tentang kelestarian alam dan pentingnya menjaga keseimbangan antara manusia dan lingkungan. Warna-warna yang digunakan sering kali mencerminkan warna alami dari hutan Kalimantan, seperti hijau dan cokelat, memberikan kesan yang natural dan elegan. Batik Bayam Raja sering dipakai dalam berbagai upacara adat dan acara formal sebagai bentuk penghormatan terhadap tradisi leluhur yang kuat.",
			City:        "Banjarmasin",
			ProvinceID:  22, // Kalimantan Selatan
			IslandID:    3,  // Kalimantan
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Kalimantan/bayam%20raja.webp",
		},
		{
			Name:        "Batik Tampuk Manggis",
			Description: "Batik Tampuk Manggis merupakan batik dari Kalimantan Barat yang terinspirasi dari bentuk buah manggis yang melambangkan kekayaan dan kemakmuran. Motif tampuk manggis di batik ini sering kali disajikan dalam pola simetris yang indah, mencerminkan keindahan alam Kalimantan Barat. Filosofi dari Batik Tampuk Manggis adalah tentang kesabaran dan ketekunan dalam menjalani kehidupan, seperti halnya proses menanam dan merawat manggis. Batik ini sering digunakan dalam upacara adat dan perayaan budaya sebagai simbol kesuburan dan kesejahteraan masyarakat Kalimantan Barat.",
			City:        "Pontianak",
			ProvinceID:  20, // Kalimantan Barat
			IslandID:    3,  // Kalimantan
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Kalimantan/Tampuk%20Manggis.jpg?t=2024-10-14T14%3A34%3A09.251Z",
		},
		{
			Name:        "Batik Kangkung Kaombakan",
			Description: "Batik Kangkung Kaombakan adalah batik khas Kalimantan Tengah yang menampilkan motif tanaman kangkung yang bergelombang, menggambarkan kehidupan masyarakat pesisir dan aliran sungai yang menjadi sumber kehidupan di Kalimantan. Filosofi dari Batik Kangkung Kaombakan adalah tentang fleksibilitas dan kemampuan beradaptasi dengan perubahan, seperti aliran air sungai yang terus mengalir. Warna-warna yang digunakan cenderung mencerminkan alam seperti hijau, biru, dan cokelat, yang memberikan nuansa alami dan segar. Batik ini sering digunakan sebagai pakaian dalam acara-acara adat, melambangkan keterikatan dengan alam dan kehidupan yang harmonis.",
			City:        "Palangka Raya",
			ProvinceID:  21, // Kalimantan Tengah
			IslandID:    3,  // Kalimantan
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Kalimantan/KangkungKaorakan.png?t=2024-10-14T14%3A34%3A15.441Z",
		},
		{
			Name:        "Batik Benang Bintik",
			Description: "Batik Benang Bintik adalah batik dari Kalimantan Tengah yang terinspirasi dari motif benang dengan pola titik-titik kecil yang membentuk gambar alam dan kehidupan masyarakat Dayak. Batik ini mencerminkan keragaman budaya dan kekayaan tradisi suku Dayak, yang sangat menghargai alam dan lingkungan. Filosofi dari Batik Benang Bintik adalah tentang ketelitian dan keindahan dalam kesederhanaan, serta pentingnya menjaga harmoni antara manusia dan alam. Warna-warna yang digunakan sering kali memiliki nuansa tanah dan hutan, menggambarkan kedekatan dengan alam Kalimantan.",
			City:        "Palangka Raya",
			ProvinceID:  21, // Kalimantan Tengah
			IslandID:    3,  // Kalimantan
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Kalimantan/BenakBintik.jpg?t=2024-10-14T14%3A34%3A22.828Z",
		},
		{
			Name:        "Batik Awan Berarak",
			Description: "Batik Awan Berarak adalah batik tradisional dari Kalimantan Timur yang memiliki motif awan yang bergerak, melambangkan kebebasan dan harapan. Motif ini terinspirasi dari bentuk awan yang berubah-ubah dan selalu bergerak, mencerminkan semangat masyarakat Kalimantan Timur dalam menghadapi perubahan. Filosofi dari Batik Awan Berarak adalah tentang fleksibilitas dan optimisme, serta pentingnya menjaga kebersamaan dalam menghadapi tantangan. Batik ini sering dipakai dalam acara adat dan sebagai simbol dari semangat kebebasan dan keterbukaan.",
			City:        "Samarinda",
			ProvinceID:  23, // Kalimantan Timur
			IslandID:    3,  // Kalimantan
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Kalimantan/AwanBerarak.jpg?t=2024-10-14T14%3A34%3A31.411Z",
		},
		{
			Name:        "Batik Tidayu",
			Description: "Batik Tidayu adalah batik khas dari Kalimantan Barat yang menggabungkan unsur-unsur dari tiga etnis utama, yaitu Tionghoa, Dayak, dan Melayu. Motif batik ini menampilkan perpaduan antara pola-pola geometris, flora, dan simbol-simbol etnis yang mencerminkan kebersamaan dan keragaman budaya. Filosofi dari Batik Tidayu adalah tentang toleransi dan kebersamaan dalam keberagaman, serta pentingnya menghargai setiap warisan budaya yang ada. Batik ini sering dipakai dalam acara kebudayaan untuk mempererat persaudaraan antar etnis di Kalimantan Barat.",
			City:        "Pontianak",
			ProvinceID:  20, // Kalimantan Barat
			IslandID:    3,  // Kalimantan
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Kalimantan/Tidayu.jpeg?t=2024-10-14T14%3A34%3A38.603Z",
		},
		{
			Name:        "Batik Batang Garing",
			Description: "Batik Batang Garing berasal dari Kalimantan Tengah dan merupakan simbol dari kehidupan suku Dayak Ngaju. Motif Batang Garing ini menggambarkan pohon kehidupan yang melambangkan kekuatan spiritual dan keterikatan dengan alam. Batik ini sering digunakan dalam ritual adat dan upacara penting suku Dayak sebagai simbol penghormatan kepada leluhur. Filosofi dari Batik Batang Garing adalah tentang keseimbangan antara alam, manusia, dan roh-roh nenek moyang. Warna-warna yang digunakan sering kali cerah dan berani, seperti merah, kuning, dan hijau, menggambarkan semangat yang hidup dan dinamis.",
			City:        "Palangka Raya",
			ProvinceID:  21, // Kalimantan Tengah
			IslandID:    3,  // Kalimantan
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Kalimantan/BatangGaring.jpg?t=2024-10-14T14%3A34%3A45.697Z",
		},
		{
			Name:        "Batik Tarakan",
			Description: "Batik Tarakan adalah batik dari Kalimantan Utara yang memiliki motif yang menggambarkan kehidupan pesisir dan kekayaan laut. Motif ini sering kali menampilkan bentuk-bentuk ikan, ombak, dan kapal, yang mencerminkan hubungan erat masyarakat Tarakan dengan laut sebagai sumber mata pencaharian. Filosofi dari Batik Tarakan adalah tentang kebersamaan dan kerja keras dalam menjaga kelestarian laut. Batik ini mencerminkan identitas masyarakat pesisir yang kuat dan tradisi maritim yang telah ada sejak lama. Warna-warna yang digunakan mencerminkan laut seperti biru dan hijau, memberikan kesan segar dan alami.",
			City:        "Tarakan",
			ProvinceID:  24, // Kalimantan Utara
			IslandID:    3,  // Kalimantan
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Kalimantan/Tarakan.jpg?t=2024-10-14T14%3A34%3A53.929Z",
		},
		{
			Name:        "Batik Sasirangan",
			Description: "Batik Sasirangan adalah batik tradisional dari Kalimantan Selatan yang awalnya digunakan sebagai kain upacara penyembuhan suku Banjar. Motif-motif pada Batik Sasirangan memiliki makna simbolis yang kuat, seperti kekuatan, keberanian, dan keteguhan. Batik ini memiliki pola yang dibuat dengan cara mencelupkan kain yang telah diikat dengan benang, sehingga menciptakan pola unik yang berbeda-beda. Filosofi dari Batik Sasirangan adalah tentang semangat hidup dan ketahanan menghadapi berbagai cobaan. Batik ini telah menjadi identitas budaya suku Banjar dan populer sebagai pakaian dalam berbagai acara adat.",
			City:        "Banjarmasin",
			ProvinceID:  22, // Kalimantan Selatan
			IslandID:    3,  // Kalimantan
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Kalimantan/Sasirangan.jpg",
		},
		{
			Name:        "Batik Dayak Kamang",
			Description: "Batik Dayak Kamang berasal dari Kalimantan Barat dan dikenal dengan motif-motif yang terinspirasi dari seni ukir Dayak. Motifnya sering kali berupa simbol-simbol mistis, flora, dan fauna yang melambangkan kekuatan spiritual dan keterikatan dengan alam. Filosofi dari Batik Dayak Kamang adalah tentang menjaga warisan budaya dan kearifan lokal suku Dayak yang telah diwariskan secara turun temurun. Batik ini mengajarkan pentingnya melestarikan tradisi dan hidup selaras dengan alam. Warna-warna yang digunakan biasanya mencerminkan warna alami, seperti cokelat, hijau, dan merah.",
			City:        "Pontianak",
			ProvinceID:  20, // Kalimantan Barat
			IslandID:    3,  // Kalimantan
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Kalimantan/Dayak%20Kamang.jpeg",
		},
		//Sulawesi
		{
			Name:        "Batik Tari Kabasaran",
			Description: "Batik Tari Kabasaran adalah batik khas Sulawesi Utara yang terinspirasi dari tari perang tradisional Kabasaran, yang merupakan tarian adat masyarakat Minahasa. Motif batik ini menampilkan elemen-elemen seperti pakaian dan senjata yang digunakan dalam tarian tersebut, menggambarkan semangat keberanian dan kepahlawanan. Filosofi dari Batik Tari Kabasaran adalah tentang keberanian dan semangat juang dalam mempertahankan kehormatan dan tradisi. Batik ini sering dipakai dalam acara adat dan perayaan budaya sebagai simbol penghormatan terhadap para leluhur yang berani.",
			City:        "Manado",
			ProvinceID:  29, // Sulawesi Utara
			IslandID:    4,  // Sulawesi
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Sulawesi/Kabasaran.png?t=2024-10-14T14%3A50%3A52.046Z",
		},
		{
			Name:        "Batik Karawo",
			Description: "Batik Karawo berasal dari Gorontalo di Sulawesi, dikenal karena teknik pembuatannya yang unik dengan cara memotong kain secara manual untuk menciptakan pola-pola yang halus dan rumit. Batik ini sering menampilkan motif-motif alam seperti bunga, tumbuhan, dan pola geometris. Filosofi dari Batik Karawo adalah tentang ketekunan, kesabaran, dan keindahan dalam detail. Batik Karawo banyak digunakan dalam berbagai acara adat dan upacara resmi, serta menjadi kebanggaan masyarakat Gorontalo sebagai salah satu warisan budaya yang bernilai tinggi.",
			City:        "Gorontalo",
			ProvinceID:  31, // Gorontalo
			IslandID:    4,  // Sulawesi
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Sulawesi/Karawo.jpeg",
		},
		{
			Name:        "Batik Bomba",
			Description: "Batik Bomba adalah batik khas dari Sulawesi Tengah yang memiliki motif-motif terinspirasi dari budaya dan alam setempat. 'Bomba' sendiri dalam bahasa setempat berarti motif atau corak. Batik ini sering menampilkan pola-pola seperti flora, fauna, dan motif khas suku Kaili. Filosofi dari Batik Bomba adalah tentang keindahan alam Sulawesi Tengah dan pentingnya melestarikan budaya lokal. Batik ini digunakan dalam berbagai upacara adat dan kegiatan budaya sebagai bentuk penghargaan terhadap tradisi dan alam.",
			City:        "Palu",
			ProvinceID:  30, // Sulawesi Tengah
			IslandID:    4,  // Sulawesi
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Sulawesi/Bomba.jpg",
		},
		{
			Name:        "Batik Wakatobi",
			Description: "Batik Wakatobi berasal dari Sulawesi Tenggara dan terinspirasi dari keindahan laut dan terumbu karang yang ada di kepulauan Wakatobi, yang merupakan salah satu kawasan wisata bahari terbaik di Indonesia. Motif batik ini sering menampilkan bentuk-bentuk biota laut seperti ikan, karang, dan ombak, mencerminkan kekayaan laut yang dimiliki oleh Sulawesi Tenggara. Filosofi dari Batik Wakatobi adalah tentang kelestarian alam laut dan hubungan erat masyarakat dengan laut sebagai sumber kehidupan. Batik ini sering digunakan dalam acara-acara resmi dan kebudayaan untuk memperkenalkan kekayaan alam daerah Wakatobi.",
			City:        "Wakatobi",
			ProvinceID:  33, // Sulawesi Tenggara
			IslandID:    4,  // Sulawesi
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Sulawesi/Wakatobi.png?t=2024-10-14T14%3A51%3A21.619Z",
		},
		{
			Name:        "Batik La Galigo",
			Description: "Batik La Galigo adalah batik dari Sulawesi Selatan yang terinspirasi dari epos kuno 'La Galigo', yang menceritakan kisah kehidupan masyarakat Bugis di masa lampau. Motif-motif pada Batik La Galigo sering menggambarkan simbol-simbol budaya dan kehidupan masyarakat Bugis seperti perahu, padi, dan rumah adat. Filosofi dari Batik La Galigo adalah tentang kebijaksanaan dan nilai-nilai luhur yang diajarkan dalam cerita tersebut. Batik ini digunakan dalam berbagai upacara adat dan kegiatan budaya untuk melestarikan warisan sastra dan budaya Bugis.",
			City:        "Makassar",
			ProvinceID:  32, // Sulawesi Selatan
			IslandID:    4,  // Sulawesi
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Sulawesi/La%20Galigo.jpeg?t=2024-10-14T14%3A51%3A27.657Z",
		},
		{
			Name:        "Batik Jagung",
			Description: "Batik Jagung merupakan batik dari Sulawesi Selatan yang terinspirasi dari tanaman jagung, yang menjadi salah satu komoditas penting di wilayah tersebut. Motif batik ini menggambarkan bentuk jagung dan daunnya, sebagai simbol kemakmuran dan keberkahan bagi masyarakat. Filosofi dari Batik Jagung adalah tentang ketekunan dalam bercocok tanam dan pentingnya menjaga kelestarian lingkungan untuk keberlangsungan hidup. Batik ini sering digunakan dalam acara panen raya atau kegiatan kebudayaan yang terkait dengan pertanian.",
			City:        "Gowa",
			ProvinceID:  32, // Sulawesi Selatan
			IslandID:    4,  // Sulawesi
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Sulawesi/Jagung.png",
		},
		{
			Name:        "Batik Toraja",
			Description: "Batik Toraja berasal dari Sulawesi Selatan dan memiliki motif-motif yang terinspirasi dari ukiran tradisional masyarakat Toraja. Motif batik ini sering menampilkan pola-pola geometris dan bentuk-bentuk simbolis yang mencerminkan kehidupan spiritual masyarakat Toraja, seperti motif rumah Tongkonan dan motif kerbau. Filosofi dari Batik Toraja adalah tentang kehidupan setelah mati, keberanian, dan hubungan antara manusia dan leluhur. Batik ini sering digunakan dalam upacara adat Toraja sebagai bentuk penghormatan kepada para leluhur.",
			City:        "Rantepao",
			ProvinceID:  32, // Sulawesi Selatan
			IslandID:    4,  // Sulawesi
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Sulawesi/Toraja.jpg?t=2024-10-14T14%3A51%3A46.627Z",
		},
		{
			Name:        "Batik Paqbarre Allo",
			Description: "Batik Paqbarre Allo adalah batik khas Sulawesi Selatan yang terinspirasi dari alam dan kebudayaan Bugis-Makassar. Nama 'Paqbarre Allo' berarti bayangan atau cahaya matahari dalam bahasa setempat, sehingga motif-motif pada batik ini sering kali menampilkan bentuk-bentuk sinar matahari dan simbol-simbol alam lainnya. Filosofi dari Batik Paqbarre Allo adalah tentang kehangatan, optimisme, dan semangat hidup. Batik ini sering digunakan dalam kegiatan kebudayaan sebagai simbol dari harapan dan semangat yang tinggi.",
			City:        "Makassar",
			ProvinceID:  32, // Sulawesi Selatan
			IslandID:    4,  // Sulawesi
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Sulawesi/Paqbarreallo.webp?t=2024-10-14T14%3A51%3A58.790Z",
		},
		{
			Name:        "Batik Taiganja",
			Description: "Batik Taiganja merupakan batik dari Sulawesi Barat yang memiliki motif-motif yang terinspirasi dari budaya Mandar, suku asli Sulawesi Barat. Motif batik ini sering menggambarkan kehidupan masyarakat pesisir dan motif-motif alam seperti perahu dan laut, yang mencerminkan ketergantungan masyarakat Mandar pada hasil laut. Filosofi dari Batik Taiganja adalah tentang ketangguhan dalam menghadapi ombak kehidupan dan pentingnya menjaga tradisi maritim. Batik ini digunakan dalam acara-acara adat sebagai simbol dari identitas masyarakat Mandar.",
			City:        "Mamuju",
			ProvinceID:  34, // Sulawesi Barat
			IslandID:    4,  // Sulawesi
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Sulawesi/Taiganja.jpg?t=2024-10-14T14%3A52%3A08.999Z",
		},
		{
			Name:        "Batik Tongkonan",
			Description: "Batik Tongkonan adalah batik dari Sulawesi Selatan yang terinspirasi dari rumah adat Tongkonan milik masyarakat Toraja. Motif-motif pada batik ini sering menggambarkan bentuk-bentuk atap dan ukiran khas rumah Tongkonan, yang melambangkan kekuatan, perlindungan, dan kehangatan keluarga. Filosofi dari Batik Tongkonan adalah tentang ikatan keluarga dan pentingnya menjaga nilai-nilai tradisi dalam kehidupan sehari-hari. Batik ini banyak digunakan dalam upacara adat Toraja sebagai bentuk penghormatan terhadap leluhur dan simbol dari persatuan keluarga.",
			City:        "Rantepao",
			ProvinceID:  32, // Sulawesi Selatan
			IslandID:    4,  // Sulawesi
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Sulawesi/Tongkonan.jpeg?t=2024-10-14T14%3A52%3A15.734Z",
		},
		//Papua
		{
			Name:        "Batik Cendrawasih",
			Description: "Batik Cendrawasih merupakan batik khas Papua yang terinspirasi dari burung Cendrawasih, yang dikenal sebagai 'burung surga'. Motif batik ini sering menampilkan gambar burung Cendrawasih dengan warna-warna cerah dan corak yang dinamis, menggambarkan keindahan alam Papua. Filosofi dari Batik Cendrawasih adalah tentang keindahan dan kebebasan, serta simbol kebanggaan terhadap kekayaan alam Papua. Batik ini banyak digunakan dalam acara adat dan perayaan budaya sebagai wujud penghargaan terhadap keanekaragaman hayati Papua.",
			City:        "Jayapura",
			ProvinceID:  35, // Papua
			IslandID:    5,  // Papua
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Papua/Cendrawasih.jpeg?t=2024-10-14T15%3A10%3A13.096Z",
		},
		{
			Name:        "Batik Kamoro",
			Description: "Batik Kamoro berasal dari suku Kamoro di Papua dan memiliki motif-motif yang menggambarkan kehidupan masyarakat suku Kamoro, seperti ukiran kayu, perahu, dan aktivitas sehari-hari. Batik ini sering menampilkan pola-pola geometris yang mencerminkan keahlian seni ukir masyarakat Kamoro. Filosofi dari Batik Kamoro adalah tentang hubungan erat antara manusia dan alam, serta pentingnya menjaga kelestarian budaya leluhur. Batik ini digunakan dalam berbagai upacara adat sebagai simbol identitas dan warisan budaya suku Kamoro.",
			City:        "Timika",
			ProvinceID:  35, // Papua
			IslandID:    5,  // Papua
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Papua/Kamoro.jpg?t=2024-10-14T15%3A10%3A25.419Z",
		},
		{
			Name:        "Batik Prada",
			Description: "Batik Prada adalah batik khas Papua yang terkenal dengan hiasan emas atau perak yang memberikan kesan mewah dan elegan. Motif-motif batik ini sering kali terinspirasi dari budaya lokal, seperti ukiran tradisional dan simbol-simbol kepercayaan masyarakat Papua. Filosofi dari Batik Prada adalah tentang kemuliaan dan kemegahan, serta penghormatan terhadap tradisi dan adat istiadat. Batik ini sering digunakan dalam acara resmi dan upacara adat sebagai simbol kebanggaan dan status sosial.",
			City:        "Jayapura",
			ProvinceID:  35, // Papua
			IslandID:    5,  // Papua
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Papua/Prada.jpeg?t=2024-10-14T15%3A10%3A31.826Z",
		},
		{
			Name:        "Batik Asmat",
			Description: "Batik Asmat berasal dari suku Asmat di Papua dan dikenal dengan motif-motif yang terinspirasi dari ukiran kayu khas suku Asmat. Motif batik ini sering menggambarkan kehidupan masyarakat Asmat, termasuk perahu, patung leluhur, dan simbol-simbol alam. Filosofi dari Batik Asmat adalah tentang penghormatan terhadap leluhur dan pentingnya menjaga keseimbangan dengan alam. Batik ini digunakan dalam upacara adat suku Asmat sebagai bagian dari tradisi yang diwariskan dari generasi ke generasi.",
			City:        "Agats",
			ProvinceID:  35, // Papua
			IslandID:    5,  // Papua
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Papua/AsmatPapua.jpeg?t=2024-10-14T15%3A10%3A41.504Z",
		},
		{
			Name:        "Batik Tifa Honai",
			Description: "Batik Tifa Honai merupakan batik yang menggambarkan alat musik tradisional Papua, yaitu Tifa, serta rumah adat Honai. Motif batik ini sering menggambarkan bentuk Tifa yang melambangkan semangat dan kebersamaan masyarakat Papua. Filosofi dari Batik Tifa Honai adalah tentang kekuatan, kebersamaan, dan kearifan lokal yang menjadi ciri khas masyarakat Papua. Batik ini sering digunakan dalam acara-acara adat dan perayaan budaya sebagai simbol identitas dan kebanggaan terhadap tradisi Papua.",
			City:        "Jayapura",
			ProvinceID:  35, // Papua
			IslandID:    5,  // Papua
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Papua/Tifa%20Honai.jpg?t=2024-10-14T15%3A10%3A47.812Z",
		},
		{
			Name:        "Batik Sentani",
			Description: "Batik Sentani berasal dari daerah Sentani di Papua dan memiliki motif-motif yang terinspirasi dari kehidupan masyarakat di sekitar Danau Sentani. Motif-motif batik ini sering menggambarkan ikan, perahu, dan ombak, mencerminkan kehidupan masyarakat yang bergantung pada danau sebagai sumber penghidupan. Filosofi dari Batik Sentani adalah tentang kelestarian alam dan pentingnya menjaga hubungan harmonis dengan lingkungan. Batik ini digunakan dalam berbagai upacara adat sebagai bentuk penghormatan terhadap alam dan budaya setempat.",
			City:        "Sentani",
			ProvinceID:  35, // Papua
			IslandID:    5,  // Papua
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Papua/Sentani.jpeg?t=2024-10-14T15%3A10%3A55.938Z",
		},
		{
			Name:        "Batik Tifa",
			Description: "Batik Tifa adalah batik khas Papua yang menampilkan motif alat musik tradisional Tifa, yang digunakan dalam berbagai upacara adat dan kegiatan kebudayaan masyarakat Papua. Motif Tifa pada batik ini melambangkan kekuatan dan kebersamaan, serta pentingnya peran musik dalam kehidupan sosial masyarakat Papua. Filosofi dari Batik Tifa adalah tentang keharmonisan dan semangat gotong royong dalam menjaga tradisi dan adat istiadat. Batik ini sering digunakan dalam acara adat sebagai bentuk penghargaan terhadap musik dan seni tradisional Papua.",
			City:        "Jayapura",
			ProvinceID:  35, // Papua
			IslandID:    5,  // Papua
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Papua/Tifa.png?t=2024-10-14T15%3A11%3A01.519Z",
		},
		{
			Name:        "Batik Kamaro Timika",
			Description: "Batik Kamaro Timika berasal dari daerah Timika di Papua, dengan motif yang terinspirasi dari kebudayaan lokal dan kehidupan masyarakat di sekitar wilayah tersebut. Batik ini sering menampilkan pola-pola geometris yang melambangkan keseimbangan alam dan kehidupan sehari-hari. Filosofi dari Batik Kamaro Timika adalah tentang keindahan kehidupan sederhana dan harmonisasi antara manusia dan alam. Batik ini digunakan dalam berbagai acara adat dan perayaan budaya sebagai bentuk penghormatan terhadap tradisi lokal.",
			City:        "Timika",
			ProvinceID:  35, // Papua
			IslandID:    5,  // Papua
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Papua/Kamaro%20Timika.png?t=2024-10-14T15%3A10%3A19.053Z",
		},
		{
			Name:        "Batik Raja Ampat",
			Description: "Batik Raja Ampat berasal dari wilayah kepulauan Raja Ampat di Papua Barat dan dikenal dengan motif-motif yang terinspirasi dari keindahan alam bawah laut dan keragaman hayati laut di wilayah tersebut. Motif batik ini sering menggambarkan ikan, terumbu karang, dan kehidupan laut lainnya, yang menjadi daya tarik wisata utama di Raja Ampat. Filosofi dari Batik Raja Ampat adalah tentang kelestarian lingkungan laut dan pentingnya menjaga kekayaan alam untuk generasi mendatang. Batik ini digunakan dalam berbagai acara kebudayaan untuk memperkenalkan keindahan alam Papua Barat.",
			City:        "Waisai",
			ProvinceID:  36, // Papua Barat
			IslandID:    5,  // Papua
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Papua/RajaAmpat.png?t=2024-10-14T15%3A11%3A16.160Z",
		},
		{
			Name:        "Batik Ukir Sentani",
			Description: "Batik Ukir Sentani adalah batik khas Papua yang terinspirasi dari seni ukir tradisional masyarakat Sentani. Motif batik ini sering menggambarkan pola-pola ukiran yang mencerminkan kekayaan budaya dan kearifan lokal masyarakat sekitar Danau Sentani. Filosofi dari Batik Ukir Sentani adalah tentang penghormatan terhadap tradisi dan seni yang diwariskan dari generasi ke generasi. Batik ini digunakan dalam berbagai upacara adat dan acara kebudayaan sebagai simbol identitas dan kebanggaan masyarakat Sentani, serta sebagai wujud cinta terhadap budaya dan alam di sekitar mereka.",
			City:        "Sentani",
			ProvinceID:  35, // Papua
			IslandID:    5,  // Papua
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Papua/Ukir%20Sentani.jpg?t=2024-10-14T15%3A11%3A22.610Z",
		},
		//Maluku
		{
			Name:        "Batik Alam Pantai",
			Description: "Batik Alam Pantai terinspirasi dari keindahan pantai-pantai di Maluku yang menampilkan keanekaragaman hayati laut. Motifnya menggambarkan ombak, pasir, dan berbagai kehidupan laut yang menjadi bagian dari budaya dan kehidupan masyarakat pesisir. Batik ini mencerminkan hubungan harmonis antara manusia dan alam serta mengajak pemakainya untuk menjaga kelestarian lingkungan. Dengan warna-warna cerah yang mencerminkan kehidupan di laut, batik ini sering digunakan dalam acara tradisional dan perayaan sebagai simbol kebanggaan akan kekayaan alam Maluku.",
			City:        "Ambon",
			ProvinceID:  31, // Maluku
			IslandID:    6,  // Maluku
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Maluku/Alam%20Pantai.png?t=2024-10-14T15%3A24%3A57.163Z",
		},
		{
			Name:        "Batik Burung Bidadari",
			Description: "Batik Burung Bidadari menggambarkan keindahan burung-burung eksotis yang ada di Maluku. Motif ini melambangkan keberanian dan keanggunan, mencerminkan filosofi kehidupan masyarakat yang selaras dengan alam. Burung-burung ini juga menjadi simbol kebebasan dan harapan, menjadikan batik ini pilihan yang populer dalam berbagai acara formal dan informal. Dengan detail yang rumit dan warna yang cerah, batik ini menunjukkan keterampilan tangan yang tinggi dari pengrajin Maluku.",
			City:        "Ambon",
			ProvinceID:  31, // Maluku
			IslandID:    6,  // Maluku
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Maluku/BurungBidadari.jpeg?t=2024-10-14T15%3A25%3A49.049Z",
		},
		{
			Name:        "Batik Cakalele",
			Description: "Batik Cakalele terinspirasi dari tarian tradisional Cakalele yang terkenal di Maluku. Motif batik ini sering menampilkan gerakan dan keanggunan tarian, melambangkan kekuatan dan semangat juang masyarakat Maluku. Dengan warna-warna yang cerah dan dinamis, batik ini sering digunakan dalam acara-acara budaya dan festival untuk menunjukkan identitas budaya yang kaya. Cakalele menjadi simbol persatuan dan kebanggaan akan warisan budaya yang harus dilestarikan.",
			City:        "Ambon",
			ProvinceID:  31, // Maluku
			IslandID:    6,  // Maluku
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Maluku/Cakalele.jpeg?t=2024-10-14T15%3A25%3A55.932Z",
		},
		{
			Name:        "Batik Tubo Kelapa",
			Description: "Batik Tubo Kelapa menggambarkan pohon kelapa yang menjadi simbol kehidupan masyarakat Maluku. Motif batik ini menunjukkan berbagai bagian dari pohon kelapa dan penggunaannya dalam kehidupan sehari-hari, dari makanan hingga kerajinan tangan. Filosofi di balik batik ini adalah tentang ketahanan dan keberlanjutan, mencerminkan bagaimana masyarakat Maluku memanfaatkan sumber daya alam dengan bijak. Batik ini sangat cocok untuk digunakan dalam berbagai kesempatan, memperlihatkan kecintaan akan alam dan budaya lokal.",
			City:        "Ambon",
			ProvinceID:  31, // Maluku
			IslandID:    6,  // Maluku
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Maluku/Tubo%20Kelapa.jpg?t=2024-10-14T15%3A26%3A02.794Z",
		},
		{
			Name:        "Batik Pattimura",
			Description: "Batik Pattimura menghormati pahlawan nasional, Thomas Matulessy, atau lebih dikenal sebagai Patimura, yang berjuang melawan penjajahan di Maluku. Motif batik ini melambangkan semangat perjuangan dan keberanian masyarakat Maluku dalam mempertahankan tanah air. Desainnya sering mencerminkan elemen-elemen yang berkaitan dengan sejarah dan budaya Maluku, menjadikannya sebagai simbol kebanggaan dan penghormatan. Batik ini sering digunakan dalam perayaan hari kemerdekaan dan acara-acara bersejarah.",
			City:        "Ambon",
			ProvinceID:  31, // Maluku
			IslandID:    6,  // Maluku
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Maluku/Pattimura.jpeg?t=2024-10-14T15%3A26%3A08.706Z",
		},
		{
			Name:        "Batik Tanah Seram",
			Description: "Batik Tanah Seram terinspirasi oleh keindahan alam dan kebudayaan pulau Seram. Motif batik ini mencerminkan kekayaan flora dan fauna serta adat istiadat masyarakat setempat. Dengan desain yang terinspirasi oleh tanah dan alam, batik ini menggambarkan hubungan yang erat antara masyarakat dan lingkungan sekitar. Batik Tanah Seram sering digunakan dalam acara adat dan tradisi untuk menunjukkan kecintaan terhadap budaya dan tanah air.",
			City:        "Seram",
			ProvinceID:  31, // Maluku
			IslandID:    6,  // Maluku
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Maluku/TanahSeram.jpg?t=2024-10-14T15%3A26%3A13.997Z",
		},
		{
			Name:        "Batik Pala Salawaku",
			Description: "Batik Pala Salawaku terinspirasi oleh rempah-rempah terkenal dari Maluku, khususnya pala. Motif ini melambangkan kekayaan alam Maluku yang diakui di seluruh dunia. Batik ini juga menceritakan sejarah perdagangan rempah yang menghubungkan Maluku dengan berbagai penjuru dunia. Dengan desain yang menggambarkan biji pala dan elemen alam lainnya, batik ini menjadi simbol identitas dan kekayaan budaya masyarakat Maluku.",
			City:        "Ambon",
			ProvinceID:  31, // Maluku
			IslandID:    6,  // Maluku
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Maluku/PalaSalawaku.jpeg?t=2024-10-14T15%3A26%3A21.287Z",
		},
		{
			Name:        "Batik Tifa Totobuang",
			Description: "Batik Tifa Totobuang mengacu pada alat musik tradisional yang digunakan dalam berbagai acara budaya di Maluku. Motif batik ini menggambarkan kesenangan dan semangat dalam merayakan kebersamaan dan tradisi. Tifa Totobuang melambangkan keterikatan masyarakat Maluku dengan seni dan budaya, serta menjadi simbol penyampaian pesan-pesan melalui irama dan nada. Batik ini sering digunakan dalam festival dan perayaan, menjadi penghubung antara generasi yang berbeda.",
			City:        "Ambon",
			ProvinceID:  31, // Maluku
			IslandID:    6,  // Maluku
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Maluku/Tifa%20Totobuang.jpeg?t=2024-10-14T15%3A26%3A31.408Z",
		},
		{
			Name:        "Batik Pantai Ambon",
			Description: "Batik Pantai Ambon terinspirasi oleh keindahan alam Pantai Ambon yang mempesona. Motifnya mencerminkan ombak, pasir, dan keanekaragaman hayati laut yang dapat ditemukan di sepanjang garis pantai. Batik ini melambangkan kedamaian dan kesejukan, dengan desain yang menunjukkan hubungan erat antara masyarakat Ambon dan laut. Warna-warna cerah dan motif yang dinamis membuat batik ini menjadi pilihan populer untuk dikenakan dalam acara santai maupun formal, sekaligus menggambarkan kebanggaan akan keindahan alam setempat. Batik Pantai Ambon sering digunakan dalam festival budaya dan perayaan sebagai simbol cinta terhadap lingkungan.",
			City:        "Ambon",
			ProvinceID:  31, // Maluku
			IslandID:    6,  // Maluku
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/Maluku/Pantai%20Ambon.png",
		},
		//NTT
		{
			Name:        "Batik Kuda Sepasang",
			Description: "Batik Kuda Sepasang merupakan lambang persatuan dan cinta, terinspirasi oleh kuda yang melambangkan kekuatan dan keindahan. Motifnya seringkali menggambarkan dua ekor kuda yang berhadapan, mencerminkan hubungan harmonis dalam kehidupan masyarakat Nusa Tenggara Timur. Warna-warna cerah dan kontras dalam desainnya menunjukkan keragaman budaya yang ada di daerah ini. Batik ini biasanya digunakan dalam berbagai upacara adat, perayaan, dan juga sebagai pakaian sehari-hari.",
			City:        "Kupang",
			ProvinceID:  19, // Nusa Tenggara Timur
			IslandID:    8,  // Nusa Tenggara Timur
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/NTT/KudaSepasang.jpg?t=2024-10-14T15%3A33%3A19.824Z",
		},
		{
			Name:        "Batik Liris Kupang",
			Description: "Batik Liris Kupang adalah motif batik yang menggambarkan keindahan flora dan fauna yang ada di Kupang. Desainnya menonjolkan garis-garis halus yang membentuk pola liris, memberikan kesan elegan dan anggun. Batik ini mencerminkan kekayaan alam Nusa Tenggara Timur, serta melambangkan harapan dan semangat baru bagi masyarakat setempat. Warna yang digunakan bervariasi dari nuansa pastel hingga warna-warna cerah yang mencerminkan kehangatan daerah ini.",
			City:        "Kupang",
			ProvinceID:  19, // Nusa Tenggara Timur
			IslandID:    8,  // Nusa Tenggara Timur
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/NTT/LirisKupang.png?t=2024-10-14T15%3A33%3A28.692Z",
		},
		{
			Name:        "Batik Pucuk Mekar",
			Description: "Batik Pucuk Mekar menggambarkan motif daun dan bunga yang sedang mekar, simbol dari kehidupan dan pertumbuhan. Motif ini sering digunakan dalam perayaan dan acara penting sebagai lambang harapan dan keberkahan. Dengan perpaduan warna yang segar, batik ini mencerminkan keindahan alam Nusa Tenggara Timur. Desain yang ceria dan penuh warna ini menjadikan batik ini sebagai pilihan yang populer di kalangan masyarakat.",
			City:        "Kupang",
			ProvinceID:  19, // Nusa Tenggara Timur
			IslandID:    8,  // Nusa Tenggara Timur
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/NTT/PucukMekar.png?t=2024-10-14T15%3A33%3A44.443Z",
		},
		{
			Name:        "Batik Teguh Bersatu",
			Description: "Batik Teguh Bersatu adalah lambang dari kekuatan dan persatuan masyarakat Nusa Tenggara Timur. Motif batik ini menunjukkan pola-pola geometris yang saling terhubung, menggambarkan hubungan erat antaranggota masyarakat. Dengan warna yang kuat dan berani, batik ini merefleksikan semangat perjuangan dan kerjasama yang ada di antara mereka. Batik ini sering dipakai dalam acara-acara resmi dan pertemuan adat.",
			City:        "Kupang",
			ProvinceID:  19, // Nusa Tenggara Timur
			IslandID:    8,  // Nusa Tenggara Timur
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/NTT/TeguhBersatu.png?t=2024-10-14T15%3A33%3A51.924Z",
		},
		{
			Name:        "Batik Kuda Kupang",
			Description: "Batik Kuda Kupang terinspirasi oleh kuda-kuda yang menjadi bagian penting dari budaya lokal. Motif ini menggambarkan kuda dengan desain yang artistik dan penuh warna, melambangkan kekuatan, kebanggaan, dan kebudayaan masyarakat Nusa Tenggara Timur. Batik ini digunakan dalam berbagai acara, termasuk perayaan dan upacara adat, serta menjadi simbol identitas masyarakat Kupang.",
			City:        "Kupang",
			ProvinceID:  19, // Nusa Tenggara Timur
			IslandID:    8,  // Nusa Tenggara Timur
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/NTT/Kuda%20Kupang.jpg?t=2024-10-14T15%3A33%3A57.641Z",
		},
		{
			Name:        "Batik Kupang Rukun",
			Description: "Batik Kupang Rukun adalah simbol persatuan dan harmoni dalam masyarakat. Motifnya sering menggambarkan elemen-elemen yang saling melengkapi, mencerminkan semangat kebersamaan dan gotong royong. Dengan penggunaan warna-warna yang harmonis, batik ini menjadi pilihan yang tepat untuk digunakan dalam berbagai acara, baik formal maupun non-formal, sebagai representasi cinta dan kesatuan.",
			City:        "Kupang",
			ProvinceID:  19, // Nusa Tenggara Timur
			IslandID:    8,  // Nusa Tenggara Timur
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/NTT/KupangRukun.jpg?t=2024-10-14T15%3A34%3A03.744Z",
		},
		//NTB
		{
			Name:        "Batik Daun Sirih",
			Description: "Batik Daun Sirih merupakan salah satu motif batik yang terinspirasi dari bentuk daun sirih. Motif ini mencerminkan nilai-nilai kehidupan masyarakat yang erat dengan tradisi dan budaya. Dengan penggunaan warna yang segar dan alami, batik ini melambangkan kesegaran dan harapan. Setiap detail pada batik ini menceritakan kisah tentang ketahanan dan kebersamaan masyarakat NTB dalam menjaga budaya dan tradisi mereka. Motif ini sering digunakan dalam berbagai acara resmi dan perayaan.",
			City:        "Mataram",
			ProvinceID:  18, // Nusa Tenggara Barat
			IslandID:    7,  // Nusa Tenggara
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/NTB/DaunSirih.jpeg?t=2024-10-14T15%3A44%3A15.426Z",
		},
		{
			Name:        "Batik Bale Lumbu",
			Description: "Batik Bale Lumbu adalah motif batik yang terinspirasi dari bentuk dan desain rumah tradisional suku Sasak. Motif ini menggambarkan arsitektur rumah yang khas dan mencerminkan keindahan budaya NTB. Dengan pola yang elegan dan warna yang hangat, batik ini membawa nuansa kearifan lokal. Setiap goresan dan bentuk dalam batik ini adalah simbol dari kehidupan sehari-hari masyarakat yang dekat dengan alam dan tradisi. Batik Bale Lumbu sering dikenakan pada acara adat dan perayaan penting.",
			City:        "Bima",
			ProvinceID:  18, // Nusa Tenggara Barat
			IslandID:    7,  // Nusa Tenggara
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/NTB/BaleLumbu.png",
		},
		{
			Name:        "Batik Tenun Bima",
			Description: "Batik Tenun Bima menggabungkan teknik batik dengan seni tenun yang telah ada sejak lama di Bima. Motif batik ini kaya akan warna dan pola, mencerminkan keindahan alam dan budaya masyarakat Bima. Batik ini melambangkan semangat dan keteguhan masyarakat dalam mempertahankan warisan budaya. Dalam setiap seratnya, terdapat cerita yang menghubungkan generasi, menjadikan batik ini tidak hanya sekedar pakaian, tetapi juga sebuah karya seni yang hidup. Batik Tenun Bima sering dikenakan dalam acara-acara formal dan budaya.",
			City:        "Bima",
			ProvinceID:  18, // Nusa Tenggara Barat
			IslandID:    7,  // Nusa Tenggara
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/NTB/TenunanBima.png",
		},
		{
			Name:        "Batik Tari Persean",
			Description: "Batik Tari Persean terinspirasi dari tarian tradisional yang terkenal di NTB, yaitu Tari Persean. Motif ini mencerminkan gerakan, semangat, dan keindahan dari tarian tersebut. Dengan pola yang dinamis dan warna yang cerah, batik ini menggambarkan keceriaan dan kekuatan budaya NTB. Setiap detail pada batik ini menggambarkan cerita dan makna dari tari tersebut, menjadikannya unik dan berharga. Batik Tari Persean sering digunakan dalam acara seni dan budaya, serta sebagai simbol kebanggaan masyarakat NTB.",
			City:        "Mataram",
			ProvinceID:  18, // Nusa Tenggara Barat
			IslandID:    7,  // Nusa Tenggara
			Link_Image:  "https://jjittozptnoaoseyrvcz.supabase.co/storage/v1/object/public/hology/Batik/NTB/TariPersean.png?t=2024-10-14T15%3A44%3A43.563Z",
		},
	}

	if err := db.CreateInBatches(batikSeeder, len(*batikSeeder)).Error; err != nil {
		panic(err)
	}
}

func generateQuiz(db *gorm.DB) {
	var quizzes = []domain.Quiz{
		// Easy questions
		{Question: "Dari mana asal motif Batik Parang?", Answer: "A", OptionA: "Jawa Tengah", OptionB: "Bali", OptionC: "Sumatra", OptionD: "Kalimantan", Difficulty: "easy", Image_Link: ""},
		{Question: "Warna apa yang paling umum ditemukan dalam Batik tradisional?", Answer: "B", OptionA: "Hijau", OptionB: "Cokelat", OptionC: "Ungu", OptionD: "Merah Muda", Difficulty: "easy", Image_Link: ""},
		{Question: "Hewan apa yang sering digambarkan dalam Batik Megamendung?", Answer: "D", OptionA: "Singa", OptionB: "Harimau", OptionC: "Rusa", OptionD: "Awan", Difficulty: "easy", Image_Link: ""},
		{Question: "Alat apa yang digunakan untuk menerapkan lilin dalam proses Batik?", Answer: "A", OptionA: "Canting", OptionB: "Alat Tenun", OptionC: "Kuasa", OptionD: "Gunting", Difficulty: "easy", Image_Link: ""},
		{Question: "Batik Tulis adalah jenis Batik yang dibuat dengan?", Answer: "B", OptionA: "Mencap", OptionB: "Menggambar dengan tangan", OptionC: "Mesin", OptionD: "Menenun", Difficulty: "easy", Image_Link: ""},
		{Question: "Motif Batik apa yang biasa dikenakan oleh keluarga kerajaan Indonesia?", Answer: "C", OptionA: "Kawung", OptionB: "Lasem", OptionC: "Parang", OptionD: "Sekar Jagad", Difficulty: "easy", Image_Link: ""},
		{Question: "Motif Batik dengan bentuk geometris sering disebut?", Answer: "D", OptionA: "Floral", OptionB: "Burung", OptionC: "Awan", OptionD: "Kawung", Difficulty: "easy", Image_Link: ""},
		{Question: "Pulau mana yang menjadi asal Batik Lasem?", Answer: "A", OptionA: "Jawa", OptionB: "Sumatra", OptionC: "Kalimantan", OptionD: "Sulawesi", Difficulty: "easy", Image_Link: ""},
		{Question: "Warna berikut yang BUKAN merupakan warna tradisional Batik adalah?", Answer: "C", OptionA: "Indigo", OptionB: "Sogan", OptionC: "Hijau Neon", OptionD: "Merah", Difficulty: "easy", Image_Link: ""},
		{Question: "Batik diakui oleh UNESCO sebagai warisan?", Answer: "B", OptionA: "Lukisan Tradisional", OptionB: "Warisan Budaya Takbenda", OptionC: "Sastra", OptionD: "Arsitektur", Difficulty: "easy", Image_Link: ""},

		// Medium questions
		{Question: "Motif Batik apa yang melambangkan kesuburan dan pertumbuhan, sering menggambarkan bunga dan tanaman?", Answer: "A", OptionA: "Sekar Jagad", OptionB: "Parang", OptionC: "Kawung", OptionD: "Lasem", Difficulty: "medium", Image_Link: ""},
		{Question: "Apa perbedaan utama antara Batik Tulis dan Batik Cap?", Answer: "B", OptionA: "Batik Tulis menggunakan lilin, sedangkan Batik Cap menggunakan pewarna", OptionB: "Batik Tulis digambar dengan tangan, sedangkan Batik Cap dicap", OptionC: "Batik Tulis menggunakan kuas, sedangkan Batik Cap menggunakan alat canting", OptionD: "Tidak ada perbedaan", Difficulty: "medium", Image_Link: ""},
		{Question: "Motif Batik 'Garuda' terinspirasi oleh tokoh mitologi apa?", Answer: "C", OptionA: "Harimau", OptionB: "Elang", OptionC: "Garuda", OptionD: "Naga", Difficulty: "medium", Image_Link: ""},
		{Question: "Daerah mana yang terkenal dengan Batik yang menggunakan warna-warna cerah dan non-tradisional?", Answer: "D", OptionA: "Jawa Tengah", OptionB: "Bali", OptionC: "Sumatra Barat", OptionD: "Pekalongan", Difficulty: "medium", Image_Link: ""},
		{Question: "Teknik Batik apa yang melibatkan pewarnaan kain beberapa kali untuk mendapatkan warna berbeda?", Answer: "A", OptionA: "Pewarnaan Resist", OptionB: "Mencap", OptionC: "Batik Tenun", OptionD: "Sablon", Difficulty: "medium", Image_Link: ""},
		{Question: "Motif Batik apa yang dilarang dikenakan oleh rakyat jelata pada zaman dahulu?", Answer: "B", OptionA: "Batik Pekalongan", OptionB: "Batik Larangan", OptionC: "Batik Pesisir", OptionD: "Batik Palembang", Difficulty: "medium", Image_Link: ""},
		{Question: "Simbolisme apa yang ada di balik motif Batik Ceplok?", Answer: "C", OptionA: "Kekuasaan dan Keagungan", OptionB: "Harmoni dan Keseimbangan", OptionC: "Keteraturan dan Struktur", OptionD: "Kekayaan dan Kemakmuran", Difficulty: "medium", Image_Link: ""},
		{Question: "Di daerah mana warna cokelat 'Sogan' sering digunakan dalam Batik?", Answer: "D", OptionA: "Bali", OptionB: "Sumatra", OptionC: "Pekalongan", OptionD: "Yogyakarta", Difficulty: "medium", Image_Link: ""},
		{Question: "Motif Batik yang menggambarkan lautan dan ombak dikenal sebagai?", Answer: "A", OptionA: "Megamendung", OptionB: "Parang", OptionC: "Garuda", OptionD: "Cendrawasih", Difficulty: "medium", Image_Link: ""},
		{Question: "Motif Batik 'Kawung' diyakini melambangkan apa?", Answer: "B", OptionA: "Perlindungan", OptionB: "Kemurnian dan Keadilan", OptionC: "Kekuatan", OptionD: "Kekayaan", Difficulty: "medium", Image_Link: ""},

		// Hard questions
		{Question: "Apa yang dilambangkan motif Batik 'Sawunggaling' yang menggambarkan ayam jantan?", Answer: "D", OptionA: "Kemenangan", OptionB: "Keagungan", OptionC: "Cinta", OptionD: "Kekuatan dan Keberanian", Difficulty: "hard", Image_Link: ""},
		{Question: "Teknik Batik apa yang menggunakan cap tembaga yang dikombinasikan dengan elemen gambar tangan?", Answer: "A", OptionA: "Batik Kombinasi", OptionB: "Batik Tulis", OptionC: "Batik Cap", OptionD: "Batik Printing", Difficulty: "hard", Image_Link: ""},
		{Question: "Motif Batik 'Ceplok' sering dikaitkan dengan konsep kosmologi apa dalam budaya Jawa?", Answer: "C", OptionA: "Keseimbangan hidup", OptionB: "Siklus reinkarnasi", OptionC: "Alam semesta dan keabadian", OptionD: "Perjalanan jiwa", Difficulty: "hard", Image_Link: ""},
		{Question: "Batik Lasem unik karena menggunakan warna apa yang jarang ditemukan dalam tradisi Batik lainnya?", Answer: "B", OptionA: "Cokelat Sogan", OptionB: "Merah Tua", OptionC: "Biru Indigo", OptionD: "Kuning Pucat", Difficulty: "hard", Image_Link: ""},
		{Question: "Motif 'Parang Rusak' dulu dilarang dikenakan oleh siapa?", Answer: "C", OptionA: "Kaum Bangsawan", OptionB: "Para Pendeta", OptionC: "Rakyat Jelata", OptionD: "Para Pedagang", Difficulty: "hard", Image_Link: ""},
		{Question: "Peristiwa sejarah apa yang menyebabkan perkembangan besar dalam pola Batik, terutama di daerah pesisir seperti Cirebon dan Lasem?", Answer: "A", OptionA: "Penjajahan Belanda", OptionB: "Munculnya Kerajaan Mataram", OptionC: "Migrasi Tiongkok", OptionD: "Masuknya Islam", Difficulty: "hard", Image_Link: ""},
		{Question: "Batik Madura dikenal dengan pola dan warna-warna berani. Faktor lingkungan apa yang mempengaruhi ciri khas ini?", Answer: "D", OptionA: "Komposisi tanah", OptionB: "Dekat dengan laut", OptionC: "Daerah pegunungan", OptionD: "Iklim kering dan gersang", Difficulty: "hard", Image_Link: ""},
		{Question: "Motif Batik mana yang diyakini memiliki kekuatan magis untuk melindungi pemakainya?", Answer: "C", OptionA: "Batik Lasem", OptionB: "Batik Garuda", OptionC: "Batik Parang", OptionD: "Batik Kawung", Difficulty: "hard", Image_Link: ""},
		{Question: "Motif Batik mana yang menjadi simbol perlawanan melawan penjajahan kolonial di Indonesia?", Answer: "B", OptionA: "Batik Megamendung", OptionB: "Batik Parang", OptionC: "Batik Kawung", OptionD: "Batik Lasem", Difficulty: "hard", Image_Link: ""},
		{Question: "Motif apa yang banyak digunakan dalam Batik Keraton dan sering mengandung makna filosofis yang dalam?", Answer: "A", OptionA: "Batik Parang", OptionB: "Batik Sekar Jagad", OptionC: "Batik Megamendung", OptionD: "Batik Garuda", Difficulty: "hard", Image_Link: ""},
	}

	if err := db.CreateInBatches(quizzes, 30).Error; err != nil {
		panic(err)
	}
}

func generateProvince(db *gorm.DB) {
	var provinces = []domain.Province{
		{Name: "Aceh"},
		{Name: "Sumatera Utara"},
		{Name: "Sumatera Barat"},
		{Name: "Riau"},
		{Name: "Kepulauan Riau"},
		{Name: "Jambi"},
		{Name: "Sumatera Selatan"},
		{Name: "Bangka Belitung"},
		{Name: "Bengkulu"},
		{Name: "Lampung"},
		{Name: "DKI Jakarta"},
		{Name: "Jawa Barat"},
		{Name: "Banten"},
		{Name: "Jawa Tengah"},
		{Name: "DI Yogyakarta"},
		{Name: "Jawa Timur"},
		{Name: "Bali"},
		{Name: "Nusa Tenggara Barat"},
		{Name: "Nusa Tenggara Timur"},
		{Name: "Kalimantan Barat"},
		{Name: "Kalimantan Tengah"},
		{Name: "Kalimantan Selatan"},
		{Name: "Kalimantan Timur"},
		{Name: "Kalimantan Utara"},
		{Name: "Sulawesi Utara"},
		{Name: "Gorontalo"},
		{Name: "Sulawesi Tengah"},
		{Name: "Sulawesi Barat"},
		{Name: "Sulawesi Selatan"},
		{Name: "Sulawesi Tenggara"},
		{Name: "Maluku"},
		{Name: "Maluku Utara"},
		{Name: "Papua"},
		{Name: "Papua Barat"},
		{Name: "Papua Tengah"},
		{Name: "Papua Pegunungan"},
		{Name: "Papua Selatan"},
		{Name: "Papua Barat Daya"},
	}

	if err := db.CreateInBatches(provinces, len(provinces)).Error; err != nil {
		panic(err)
	}
}

func generateIsland(db *gorm.DB) {
	island := []domain.Island{
		{
			Name: "Jawa",
		},
		{
			Name: "Sumatra",
		},
		{
			Name: "Kalimantan",
		},
		{
			Name: "Sulawesi",
		},
		{
			Name: "Papua",
		},
		{
			Name: "Maluku",
		},
		{
			Name: "Nusa Tenggara Barat",
		},
		{
			Name: "Nusa Tenggara Timur",
		},
	}

	if err := db.CreateInBatches(island, len(island)).Error; err != nil {
		panic(err)
	}
}
