-- MySQL dump 10.14  Distrib 5.5.65-MariaDB, for Linux (x86_64)
--
-- Host: post-db.cdvmgnpfzvjx.ap-northeast-1.rds.amazonaws.com    Database: ramen_app
-- ------------------------------------------------------
-- Server version	8.0.17

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `likes`
--

DROP TABLE IF EXISTS `likes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `likes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `likes`
--

LOCK TABLES `likes` WRITE;
/*!40000 ALTER TABLE `likes` DISABLE KEYS */;
INSERT INTO `likes` VALUES (1,2),(2,2),(3,2),(4,3),(5,3),(6,3),(8,1),(11,3),(12,2),(13,2),(14,3);
/*!40000 ALTER TABLE `likes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `post_likes`
--

DROP TABLE IF EXISTS `post_likes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `post_likes` (
  `post_id` int(11) NOT NULL,
  `like_id` int(11) NOT NULL,
  PRIMARY KEY (`post_id`,`like_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `post_likes`
--

LOCK TABLES `post_likes` WRITE;
/*!40000 ALTER TABLE `post_likes` DISABLE KEYS */;
INSERT INTO `post_likes` VALUES (1,2),(4,1),(4,6),(6,3),(6,5),(7,14),(9,13),(10,4),(10,8),(12,11),(12,12);
/*!40000 ALTER TABLE `post_likes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `posts`
--

DROP TABLE IF EXISTS `posts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `posts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `content` varchar(400) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `photo_url` varchar(400) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `like_count` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `store_info` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `posts`
--

LOCK TABLES `posts` WRITE;
/*!40000 ALTER TABLE `posts` DISABLE KEYS */;
INSERT INTO `posts` VALUES (1,'二郎上野毛店に行ってきました！','金曜日、開店10分前に到着すると4人待ち。\n駅からめちゃくちゃ近いな〜と思ってると11時より3〜4分早くシャッターが開いた。\n噂のデロ麺が食べたいので麺の硬さはデフォで。\nヤサイ・ニンニク・アブラをコール。\n液体油が熱々をさらに熱々にしている。\nやや細めの麺を持ち上げ、これがデロ麺なのか、と感慨深くすする。\n豚はみっちりしていて歯応えのあるもの。しょっぱいのでヤサイと相性バッチリだった。\n固形のアブラがとんでもない美味しさなので次があればアブラマシマシにしたい。\n麺の量などもいろいろちょうど良かった。\n駅から近すぎて、もう少し歩きたいなと二子玉川の蔦屋家電まで歩いたが、オシャレとニンニクは真逆のものだな…と長居せずそそくさ帰った。','https://res.cloudinary.com/dxo10noyu/raw/upload/v1592395061/o6rjnmz7wavuly5fd9m4.jpg',1,0,'2020-06-17 11:57:42','2020-06-17 11:57:42',NULL,'https://tabelog.com/tokyo/A1317/A131715/13010417/'),(2,'無邪気に無邪気 自由が丘南口店 ','最早老舗といってもよいこちらのお店のラーメンは完成度が高い無邪気 。\nあーこのとんこつベースのラーメンの味わいは日本でこそ味わえるもの。\n身体に染み渡るのがわかる。あらためてありがたい。ごちそう様でした。','https://res.cloudinary.com/dxo10noyu/raw/upload/v1592395220/dtpgtjk1xusm8dttpcvf.jpg',1,0,'2020-06-17 12:00:21','2020-06-17 12:00:21',NULL,NULL),(3,'大井町であっさり塩ラーメンの焔 ','スープは相変わらず優しい味です。塩味はそれ程強くなく、鶏ベースだけど、主張しない程度に魚介も感じられて、良い味しています。\nチャーシューは、流行りの低温調理ではなく、しっかり火を通した上に、表面を炙っていて、肉の味がしっかり感じられました。\n味玉も良い感じの半熟だし、相変わらずレベルの高いラーメンだと思います。','https://res.cloudinary.com/dxo10noyu/raw/upload/v1592395347/cfdazadfqmo10gf6ww87.jpg',1,0,'2020-06-17 12:02:28','2020-06-17 12:02:28',NULL,NULL),(4,'麺匠 濱星＠溝の口','券売機制で券を先に買って入店。濃厚煮干しそば(750円)のチケットを購入して入店。\n「肉にこだわった」と店が豪語するように柔らかめのチャーシューは食感も良い。2枚も入っておりお得感がある。\n飲んだ締めにも良さそうだが、「濃厚」は言葉通り、とんこつラーメンのようなクリーミーな食感で少し重さを感じた。','https://res.cloudinary.com/dxo10noyu/raw/upload/v1592395645/clfwktcsx87vx82dkkr5.jpg',1,0,'2020-06-17 12:07:26','2020-06-17 12:07:26',NULL,'https://tabelog.com/kanagawa/A1405/A140505/14064603/'),(5,'マッチでぇ～す！！','溝の口駅の近くのまっち棒 溝の口店 \nどんぶりのヘリに沢山の脂が・・・スープの色合いがドスいというか、見るからにドロっとしてそうな感じ。屈強な和歌山ラーメンという印象。\nほっほ～、やっぱり口の中に豚骨の強い香りが広がる。というか、食べる→スープが鼻に近付く→嗅覚が豚をキャッチ。そんな感じ。けどね、そんな香りの割にそこまで口の中には残らないというか。味も割と主張が強いのに、喉を通るとしつこい感じはない。香りこそあれですが、豚骨の旨みをしっかり生かしているなって。\n二郎系とはまた違うけど、こんなスープを使ってて、ただ麺だけを味わうのは違う、がっつり行きたいなって思わせるラーメン。','https://res.cloudinary.com/dxo10noyu/raw/upload/v1592395818/mhf9bgi9wvyrixysdqpj.jpg',1,0,'2020-06-17 12:10:19','2020-06-17 12:10:19',NULL,'https://tabelog.com/kanagawa/A1405/A140505/14006090/'),(6,'やっじゃがな','しこたま飲んだ後の定番のやっじゃがな。\nお腹いっぱいのはずが替え玉まで。\n二玉目は辛味をいれて。。。\nあー今日も飲み過ぎた！\n飲んだ後にはこのシンプルなラーメンが最高！\nご馳走さまでした！','https://res.cloudinary.com/dxo10noyu/raw/upload/v1592396618/hxsiy7rhbjo1wwhh7bq5.jpg',2,0,'2020-06-17 12:23:39','2020-06-17 12:23:39',NULL,NULL),(7,'田中は王道の家系ラーメン','いや〜久しぶりの家系ラーメン！ですが、前夜の仕事飲みのダメージが大きく苦渋のどノーマル…\n嫌な臭みのないクリーミーなスープは二日酔いの胃袋に膜を張って守ってくれる。\nほうれん草に含まれる豊富なビタミンが身体の疲れを癒してくれる。\n薄切りのチャーシューはさっぱり系で二日酔いの身体に優しく寄り添いチカラを与えてくれる。\n家系ラーメン特有の海苔たちはダメダメな心をシャキッと！させてくれるほどパリッとした食感で磯を思い出す。','https://res.cloudinary.com/dxo10noyu/raw/upload/v1592396875/lzf1juyibaqhx3dunbnw.jpg',2,0,'2020-06-17 12:27:56','2020-06-17 12:27:56',NULL,NULL),(8,'五目蒙古タンメン旨し','たっぷり野菜の“味噌タンメン”に当店秘伝の辛し麻婆をトッピングした“蒙古タンメン”、\nそれに辛し肉とゆで卵をトッピングした“五目蒙古タンメン”～\nゆで卵は超固茹で、スライスしてありますが１個分。\n私は断然半熟派ですが、このスープには固茹でが合うかも！！！\n秘伝の味噌ダレでじっくりと煮込まれた、ごそっと肉、野菜がたまらんね～\nヘルシー感満載。\nもっちり太麺をずずっとうまい！！！','https://res.cloudinary.com/dxo10noyu/raw/upload/v1592397174/jzjhg2qpvo8v0tj1rj5l.jpg',2,0,'2020-06-17 12:32:55','2020-06-17 12:32:55',NULL,NULL),(9,'煮干し中毒に注意！！','JR京浜東北線東神奈川駅より徒歩5分。\n改札を出て西口の陸橋で国道1号線を越えて階段を降りる。そのまま国道1号線を横浜方面へ300ｍ程進んだところに店がある「らぁめん夢」\nラーメンに注射器に入った煮干油が付いてきて、それを途中で注入して味変させるという斬新なアイデアで評判らしい。\n\n具は特製だけに具だくさん。始めからスープには味玉とワンタンと海苔がセットされてますが別皿に厚めのチャーシュー・メンマ・青菜・薬味のネギとかなりたっぷり。麺大盛なので1,000円を超えちゃいましたがこの内容ならお得かも。\nという感じで相変わらず美味いのにお得感ありありな一杯で満足できました。美味かったです。比較的近いのでまた伺います。ごちそうさまでした。','https://res.cloudinary.com/dxo10noyu/raw/upload/v1592397949/z7cybhyuc4a6ejghh6eg.jpg',3,0,'2020-06-17 12:45:50','2020-06-17 12:45:50',NULL,'https://tabelog.com/kanagawa/A1401/A140212/14053385/'),(10,'気持ちもお腹もほっこりこ(*´ー｀*)','八王子の雄と名高い『ほっこり中華そば もつけ』さんを訪問です。\n煮干し、節、鶏などでひいたスープに醤油のカエシを合わせたスープは、いかにも正統派中華そばというビジュアル♪\nレンゲで掬って飲むと、煮干しと魚介系の節が前面に出たしっかりと力強いスープです( ´艸｀)\n自家製という麺はのど越しが良くシルクのようでありながらも、噛み締めるとぷりっとした弾力があって、こいつはめっちゃ美味い＼(^o^)／\nこの日は店主とお母さんらしき方の2名体制でしたが、お二人ともホスピタリティ溢れる優しい接客で素晴らしかった！\nお腹ばかりか気持ちまでほっこりさせられ、大満足でお店を後にしました。\n美味しかった～♪','https://res.cloudinary.com/dxo10noyu/raw/upload/v1592398359/jndytqh8lkvkace3wfcr.jpg',3,0,'2020-06-17 12:52:40','2020-06-17 12:52:40',NULL,'https://retty.me/area/PRE13/ARE3/SUB302/100001266379/?utm_y_pc_res_title'),(11,'投稿テスト','大丈夫かな\nどうかな？','https://res.cloudinary.com/dxo10noyu/raw/upload/v1592930088/fqmlouh36dcdvbnrzfbb.jpg',1,0,'2020-06-23 16:34:50','2020-06-23 16:34:50','2020-06-23 16:35:15',NULL),(12,'大岡山の店','食感は二郎というよりも、やはり大勝軒寄り。\nつけの方が断然美味しいかもw\n\nスープはやたらと甘い。ぶたを観て予想できたけど\n予想通りの甘さかな。かえしの作りが今一歩。\n\nベースのスープは悪くない。むしろ美味しいと思うけど。\nかえしが良く無いと仕上げも変わってしまう。\n少し勿体ないかな。\n\nぶたは薄切りの豚肉を網で焼いてある。焼肉風でご飯に\n良く合うと思う。これはこれでなかなか美味しい。\n全体的には悪く無いけど、甘いのが残念。\n甘いのが食べたくなった時、また行くかも。\n\nごちそうさまでした。','https://res.cloudinary.com/dxo10noyu/raw/upload/v1594559886/qspsgckznndpwdcrpwqb.jpg',1,NULL,'2020-07-12 12:14:23','2020-07-12 12:14:23',NULL,'https://tabelog.com/tokyo/A1317/A131711/13048427/');
/*!40000 ALTER TABLE `posts` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-08-02 19:28:47
