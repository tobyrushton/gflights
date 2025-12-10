// Package iata contains IATA airport codes, which are supported by the Google Flights API, along with time zones.
// This package was generated using an airport list (which can be found at this address: [airports.json])
// and the Google Flights API.
//
// Command: go run ./iata/cmd/gen.go
//
// Generation date: 2025-12-10
//
// [airports.json]: https://raw.githubusercontent.com/mwgg/Airports/refs/heads/master/airports.json
package iata

type Location struct{ City, Tz string }

// IATATimeZone turns IATA airport codes into the time zone where the airport is located.
// If IATATimeZone can't find an IATA airport code, then it returns "Not supported IATA Code".
func IATATimeZone(iata string) Location {
	switch iata {
	case "AAA":
		return Location{"", "Pacific/Tahiti"}
	case "AAB":
		return Location{"", "Australia/Brisbane"}
	case "AAC":
		return Location{"El Arish", "Africa/Cairo"}
	case "AAD":
		return Location{"Adado", "Africa/Mogadishu"}
	case "AAE":
		return Location{"Annabah", "Africa/Algiers"}
	case "AAF":
		return Location{"Apalachicola", "America/New_York"}
	case "AAG":
		return Location{"Arapoti", "America/Sao_Paulo"}
	case "AAH":
		return Location{"Aachen", "Europe/Berlin"}
	case "AAI":
		return Location{"Arraias", "America/Araguaina"}
	case "AAJ":
		return Location{"Awaradam", "America/Paramaribo"}
	case "AAK":
		return Location{"Buariki", "Pacific/Tarawa"}
	case "AAL":
		return Location{"Aalborg", "Europe/Copenhagen"}
	case "AAM":
		return Location{"Malamala", "Africa/Johannesburg"}
	case "AAN":
		return Location{"Al Ain", "Asia/Dubai"}
	case "AAO":
		return Location{"Anaco", "America/Caracas"}
	case "AAP":
		return Location{"Samarinda", "Asia/Makassar"}
	case "AAQ":
		return Location{"Anapa", "Europe/Moscow"}
	case "AAR":
		return Location{"Aarhus", "Europe/Copenhagen"}
	case "AAT":
		return Location{"Altay", "Asia/Shanghai"}
	case "AAU":
		return Location{"Asau", "Pacific/Apia"}
	case "AAV":
		return Location{"Surallah", "Asia/Manila"}
	case "AAW":
		return Location{"Abbottabad", "Asia/Karachi"}
	case "AAX":
		return Location{"Araxa", "America/Sao_Paulo"}
	case "AAY":
		return Location{"", "Asia/Aden"}
	case "AAZ":
		return Location{"Quezaltenango", "America/Guatemala"}
	case "ABA":
		return Location{"Abakan", "Asia/Krasnoyarsk"}
	case "ABB":
		return Location{"Asaba", "Africa/Lagos"}
	case "ABC":
		return Location{"Albacete", "Europe/Madrid"}
	case "ABD":
		return Location{"Abadan", "Asia/Tehran"}
	case "ABE":
		return Location{"Allentown", "America/New_York"}
	case "ABF":
		return Location{"Abaiang", "Pacific/Tarawa"}
	case "ABG":
		return Location{"", "Australia/Brisbane"}
	case "ABH":
		return Location{"", "Australia/Brisbane"}
	case "ABI":
		return Location{"Abilene", "America/Chicago"}
	case "ABJ":
		return Location{"Abidjan", "Africa/Abidjan"}
	case "ABK":
		return Location{"Kabri Dehar", "Africa/Addis_Ababa"}
	case "ABL":
		return Location{"Ambler", "America/Anchorage"}
	case "ABM":
		return Location{"", "Australia/Brisbane"}
	case "ABN":
		return Location{"Albina", "America/Cayenne"}
	case "ABO":
		return Location{"Aboisso", "Africa/Abidjan"}
	case "ABQ":
		return Location{"Albuquerque", "America/Denver"}
	case "ABR":
		return Location{"Aberdeen", "America/Chicago"}
	case "ABS":
		return Location{"Abu Simbel", "Africa/Cairo"}
	case "ABT":
		return Location{"", "Asia/Riyadh"}
	case "ABU":
		return Location{"Atambua-Timor Island", "Asia/Makassar"}
	case "ABV":
		return Location{"Abuja", "Africa/Lagos"}
	case "ABX":
		return Location{"Albury", "Australia/Melbourne"}
	case "ABY":
		return Location{"Albany", "America/New_York"}
	case "ABZ":
		return Location{"Aberdeen", "Europe/London"}
	case "ACA":
		return Location{"Acapulco", "America/Mexico_City"}
	case "ACB":
		return Location{"Bellaire", "America/Detroit"}
	case "ACC":
		return Location{"Accra", "Africa/Accra"}
	case "ACD":
		return Location{"Acandi", "America/Bogota"}
	case "ACE":
		return Location{"Lanzarote Island", "Atlantic/Canary"}
	case "ACF":
		return Location{"Aral", "Asia/Shanghai"}
	case "ACH":
		return Location{"Altenrhein", "Europe/Vienna"}
	case "ACI":
		return Location{"Saint Anne", "Europe/Guernsey"}
	case "ACJ":
		return Location{"Anuradhapura", "Asia/Colombo"}
	case "ACK":
		return Location{"Nantucket", "America/New_York"}
	case "ACN":
		return Location{"Ciudad Acuna", "America/Matamoros"}
	case "ACP":
		return Location{"Maragheh", "Asia/Tehran"}
	case "ACR":
		return Location{"Araracuara", "America/Bogota"}
	case "ACS":
		return Location{"Achinsk", "Asia/Krasnoyarsk"}
	case "ACT":
		return Location{"Waco", "America/Chicago"}
	case "ACV":
		return Location{"Arcata/Eureka", "America/Los_Angeles"}
	case "ACX":
		return Location{"Xingyi", "Asia/Shanghai"}
	case "ACY":
		return Location{"Atlantic City", "America/New_York"}
	case "ACZ":
		return Location{"", "Asia/Tehran"}
	case "ADA":
		return Location{"Adana", "Europe/Istanbul"}
	case "ADB":
		return Location{"Izmir", "Europe/Istanbul"}
	case "ADC":
		return Location{"Andekombe", "Pacific/Port_Moresby"}
	case "ADD":
		return Location{"Addis Ababa", "Africa/Addis_Ababa"}
	case "ADE":
		return Location{"Aden", "Asia/Aden"}
	case "ADF":
		return Location{"Adiyaman", "Europe/Istanbul"}
	case "ADG":
		return Location{"Adrian", "America/Detroit"}
	case "ADH":
		return Location{"Aldan", "Asia/Yakutsk"}
	case "ADI":
		return Location{"Arandis", "Africa/Windhoek"}
	case "ADJ":
		return Location{"Amman", "Asia/Amman"}
	case "ADK":
		return Location{"Adak Island", "America/Adak"}
	case "ADL":
		return Location{"Adelaide", "Australia/Adelaide"}
	case "ADM":
		return Location{"Ardmore", "America/Chicago"}
	case "ADN":
		return Location{"Andes", "America/Bogota"}
	case "ADO":
		return Location{"", "Australia/Adelaide"}
	case "ADP":
		return Location{"Ampara", "Asia/Colombo"}
	case "ADQ":
		return Location{"Kodiak", "America/Anchorage"}
	case "ADR":
		return Location{"Andrews", "America/New_York"}
	case "ADS":
		return Location{"Dallas", "America/Chicago"}
	case "ADT":
		return Location{"Ada", "America/Chicago"}
	case "ADU":
		return Location{"Ardabil", "Asia/Tehran"}
	case "ADW":
		return Location{"Camp Springs", "America/New_York"}
	case "ADX":
		return Location{"St. Andrews", "Europe/London"}
	case "ADY":
		return Location{"Alldays", "Africa/Johannesburg"}
	case "ADZ":
		return Location{"San Andres", "America/Bogota"}
	case "AEA":
		return Location{"Abemama Atoll", "Pacific/Tarawa"}
	case "AEB":
		return Location{"Baise", "Asia/Shanghai"}
	case "AEG":
		return Location{"Padang Sidempuan", "Asia/Jakarta"}
	case "AEH":
		return Location{"", "Africa/Ndjamena"}
	case "AEL":
		return Location{"Albert Lea", "America/Chicago"}
	case "AEM":
		return Location{"Amgu", "Asia/Vladivostok"}
	case "AEO":
		return Location{"Aioun El Atrouss", "Africa/Nouakchott"}
	case "AEP":
		return Location{"Buenos Aires", "America/Argentina/Buenos_Aires"}
	case "AER":
		return Location{"Sochi", "Europe/Moscow"}
	case "AES":
		return Location{"Alesund", "Europe/Oslo"}
	case "AET":
		return Location{"Allakaket", "America/Anchorage"}
	case "AEU":
		return Location{"", "Asia/Tehran"}
	case "AEX":
		return Location{"Alexandria", "America/Chicago"}
	case "AEY":
		return Location{"Akureyri", "Atlantic/Reykjavik"}
	case "AFA":
		return Location{"San Rafael", "America/Argentina/Mendoza"}
	case "AFD":
		return Location{"Port Alfred", "Africa/Johannesburg"}
	case "AFE":
		return Location{"Kake", "America/Sitka"}
	case "AFF":
		return Location{"Colorado Springs", "America/Denver"}
	case "AFI":
		return Location{"Amalfi", "America/Bogota"}
	case "AFL":
		return Location{"Alta Floresta", "America/Cuiaba"}
	case "AFN":
		return Location{"Jaffrey", "America/New_York"}
	case "AFO":
		return Location{"Afton", "America/Denver"}
	case "AFR":
		return Location{"Afore", "Pacific/Port_Moresby"}
	case "AFS":
		return Location{"Zarafshan", "Asia/Samarkand"}
	case "AFT":
		return Location{"Bila", "Pacific/Guadalcanal"}
	case "AFW":
		return Location{"Fort Worth", "America/Chicago"}
	case "AFY":
		return Location{"Afyonkarahisar", "Europe/Istanbul"}
	case "AFZ":
		return Location{"Sabzevar", "Asia/Tehran"}
	case "AGA":
		return Location{"Agadir", "Africa/Casablanca"}
	case "AGB":
		return Location{"Augsburg", "Europe/Berlin"}
	case "AGC":
		return Location{"Pittsburgh", "America/New_York"}
	case "AGE":
		return Location{"Wangerooge", "Europe/Berlin"}
	case "AGF":
		return Location{"Agen/La Garenne", "Europe/Paris"}
	case "AGH":
		return Location{"Angelholm", "Europe/Stockholm"}
	case "AGI":
		return Location{"Wageningen Airport", "America/Paramaribo"}
	case "AGJ":
		return Location{"Aguni", "Asia/Tokyo"}
	case "AGL":
		return Location{"", "Pacific/Port_Moresby"}
	case "AGN":
		return Location{"Angoon", "America/Juneau"}
	case "AGO":
		return Location{"Magnolia", "America/Chicago"}
	case "AGP":
		return Location{"Malaga", "Europe/Madrid"}
	case "AGQ":
		return Location{"Agrinion", "Europe/Athens"}
	case "AGR":
		return Location{"", "Asia/Kolkata"}
	case "AGS":
		return Location{"Augusta", "America/New_York"}
	case "AGT":
		return Location{"Ciudad del Este", "America/Asuncion"}
	case "AGU":
		return Location{"Aguascalientes", "America/Mexico_City"}
	case "AGV":
		return Location{"Acarigua", "America/Caracas"}
	case "AGX":
		return Location{"", "Asia/Kolkata"}
	case "AGZ":
		return Location{"Aggeneys", "Africa/Johannesburg"}
	case "AHB":
		return Location{"Abha", "Asia/Riyadh"}
	case "AHC":
		return Location{"Herlong", "America/Los_Angeles"}
	case "AHE":
		return Location{"Ahe Atoll", "Pacific/Tahiti"}
	case "AHF":
		return Location{"Arapahoe", "America/Chicago"}
	case "AHH":
		return Location{"Amery", "America/Chicago"}
	case "AHI":
		return Location{"Amahai-Seram Island", "Asia/Jayapura"}
	case "AHJ":
		return Location{"Hongyuan", "Asia/Shanghai"}
	case "AHL":
		return Location{"Aishalton", "America/Guyana"}
	case "AHM":
		return Location{"Ashland", "America/Los_Angeles"}
	case "AHN":
		return Location{"Athens", "America/New_York"}
	case "AHO":
		return Location{"Alghero", "Europe/Rome"}
	case "AHS":
		return Location{"Ahuas", "America/Tegucigalpa"}
	case "AHU":
		return Location{"Al Hoceima", "Africa/Casablanca"}
	case "AHZ":
		return Location{"Bourg", "Europe/Paris"}
	case "AIA":
		return Location{"Alliance", "America/Denver"}
	case "AID":
		return Location{"Anderson", "America/Indiana/Indianapolis"}
	case "AIE":
		return Location{"Aiome", "Pacific/Port_Moresby"}
	case "AIF":
		return Location{"Assis", "America/Sao_Paulo"}
	case "AIG":
		return Location{"Yalinga", "Africa/Bangui"}
	case "AII":
		return Location{"Ali-Sabieh", "Africa/Djibouti"}
	case "AIK":
		return Location{"Aiken", "America/New_York"}
	case "AIN":
		return Location{"Wainwright", "America/Anchorage"}
	case "AIO":
		return Location{"Atlantic", "America/Chicago"}
	case "AIP":
		return Location{"", "Asia/Kolkata"}
	case "AIR":
		return Location{"Aripuanã", "America/Cuiaba"}
	case "AIS":
		return Location{"Arorae Island", "Pacific/Tarawa"}
	case "AIT":
		return Location{"Aitutaki", "Pacific/Rarotonga"}
	case "AIU":
		return Location{"Atiu Island", "Pacific/Rarotonga"}
	case "AIV":
		return Location{"Aliceville", "America/Chicago"}
	case "AIZ":
		return Location{"Kaiser Lake Ozark", "America/Chicago"}
	case "AJA":
		return Location{"Ajaccio/Napoleon Bonaparte", "Europe/Paris"}
	case "AJF":
		return Location{"Al-Jawf", "Asia/Riyadh"}
	case "AJI":
		return Location{"Agri", "Europe/Istanbul"}
	case "AJJ":
		return Location{"Akjoujt", "Africa/Nouakchott"}
	case "AJK":
		return Location{"Araak", "Asia/Tehran"}
	case "AJL":
		return Location{"Aizawl", "Asia/Kolkata"}
	case "AJN":
		return Location{"Ouani", "Indian/Comoro"}
	case "AJR":
		return Location{"Arvidsjaur", "Europe/Stockholm"}
	case "AJU":
		return Location{"Aracaju", "America/Maceio"}
	case "AJY":
		return Location{"Agadez", "Africa/Niamey"}
	case "AKA":
		return Location{"Ankang", "Asia/Shanghai"}
	case "AKB":
		return Location{"Atka", "America/Adak"}
	case "AKC":
		return Location{"Akron", "America/New_York"}
	case "AKD":
		return Location{"", "Asia/Kolkata"}
	case "AKE":
		return Location{"Akiéni", "Africa/Libreville"}
	case "AKF":
		return Location{"Kufra", "Africa/Tripoli"}
	case "AKH":
		return Location{"", "Asia/Riyadh"}
	case "AKI":
		return Location{"Akiak", "America/Anchorage"}
	case "AKJ":
		return Location{"Asahikawa", "Asia/Tokyo"}
	case "AKK":
		return Location{"Akhiok", "America/Anchorage"}
	case "AKL":
		return Location{"Auckland", "Pacific/Auckland"}
	case "AKN":
		return Location{"King Salmon", "America/Anchorage"}
	case "AKO":
		return Location{"Akron", "America/Denver"}
	case "AKP":
		return Location{"Anaktuvuk Pass", "America/Anchorage"}
	case "AKQ":
		return Location{"Menggala-Sumatra Island", "Asia/Jakarta"}
	case "AKR":
		return Location{"Akure", "Africa/Lagos"}
	case "AKS":
		return Location{"Auki", "Pacific/Guadalcanal"}
	case "AKT":
		return Location{"", "Asia/Nicosia"}
	case "AKU":
		return Location{"Aksu", "Asia/Shanghai"}
	case "AKV":
		return Location{"Akulivik", "America/Iqaluit"}
	case "AKW":
		return Location{"Omidiyeh", "Asia/Tehran"}
	case "AKX":
		return Location{"Aktyubinsk", "Asia/Aqtobe"}
	case "AKY":
		return Location{"Sittwe", "Asia/Yangon"}
	case "ALA":
		return Location{"Almaty", "Asia/Almaty"}
	case "ALB":
		return Location{"Albany", "America/New_York"}
	case "ALC":
		return Location{"Alicante", "Europe/Madrid"}
	case "ALD":
		return Location{"Fortaleza", "America/Lima"}
	case "ALF":
		return Location{"Alta", "Europe/Oslo"}
	case "ALG":
		return Location{"Algiers", "Africa/Algiers"}
	case "ALH":
		return Location{"Albany", "Australia/Perth"}
	case "ALI":
		return Location{"Alice", "America/Chicago"}
	case "ALJ":
		return Location{"Alexander Bay", "Africa/Johannesburg"}
	case "ALL":
		return Location{"Albenga", "Europe/Rome"}
	case "ALM":
		return Location{"Alamogordo", "America/Denver"}
	case "ALN":
		return Location{"Alton/St Louis", "America/Chicago"}
	case "ALO":
		return Location{"Waterloo", "America/Chicago"}
	case "ALP":
		return Location{"Aleppo", "Asia/Damascus"}
	case "ALQ":
		return Location{"Alegrete", "America/Sao_Paulo"}
	case "ALR":
		return Location{"Alexandra", "Pacific/Auckland"}
	case "ALS":
		return Location{"Alamosa", "America/Denver"}
	case "ALT":
		return Location{"Alenquer", "America/Santarem"}
	case "ALU":
		return Location{"Alula", "Africa/Mogadishu"}
	case "ALW":
		return Location{"Walla Walla", "America/Los_Angeles"}
	case "ALX":
		return Location{"Alexander City", "America/Chicago"}
	case "ALY":
		return Location{"Alexandria", "Africa/Cairo"}
	case "AMA":
		return Location{"Amarillo", "America/Chicago"}
	case "AMB":
		return Location{"", "Indian/Antananarivo"}
	case "AMC":
		return Location{"Am Timan", "Africa/Ndjamena"}
	case "AMD":
		return Location{"Ahmedabad", "Asia/Kolkata"}
	case "AMH":
		return Location{"", "Africa/Addis_Ababa"}
	case "AMI":
		return Location{"Mataram-Lombok Island", "Asia/Makassar"}
	case "AMJ":
		return Location{"Almenara", "America/Sao_Paulo"}
	case "AMK":
		return Location{"Durango", "America/Denver"}
	case "AMM":
		return Location{"Amman", "Asia/Amman"}
	case "AMN":
		return Location{"Alma", "America/Detroit"}
	case "AMO":
		return Location{"Mao", "Africa/Ndjamena"}
	case "AMP":
		return Location{"Ampanihy", "Indian/Antananarivo"}
	case "AMQ":
		return Location{"Ambon", "Asia/Jayapura"}
	case "AMS":
		return Location{"Amsterdam", "Europe/Amsterdam"}
	case "AMT":
		return Location{"", "Australia/Adelaide"}
	case "AMU":
		return Location{"Amanab", "Pacific/Port_Moresby"}
	case "AMV":
		return Location{"Amderma", "Europe/Moscow"}
	case "AMW":
		return Location{"Ames", "America/Chicago"}
	case "AMX":
		return Location{"", "Australia/Darwin"}
	case "AMZ":
		return Location{"Manurewa", "Pacific/Auckland"}
	case "ANB":
		return Location{"Anniston", "America/Chicago"}
	case "ANC":
		return Location{"Anchorage", "America/Anchorage"}
	case "AND":
		return Location{"Anderson", "America/New_York"}
	case "ANE":
		return Location{"Angers/Marce", "Europe/Paris"}
	case "ANF":
		return Location{"Antofagasta", "America/Santiago"}
	case "ANG":
		return Location{"Angouleme/Brie/Champniers", "Europe/Paris"}
	case "ANI":
		return Location{"Aniak", "America/Anchorage"}
	case "ANJ":
		return Location{"Zanaga", "Africa/Brazzaville"}
	case "ANK":
		return Location{"Ankara", "Europe/Istanbul"}
	case "ANM":
		return Location{"", "Indian/Antananarivo"}
	case "ANN":
		return Location{"Annette", "America/Metlakatla"}
	case "ANO":
		return Location{"Angoche", "Africa/Maputo"}
	case "ANP":
		return Location{"Annapolis", "America/New_York"}
	case "ANQ":
		return Location{"Angola", "America/Indiana/Indianapolis"}
	case "ANR":
		return Location{"Antwerp", "Europe/Brussels"}
	case "ANS":
		return Location{"Andahuaylas", "America/Lima"}
	case "ANU":
		return Location{"St. George", "America/Antigua"}
	case "ANV":
		return Location{"Anvik", "America/Anchorage"}
	case "ANW":
		return Location{"Ainsworth", "America/Chicago"}
	case "ANX":
		return Location{"Andenes", "Europe/Oslo"}
	case "ANY":
		return Location{"Anthony", "America/Chicago"}
	case "AOC":
		return Location{"Altenburg", "Europe/Berlin"}
	case "AOE":
		return Location{"Eskisehir", "Europe/Istanbul"}
	case "AOG":
		return Location{"Anshan", "Asia/Shanghai"}
	case "AOH":
		return Location{"Lima", "America/New_York"}
	case "AOI":
		return Location{"Ancona", "Europe/Rome"}
	case "AOJ":
		return Location{"Aomori", "Asia/Tokyo"}
	case "AOK":
		return Location{"Karpathos Island", "Europe/Athens"}
	case "AOL":
		return Location{"Paso de los Libres", "America/Argentina/Cordoba"}
	case "AOM":
		return Location{"Adam", "Asia/Muscat"}
	case "AOO":
		return Location{"Altoona", "America/New_York"}
	case "AOP":
		return Location{"Andoas", "America/Lima"}
	case "AOR":
		return Location{"Alor Satar", "Asia/Kuala_Lumpur"}
	case "AOT":
		return Location{"Aosta", "Europe/Rome"}
	case "AOU":
		return Location{"Attopeu", "Asia/Vientiane"}
	case "AOY":
		return Location{"Asaloyeh", "Asia/Tehran"}
	case "APA":
		return Location{"Denver", "America/Denver"}
	case "APB":
		return Location{"Apolo", "America/La_Paz"}
	case "APC":
		return Location{"Napa", "America/Los_Angeles"}
	case "APE":
		return Location{"San Juan Aposento", "America/Lima"}
	case "APF":
		return Location{"Naples", "America/New_York"}
	case "APG":
		return Location{"Aberdeen Proving Grounds(Aberdeen)", "America/New_York"}
	case "APH":
		return Location{"Fort A. P. Hill", "America/New_York"}
	case "API":
		return Location{"Apiay", "America/Bogota"}
	case "APJ":
		return Location{"Burang", "Asia/Shanghai"}
	case "APK":
		return Location{"Apataki", "Pacific/Tahiti"}
	case "APL":
		return Location{"Nampula", "Africa/Maputo"}
	case "APN":
		return Location{"Alpena", "America/Detroit"}
	case "APO":
		return Location{"Carepa", "America/Bogota"}
	case "APQ":
		return Location{"Arapiraca", "America/Maceio"}
	case "APS":
		return Location{"Anapolis", "America/Sao_Paulo"}
	case "APT":
		return Location{"Jasper", "America/Chicago"}
	case "APU":
		return Location{"Apucarana", "America/Sao_Paulo"}
	case "APV":
		return Location{"Apple Valley", "America/Los_Angeles"}
	case "APW":
		return Location{"Apia", "Pacific/Apia"}
	case "APX":
		return Location{"Arapongas", "America/Sao_Paulo"}
	case "APY":
		return Location{"Alto Parnaiba", "America/Fortaleza"}
	case "APZ":
		return Location{"Zapala", "America/Argentina/Salta"}
	case "AQA":
		return Location{"Araraquara", "America/Sao_Paulo"}
	case "AQB":
		return Location{"Santa Cruz del Quiche", "America/Guatemala"}
	case "AQG":
		return Location{"Anqing", "Asia/Shanghai"}
	case "AQI":
		return Location{"Qaisumah", "Asia/Riyadh"}
	case "AQJ":
		return Location{"Aqaba", "Asia/Amman"}
	case "AQM":
		return Location{"Ariquemes", "America/Porto_Velho"}
	case "AQP":
		return Location{"Arequipa", "America/Lima"}
	case "ARA":
		return Location{"New Iberia", "America/Chicago"}
	case "ARB":
		return Location{"Ann Arbor", "America/Detroit"}
	case "ARC":
		return Location{"Arctic Village", "America/Anchorage"}
	case "ARD":
		return Location{"Alor Island", "Asia/Makassar"}
	case "ARE":
		return Location{"Arecibo", "America/Puerto_Rico"}
	case "ARG":
		return Location{"Walnut Ridge", "America/Chicago"}
	case "ARH":
		return Location{"Archangelsk", "Europe/Moscow"}
	case "ARI":
		return Location{"Arica", "America/Lima"}
	case "ARJ":
		return Location{"Arso-Papua Island", "Asia/Jayapura"}
	case "ARK":
		return Location{"Arusha", "Africa/Dar_es_Salaam"}
	case "ARL":
		return Location{"Arly", "Africa/Ouagadougou"}
	case "ARM":
		return Location{"Armidale", "Australia/Sydney"}
	case "ARN":
		return Location{"Stockholm", "Europe/Stockholm"}
	case "ARR":
		return Location{"Alto Rio Senguerr", "America/Argentina/Catamarca"}
	case "ARS":
		return Location{"Aragarcas", "America/Cuiaba"}
	case "ART":
		return Location{"Watertown", "America/New_York"}
	case "ARU":
		return Location{"Aracatuba", "America/Sao_Paulo"}
	case "ARV":
		return Location{"Minocqua-Woodruff", "America/Chicago"}
	case "ARW":
		return Location{"Arad", "Europe/Bucharest"}
	case "ARY":
		return Location{"", "Australia/Melbourne"}
	case "ARZ":
		return Location{"N'zeto", "Africa/Luanda"}
	case "ASA":
		return Location{"Asab", "Africa/Asmara"}
	case "ASB":
		return Location{"Ashgabat", "Asia/Ashgabat"}
	case "ASC":
		return Location{"Ascension de Guarayos", "America/La_Paz"}
	case "ASD":
		return Location{"", "America/Nassau"}
	case "ASE":
		return Location{"Aspen", "America/Denver"}
	case "ASF":
		return Location{"Astrakhan", "Europe/Astrakhan"}
	case "ASG":
		return Location{"", "Pacific/Auckland"}
	case "ASH":
		return Location{"Nashua", "America/New_York"}
	case "ASI":
		return Location{"Ascension Island", "Atlantic/St_Helena"}
	case "ASJ":
		return Location{"Amami", "Asia/Tokyo"}
	case "ASK":
		return Location{"Yamoussoukro", "Africa/Abidjan"}
	case "ASL":
		return Location{"Marshall", "America/Chicago"}
	case "ASM":
		return Location{"Asmara", "Africa/Asmara"}
	case "ASN":
		return Location{"Talladega", "America/Chicago"}
	case "ASO":
		return Location{"Asosa", "Africa/Addis_Ababa"}
	case "ASP":
		return Location{"Alice Springs", "Australia/Darwin"}
	case "ASQ":
		return Location{"Austin", "America/Los_Angeles"}
	case "ASR":
		return Location{"Kayseri", "Europe/Istanbul"}
	case "AST":
		return Location{"Astoria", "America/Los_Angeles"}
	case "ASU":
		return Location{"Asuncion", "America/Asuncion"}
	case "ASV":
		return Location{"Amboseli National Park", "Africa/Nairobi"}
	case "ASW":
		return Location{"Aswan", "Africa/Cairo"}
	case "ASX":
		return Location{"Ashland", "America/Chicago"}
	case "ASY":
		return Location{"Ashley", "America/Chicago"}
	case "ATA":
		return Location{"Anta", "America/Lima"}
	case "ATB":
		return Location{"Atbara", "Africa/Khartoum"}
	case "ATC":
		return Location{"Arthur's Town", "America/Nassau"}
	case "ATD":
		return Location{"Atoifi", "Pacific/Guadalcanal"}
	case "ATE":
		return Location{"Antlers", "America/Chicago"}
	case "ATF":
		return Location{"Ambato", "America/Guayaquil"}
	case "ATH":
		return Location{"Athens", "Europe/Athens"}
	case "ATI":
		return Location{"Artigas", "America/Montevideo"}
	case "ATJ":
		return Location{"Antsirabe", "Indian/Antananarivo"}
	case "ATK":
		return Location{"Atqasuk", "America/Anchorage"}
	case "ATL":
		return Location{"Atlanta", "America/New_York"}
	case "ATM":
		return Location{"Altamira", "America/Santarem"}
	case "ATO":
		return Location{"Athens/Albany", "America/New_York"}
	case "ATP":
		return Location{"Aitape", "Pacific/Port_Moresby"}
	case "ATQ":
		return Location{"Amritsar", "Asia/Kolkata"}
	case "ATR":
		return Location{"Atar", "Africa/Nouakchott"}
	case "ATS":
		return Location{"Artesia", "America/Denver"}
	case "ATU":
		return Location{"Attu", "America/Adak"}
	case "ATV":
		return Location{"Ati", "Africa/Ndjamena"}
	case "ATW":
		return Location{"Appleton", "America/Chicago"}
	case "ATY":
		return Location{"Watertown", "America/Chicago"}
	case "ATZ":
		return Location{"Assiut", "Africa/Cairo"}
	case "AUA":
		return Location{"Oranjestad", "America/Aruba"}
	case "AUC":
		return Location{"Arauca", "America/Bogota"}
	case "AUD":
		return Location{"", "Australia/Brisbane"}
	case "AUF":
		return Location{"Auxerre/Branches", "Europe/Paris"}
	case "AUG":
		return Location{"Augusta", "America/New_York"}
	case "AUH":
		return Location{"Abu Dhabi", "Asia/Dubai"}
	case "AUJ":
		return Location{"Ambunti", "Pacific/Port_Moresby"}
	case "AUK":
		return Location{"Alakanuk", "America/Nome"}
	case "AUM":
		return Location{"Austin", "America/Chicago"}
	case "AUN":
		return Location{"Auburn", "America/Los_Angeles"}
	case "AUO":
		return Location{"Auburn", "America/Chicago"}
	case "AUQ":
		return Location{"", "Pacific/Marquesas"}
	case "AUR":
		return Location{"Aurillac", "Europe/Paris"}
	case "AUS":
		return Location{"Austin", "America/Chicago"}
	case "AUT":
		return Location{"Atauro", "Asia/Dili"}
	case "AUU":
		return Location{"", "Australia/Brisbane"}
	case "AUW":
		return Location{"Wausau", "America/Chicago"}
	case "AUX":
		return Location{"Araguaina", "America/Araguaina"}
	case "AUY":
		return Location{"Anelghowhat", "Pacific/Efate"}
	case "AUZ":
		return Location{"Chicago/Aurora", "America/Chicago"}
	case "AVA":
		return Location{"Anshun", "Asia/Shanghai"}
	case "AVB":
		return Location{"Aviano", "Europe/Rome"}
	case "AVG":
		return Location{"", "Australia/Darwin"}
	case "AVI":
		return Location{"Ciego de Avila", "America/Havana"}
	case "AVK":
		return Location{"Arvaikheer", "Asia/Ulaanbaatar"}
	case "AVL":
		return Location{"Asheville", "America/New_York"}
	case "AVN":
		return Location{"Avignon/Caumont", "Europe/Paris"}
	case "AVO":
		return Location{"Avon Park", "America/New_York"}
	case "AVP":
		return Location{"Wilkes-Barre/Scranton", "America/New_York"}
	case "AVR":
		return Location{"Alverca", "Europe/Lisbon"}
	case "AVU":
		return Location{"", "Pacific/Guadalcanal"}
	case "AVV":
		return Location{"Melbourne", "Australia/Melbourne"}
	case "AVW":
		return Location{"Tucson", "America/Phoenix"}
	case "AVX":
		return Location{"Avalon", "America/Los_Angeles"}
	case "AWA":
		return Location{"Awassa", "Africa/Addis_Ababa"}
	case "AWB":
		return Location{"Awaba", "Pacific/Port_Moresby"}
	case "AWD":
		return Location{"Aniwa", "Pacific/Efate"}
	case "AWK":
		return Location{"Wake Island", "Pacific/Wake"}
	case "AWM":
		return Location{"West Memphis", "America/Chicago"}
	case "AWN":
		return Location{"", "Australia/Adelaide"}
	case "AWP":
		return Location{"", "Australia/Darwin"}
	case "AWZ":
		return Location{"Ahwaz", "Asia/Tehran"}
	case "AXA":
		return Location{"The Valley", "America/Anguilla"}
	case "AXC":
		return Location{"", "Australia/Brisbane"}
	case "AXD":
		return Location{"Alexandroupolis", "Europe/Athens"}
	case "AXE":
		return Location{"Xanxere", "America/Sao_Paulo"}
	case "AXF":
		return Location{"Bayanhot", "Asia/Shanghai"}
	case "AXG":
		return Location{"Algona", "America/Chicago"}
	case "AXJ":
		return Location{"", "Asia/Tokyo"}
	case "AXK":
		return Location{"", "Asia/Aden"}
	case "AXL":
		return Location{"", "Australia/Darwin"}
	case "AXM":
		return Location{"Armenia", "America/Bogota"}
	case "AXN":
		return Location{"Alexandria", "America/Chicago"}
	case "AXP":
		return Location{"Spring Point", "America/Nassau"}
	case "AXR":
		return Location{"", "Pacific/Tahiti"}
	case "AXS":
		return Location{"Altus", "America/Chicago"}
	case "AXT":
		return Location{"Akita", "Asia/Tokyo"}
	case "AXU":
		return Location{"", "Africa/Addis_Ababa"}
	case "AXV":
		return Location{"Wapakoneta", "America/New_York"}
	case "AXX":
		return Location{"Angel Fire", "America/Denver"}
	case "AYG":
		return Location{"San Vicente Del Caguan", "America/Bogota"}
	case "AYJ":
		return Location{"Ayodhya", "Asia/Kolkata"}
	case "AYK":
		return Location{"Arkalyk", "Asia/Qostanay"}
	case "AYL":
		return Location{"", "Australia/Darwin"}
	case "AYN":
		return Location{"Anyang", "Asia/Shanghai"}
	case "AYO":
		return Location{"Ayolas", "America/Asuncion"}
	case "AYP":
		return Location{"Ayacucho", "America/Lima"}
	case "AYQ":
		return Location{"Ayers Rock", "Australia/Darwin"}
	case "AYR":
		return Location{"", "Australia/Brisbane"}
	case "AYS":
		return Location{"Waycross", "America/New_York"}
	case "AYT":
		return Location{"Antalya", "Europe/Istanbul"}
	case "AYU":
		return Location{"Aiyura Valley", "Pacific/Port_Moresby"}
	case "AYX":
		return Location{"Atalaya", "America/Lima"}
	case "AZA":
		return Location{"Phoenix", "America/Phoenix"}
	case "AZD":
		return Location{"Yazd", "Asia/Tehran"}
	case "AZH":
		return Location{"Azamgarh", "Asia/Kolkata"}
	case "AZI":
		return Location{"", "Asia/Dubai"}
	case "AZL":
		return Location{"Sapezal", "America/Cuiaba"}
	case "AZN":
		return Location{"Andijan", "Asia/Tashkent"}
	case "AZO":
		return Location{"Kalamazoo", "America/Detroit"}
	case "AZP":
		return Location{"", "America/Mexico_City"}
	case "AZR":
		return Location{"", "Africa/Algiers"}
	case "AZS":
		return Location{"Samana", "America/Santo_Domingo"}
	case "AZZ":
		return Location{"Ambriz", "Africa/Luanda"}
	case "BAA":
		return Location{"Bialla", "Pacific/Port_Moresby"}
	case "BAB":
		return Location{"Marysville", "America/Los_Angeles"}
	case "BAD":
		return Location{"Bossier City", "America/Chicago"}
	case "BAE":
		return Location{"Le Castellet", "Europe/Paris"}
	case "BAF":
		return Location{"Westfield/Springfield", "America/New_York"}
	case "BAG":
		return Location{"Baguio City", "Asia/Manila"}
	case "BAH":
		return Location{"Manama", "Asia/Bahrain"}
	case "BAI":
		return Location{"Punta Arenas", "America/Costa_Rica"}
	case "BAL":
		return Location{"Batman", "Europe/Istanbul"}
	case "BAM":
		return Location{"Battle Mountain", "America/Los_Angeles"}
	case "BAN":
		return Location{"Basongo", "Africa/Lubumbashi"}
	case "BAO":
		return Location{"Ban Mak Khaen", "Asia/Bangkok"}
	case "BAQ":
		return Location{"Barranquilla", "America/Bogota"}
	case "BAR":
		return Location{"Qionghai", "Asia/Shanghai"}
	case "BAS":
		return Location{"Ballalae", "Pacific/Guadalcanal"}
	case "BAT":
		return Location{"Barretos", "America/Sao_Paulo"}
	case "BAU":
		return Location{"Bauru", "America/Sao_Paulo"}
	case "BAV":
		return Location{"Baotou", "Asia/Shanghai"}
	case "BAX":
		return Location{"Barnaul", "Asia/Barnaul"}
	case "BAY":
		return Location{"Baia Mare", "Europe/Bucharest"}
	case "BAZ":
		return Location{"Barcelos", "America/Manaus"}
	case "BBA":
		return Location{"Balmaceda", "America/Santiago"}
	case "BBB":
		return Location{"Benson", "America/Chicago"}
	case "BBC":
		return Location{"Bay City", "America/Chicago"}
	case "BBD":
		return Location{"Brady", "America/Chicago"}
	case "BBE":
		return Location{"Big Bell", "Australia/Perth"}
	case "BBG":
		return Location{"Butaritari Atoll", "Pacific/Tarawa"}
	case "BBH":
		return Location{"", "Europe/Berlin"}
	case "BBI":
		return Location{"Bhubaneswar", "Asia/Kolkata"}
	case "BBJ":
		return Location{"Bitburg", "Europe/Berlin"}
	case "BBK":
		return Location{"Kasane", "Africa/Gaborone"}
	case "BBL":
		return Location{"", "Australia/Brisbane"}
	case "BBM":
		return Location{"Battambang", "Asia/Phnom_Penh"}
	case "BBN":
		return Location{"Bario", "Asia/Kuching"}
	case "BBO":
		return Location{"Berbera", "Africa/Mogadishu"}
	case "BBP":
		return Location{"Bembridge", "Europe/London"}
	case "BBQ":
		return Location{"Codrington", "America/Antigua"}
	case "BBR":
		return Location{"Basse Terre", "America/Guadeloupe"}
	case "BBS":
		return Location{"Yateley", "Europe/London"}
	case "BBT":
		return Location{"Berberati", "Africa/Bangui"}
	case "BBU":
		return Location{"Bucharest", "Europe/Bucharest"}
	case "BBV":
		return Location{"Grand-Bereby", "Africa/Abidjan"}
	case "BBW":
		return Location{"Broken Bow", "America/Chicago"}
	case "BBX":
		return Location{"Philadelphia", "America/New_York"}
	case "BBY":
		return Location{"Bambari", "Africa/Bangui"}
	case "BBZ":
		return Location{"Zambezi", "Africa/Lusaka"}
	case "BCA":
		return Location{"Baracoa", "America/Havana"}
	case "BCB":
		return Location{"Blacksburg", "America/New_York"}
	case "BCD":
		return Location{"Bacolod City", "Asia/Manila"}
	case "BCE":
		return Location{"Bryce Canyon", "America/Denver"}
	case "BCF":
		return Location{"Bouca", "Africa/Bangui"}
	case "BCG":
		return Location{"Bemichi", "America/Guyana"}
	case "BCH":
		return Location{"Baucau", "Asia/Dili"}
	case "BCI":
		return Location{"Barcaldine", "Australia/Brisbane"}
	case "BCL":
		return Location{"Pococi", "America/Costa_Rica"}
	case "BCM":
		return Location{"Bacau", "Europe/Bucharest"}
	case "BCN":
		return Location{"Barcelona", "Europe/Madrid"}
	case "BCO":
		return Location{"Baco", "Africa/Addis_Ababa"}
	case "BCR":
		return Location{"Boca Do Acre", "America/Manaus"}
	case "BCS":
		return Location{"Belle Chasse", "America/Chicago"}
	case "BCT":
		return Location{"Boca Raton", "America/New_York"}
	case "BCX":
		return Location{"Beloretsk", "Asia/Yekaterinburg"}
	case "BDA":
		return Location{"Hamilton", "Atlantic/Bermuda"}
	case "BDB":
		return Location{"Bundaberg", "Australia/Brisbane"}
	case "BDC":
		return Location{"Barra Do Corda", "America/Fortaleza"}
	case "BDD":
		return Location{"", "Australia/Brisbane"}
	case "BDE":
		return Location{"Baudette", "America/Chicago"}
	case "BDF":
		return Location{"Bradford", "America/Chicago"}
	case "BDG":
		return Location{"Blanding", "America/Denver"}
	case "BDH":
		return Location{"Bandar Lengeh", "Asia/Tehran"}
	case "BDI":
		return Location{"Bird Island", "Indian/Mahe"}
	case "BDJ":
		return Location{"Banjarmasin-Borneo Island", "Asia/Makassar"}
	case "BDK":
		return Location{"Bondoukou", "Africa/Abidjan"}
	case "BDL":
		return Location{"Hartford", "America/New_York"}
	case "BDM":
		return Location{"Bandirma", "Europe/Istanbul"}
	case "BDN":
		return Location{"Badin", "Asia/Karachi"}
	case "BDO":
		return Location{"Bandung-Java Island", "Asia/Jakarta"}
	case "BDP":
		return Location{"Bhadrapur", "Asia/Kathmandu"}
	case "BDQ":
		return Location{"Vadodara", "Asia/Kolkata"}
	case "BDR":
		return Location{"Bridgeport", "America/New_York"}
	case "BDS":
		return Location{"Brindisi", "Europe/Rome"}
	case "BDT":
		return Location{"", "Africa/Kinshasa"}
	case "BDU":
		return Location{"Malselv", "Europe/Oslo"}
	case "BDV":
		return Location{"Moba", "Africa/Lubumbashi"}
	case "BDW":
		return Location{"", "Australia/Perth"}
	case "BDX":
		return Location{"Broadus", "America/Denver"}
	case "BDY":
		return Location{"Bandon", "America/Los_Angeles"}
	case "BDZ":
		return Location{"", "Pacific/Port_Moresby"}
	case "BEB":
		return Location{"Balivanich", "Europe/London"}
	case "BEC":
		return Location{"Wichita", "America/Chicago"}
	case "BED":
		return Location{"Bedford", "America/New_York"}
	case "BEF":
		return Location{"Bluefileds", "America/Managua"}
	case "BEG":
		return Location{"Belgrad", "Europe/Belgrade"}
	case "BEH":
		return Location{"Benton Harbor", "America/Detroit"}
	case "BEI":
		return Location{"Beica", "Africa/Addis_Ababa"}
	case "BEJ":
		return Location{"Tanjung Redep-Borneo Island", "Asia/Makassar"}
	case "BEK":
		return Location{"", "Asia/Kolkata"}
	case "BEL":
		return Location{"Belem", "America/Belem"}
	case "BEM":
		return Location{"", "Africa/Casablanca"}
	case "BEN":
		return Location{"Benghazi", "Africa/Tripoli"}
	case "BEO":
		return Location{"", "Australia/Sydney"}
	case "BEP":
		return Location{"Bellary", "Asia/Kolkata"}
	case "BEQ":
		return Location{"Thetford", "Europe/London"}
	case "BER":
		return Location{"Berlin", "Europe/Berlin"}
	case "BES":
		return Location{"Brest/Guipavas", "Europe/Paris"}
	case "BET":
		return Location{"Bethel", "America/Anchorage"}
	case "BEU":
		return Location{"", "Australia/Brisbane"}
	case "BEV":
		return Location{"Beersheva", "Asia/Jerusalem"}
	case "BEW":
		return Location{"Beira", "Africa/Maputo"}
	case "BEX":
		return Location{"Benson", "Europe/London"}
	case "BEY":
		return Location{"Beirut", "Asia/Beirut"}
	case "BEZ":
		return Location{"Beru", "Pacific/Tarawa"}
	case "BFA":
		return Location{"Bahia Negra", "America/Asuncion"}
	case "BFD":
		return Location{"Bradford", "America/New_York"}
	case "BFE":
		return Location{"Bielefeld", "Europe/Berlin"}
	case "BFF":
		return Location{"Scottsbluff", "America/Denver"}
	case "BFG":
		return Location{"Glen Canyon Natl Rec Area", "America/Denver"}
	case "BFH":
		return Location{"Curitiba", "America/Sao_Paulo"}
	case "BFI":
		return Location{"Seattle", "America/Los_Angeles"}
	case "BFJ":
		return Location{"Bijie", "Asia/Shanghai"}
	case "BFK":
		return Location{"Aurora", "America/Denver"}
	case "BFL":
		return Location{"Bakersfield", "America/Los_Angeles"}
	case "BFM":
		return Location{"Mobile", "America/Chicago"}
	case "BFN":
		return Location{"Bloemfontain", "Africa/Johannesburg"}
	case "BFO":
		return Location{"Chiredzi", "Africa/Harare"}
	case "BFP":
		return Location{"Beaver Falls", "America/New_York"}
	case "BFR":
		return Location{"Bedford", "America/Indiana/Indianapolis"}
	case "BFS":
		return Location{"Belfast", "Europe/London"}
	case "BFT":
		return Location{"Beaufort", "America/New_York"}
	case "BFU":
		return Location{"Bengbu", "Asia/Shanghai"}
	case "BFV":
		return Location{"", "Asia/Bangkok"}
	case "BFW":
		return Location{"", "Africa/Algiers"}
	case "BFX":
		return Location{"Bafoussam", "Africa/Douala"}
	case "BGA":
		return Location{"Bucaramanga", "America/Bogota"}
	case "BGB":
		return Location{"Booue", "Africa/Libreville"}
	case "BGC":
		return Location{"", "Europe/Lisbon"}
	case "BGD":
		return Location{"Borger", "America/Chicago"}
	case "BGE":
		return Location{"Bainbridge", "America/New_York"}
	case "BGF":
		return Location{"Bangui", "Africa/Bangui"}
	case "BGG":
		return Location{"Bingöl", "Europe/Istanbul"}
	case "BGH":
		return Location{"Boghe", "Africa/Nouakchott"}
	case "BGI":
		return Location{"Bridgetown", "America/Barbados"}
	case "BGJ":
		return Location{"Borgarfjordur eystri", "Atlantic/Reykjavik"}
	case "BGL":
		return Location{"Baglung", "Asia/Kathmandu"}
	case "BGM":
		return Location{"Binghamton", "America/New_York"}
	case "BGN":
		return Location{"Belaya Gora", "Asia/Srednekolymsk"}
	case "BGO":
		return Location{"Bergen", "Europe/Oslo"}
	case "BGQ":
		return Location{"Big Lake", "America/Anchorage"}
	case "BGR":
		return Location{"Bangor", "America/New_York"}
	case "BGS":
		return Location{"Big Spring", "America/Chicago"}
	case "BGT":
		return Location{"Bagdad", "America/Phoenix"}
	case "BGU":
		return Location{"Bangassou", "Africa/Bangui"}
	case "BGV":
		return Location{"Bento Goncalves", "America/Sao_Paulo"}
	case "BGW":
		return Location{"Baghdad", "Asia/Baghdad"}
	case "BGX":
		return Location{"Bage", "America/Sao_Paulo"}
	case "BGY":
		return Location{"Bergamo", "Europe/Rome"}
	case "BGZ":
		return Location{"Braga", "Europe/Lisbon"}
	case "BHA":
		return Location{"Bahia de Caraquez", "America/Guayaquil"}
	case "BHB":
		return Location{"Bar Harbor", "America/New_York"}
	case "BHD":
		return Location{"Belfast", "Europe/London"}
	case "BHE":
		return Location{"Blenheim", "Pacific/Auckland"}
	case "BHF":
		return Location{"Bahia Solano", "America/Bogota"}
	case "BHG":
		return Location{"Brus Laguna", "America/Tegucigalpa"}
	case "BHH":
		return Location{"", "Asia/Riyadh"}
	case "BHI":
		return Location{"Bahia Blanca", "America/Argentina/Buenos_Aires"}
	case "BHJ":
		return Location{"Bhuj", "Asia/Kolkata"}
	case "BHK":
		return Location{"Bukhara", "Asia/Samarkand"}
	case "BHM":
		return Location{"Birmingham", "America/Chicago"}
	case "BHN":
		return Location{"", "Asia/Aden"}
	case "BHO":
		return Location{"Bhopal", "Asia/Kolkata"}
	case "BHP":
		return Location{"Bhojpur", "Asia/Kathmandu"}
	case "BHQ":
		return Location{"Broken Hill", "Australia/Broken_Hill"}
	case "BHR":
		return Location{"Bharatpur", "Asia/Kathmandu"}
	case "BHS":
		return Location{"Bathurst", "Australia/Sydney"}
	case "BHU":
		return Location{"Bhavnagar", "Asia/Kolkata"}
	case "BHV":
		return Location{"Bahawalpur", "Asia/Karachi"}
	case "BHW":
		return Location{"Bhagatanwala", "Asia/Karachi"}
	case "BHX":
		return Location{"Birmingham", "Europe/London"}
	case "BHY":
		return Location{"Beihai", "Asia/Shanghai"}
	case "BIA":
		return Location{"Bastia/Poretta", "Europe/Paris"}
	case "BIB":
		return Location{"Baidoa", "Africa/Mogadishu"}
	case "BID":
		return Location{"Block Island", "America/New_York"}
	case "BIE":
		return Location{"Beatrice", "America/Chicago"}
	case "BIF":
		return Location{"Fort Bliss/El Paso", "America/Denver"}
	case "BIG":
		return Location{"Delta Junction Ft Greely", "America/Anchorage"}
	case "BIH":
		return Location{"Bishop", "America/Los_Angeles"}
	case "BIK":
		return Location{"Biak-Supiori Island", "Asia/Jayapura"}
	case "BIL":
		return Location{"Billings", "America/Denver"}
	case "BIM":
		return Location{"South Bimini", "America/Nassau"}
	case "BIN":
		return Location{"Bamiyan", "Asia/Kabul"}
	case "BIO":
		return Location{"Bilbao", "Europe/Madrid"}
	case "BIP":
		return Location{"", "Australia/Brisbane"}
	case "BIQ":
		return Location{"Biarritz/Anglet/Bayonne", "Europe/Paris"}
	case "BIR":
		return Location{"Biratnagar", "Asia/Kathmandu"}
	case "BIS":
		return Location{"Bismarck", "America/Chicago"}
	case "BIT":
		return Location{"Baitadi", "Asia/Kathmandu"}
	case "BIU":
		return Location{"Bildudalur", "Atlantic/Reykjavik"}
	case "BIV":
		return Location{"Bria", "Africa/Bangui"}
	case "BIW":
		return Location{"", "Australia/Perth"}
	case "BIX":
		return Location{"Biloxi", "America/Chicago"}
	case "BIY":
		return Location{"Bisho", "Africa/Johannesburg"}
	case "BJA":
		return Location{"Bejaia", "Africa/Algiers"}
	case "BJB":
		return Location{"Bojnord", "Asia/Tehran"}
	case "BJC":
		return Location{"Denver", "America/Denver"}
	case "BJD":
		return Location{"Bakkafjordur", "Atlantic/Reykjavik"}
	case "BJF":
		return Location{"Batsfjord", "Europe/Oslo"}
	case "BJH":
		return Location{"Bajhang", "Asia/Kathmandu"}
	case "BJI":
		return Location{"Bemidji", "America/Chicago"}
	case "BJJ":
		return Location{"Wooster", "America/New_York"}
	case "BJK":
		return Location{"Maikoor Island", "Asia/Jayapura"}
	case "BJL":
		return Location{"Banjul", "Africa/Banjul"}
	case "BJM":
		return Location{"Bujumbura", "Africa/Bujumbura"}
	case "BJO":
		return Location{"Bermejo", "America/Argentina/Salta"}
	case "BJP":
		return Location{"Braganca Paulista", "America/Sao_Paulo"}
	case "BJR":
		return Location{"Bahir Dar", "Africa/Addis_Ababa"}
	case "BJU":
		return Location{"Bajura", "Asia/Kathmandu"}
	case "BJV":
		return Location{"Bodrum", "Europe/Istanbul"}
	case "BJW":
		return Location{"Bajawa", "Asia/Makassar"}
	case "BJX":
		return Location{"Silao", "America/Mexico_City"}
	case "BJY":
		return Location{"Batajnica", "Europe/Belgrade"}
	case "BJZ":
		return Location{"Badajoz", "Europe/Madrid"}
	case "BKA":
		return Location{"Moscow", "Europe/Moscow"}
	case "BKB":
		return Location{"Bikaner", "Asia/Kolkata"}
	case "BKC":
		return Location{"Buckland", "America/Anchorage"}
	case "BKD":
		return Location{"Breckenridge", "America/Chicago"}
	case "BKE":
		return Location{"Baker City", "America/Los_Angeles"}
	case "BKG":
		return Location{"Branson", "America/Chicago"}
	case "BKH":
		return Location{"Kekaha", "Pacific/Honolulu"}
	case "BKI":
		return Location{"Kota Kinabalu", "Asia/Kuching"}
	case "BKJ":
		return Location{"Boke", "Africa/Conakry"}
	case "BKK":
		return Location{"Bangkok", "Asia/Bangkok"}
	case "BKL":
		return Location{"Cleveland", "America/New_York"}
	case "BKM":
		return Location{"Bakalalan", "Asia/Kuching"}
	case "BKN":
		return Location{"Jebel", "Asia/Ashgabat"}
	case "BKO":
		return Location{"Senou", "Africa/Bamako"}
	case "BKP":
		return Location{"", "Australia/Brisbane"}
	case "BKQ":
		return Location{"Blackall", "Australia/Brisbane"}
	case "BKR":
		return Location{"Bokoro", "Africa/Ndjamena"}
	case "BKS":
		return Location{"Bengkulu-Sumatra Island", "Asia/Jakarta"}
	case "BKT":
		return Location{"Blackstone", "America/New_York"}
	case "BKU":
		return Location{"Betioky", "Indian/Antananarivo"}
	case "BKW":
		return Location{"Beckley", "America/New_York"}
	case "BKX":
		return Location{"Brookings", "America/Chicago"}
	case "BKY":
		return Location{"", "Africa/Lubumbashi"}
	case "BKZ":
		return Location{"Bukoba", "Africa/Dar_es_Salaam"}
	case "BLA":
		return Location{"Barcelona", "America/Caracas"}
	case "BLB":
		return Location{"Panama City", "America/Panama"}
	case "BLC":
		return Location{"Bali", "Africa/Douala"}
	case "BLD":
		return Location{"Boulder City", "America/Los_Angeles"}
	case "BLE":
		return Location{"", "Europe/Stockholm"}
	case "BLF":
		return Location{"Bluefield", "America/New_York"}
	case "BLG":
		return Location{"Belaga", "Asia/Kuching"}
	case "BLH":
		return Location{"Blythe", "America/Los_Angeles"}
	case "BLI":
		return Location{"Bellingham", "America/Los_Angeles"}
	case "BLJ":
		return Location{"Batna", "Africa/Algiers"}
	case "BLK":
		return Location{"Blackpool", "Europe/London"}
	case "BLL":
		return Location{"Billund", "Europe/Copenhagen"}
	case "BLM":
		return Location{"Belmar/Farmingdale", "America/New_York"}
	case "BLN":
		return Location{"", "Australia/Melbourne"}
	case "BLO":
		return Location{"Blonduos", "Atlantic/Reykjavik"}
	case "BLP":
		return Location{"Bellavista", "America/Lima"}
	case "BLQ":
		return Location{"Bologna", "Europe/Rome"}
	case "BLR":
		return Location{"Bangalore", "Asia/Kolkata"}
	case "BLS":
		return Location{"", "Australia/Brisbane"}
	case "BLT":
		return Location{"", "Australia/Brisbane"}
	case "BLU":
		return Location{"Emigrant Gap", "America/Los_Angeles"}
	case "BLV":
		return Location{"Belleville", "America/Chicago"}
	case "BLX":
		return Location{"Belluno", "Europe/Rome"}
	case "BLY":
		return Location{"Belmullet", "Europe/Dublin"}
	case "BLZ":
		return Location{"Blantyre", "Africa/Blantyre"}
	case "BMA":
		return Location{"Stockholm", "Europe/Stockholm"}
	case "BMB":
		return Location{"Bumbar", "Africa/Kinshasa"}
	case "BMC":
		return Location{"Brigham City", "America/Denver"}
	case "BMD":
		return Location{"Belo sur Tsiribihina", "Indian/Antananarivo"}
	case "BME":
		return Location{"Broome", "Australia/Perth"}
	case "BMF":
		return Location{"Bakouma", "Africa/Bangui"}
	case "BMG":
		return Location{"Bloomington", "America/Indiana/Indianapolis"}
	case "BMI":
		return Location{"Bloomington-Normal", "America/Chicago"}
	case "BMJ":
		return Location{"Baramita", "America/Guyana"}
	case "BMK":
		return Location{"Borkum", "Europe/Berlin"}
	case "BML":
		return Location{"Berlin", "America/New_York"}
	case "BMM":
		return Location{"Bitam", "Africa/Libreville"}
	case "BMN":
		return Location{"Bamarni", "Asia/Baghdad"}
	case "BMO":
		return Location{"Banmaw", "Asia/Yangon"}
	case "BMP":
		return Location{"", "Australia/Brisbane"}
	case "BMR":
		return Location{"Baltrum", "Europe/Berlin"}
	case "BMS":
		return Location{"Brumado", "America/Bahia"}
	case "BMT":
		return Location{"Beaumont", "America/Chicago"}
	case "BMU":
		return Location{"Bima-Sumbawa Island", "Asia/Makassar"}
	case "BMV":
		return Location{"Buon Ma Thuot", "Asia/Ho_Chi_Minh"}
	case "BMW":
		return Location{"Bordj Badji Mokhtar", "Africa/Algiers"}
	case "BMX":
		return Location{"Big Mountain", "America/Anchorage"}
	case "BMY":
		return Location{"Waala", "Pacific/Noumea"}
	case "BNA":
		return Location{"Nashville", "America/Chicago"}
	case "BNB":
		return Location{"Boende", "Africa/Kinshasa"}
	case "BNC":
		return Location{"Beni", "Africa/Lubumbashi"}
	case "BND":
		return Location{"Bandar Abbas", "Asia/Tehran"}
	case "BNE":
		return Location{"Brisbane", "Australia/Brisbane"}
	case "BNG":
		return Location{"Banning", "America/Los_Angeles"}
	case "BNI":
		return Location{"Benin", "Africa/Lagos"}
	case "BNJ":
		return Location{"Bonn", "Europe/Berlin"}
	case "BNK":
		return Location{"Ballina", "Australia/Sydney"}
	case "BNL":
		return Location{"Barnwell", "America/New_York"}
	case "BNN":
		return Location{"Bronnoy", "Europe/Oslo"}
	case "BNO":
		return Location{"Burns", "America/Los_Angeles"}
	case "BNP":
		return Location{"Bannu", "Asia/Karachi"}
	case "BNR":
		return Location{"Banfora", "Africa/Ouagadougou"}
	case "BNS":
		return Location{"Barinas", "America/Caracas"}
	case "BNU":
		return Location{"Blumenau", "America/Sao_Paulo"}
	case "BNW":
		return Location{"Boone", "America/Chicago"}
	case "BNX":
		return Location{"Banja Luka", "Europe/Sarajevo"}
	case "BNY":
		return Location{"Anua", "Pacific/Guadalcanal"}
	case "BOA":
		return Location{"Boma", "Africa/Kinshasa"}
	case "BOB":
		return Location{"Motu Mute", "Pacific/Tahiti"}
	case "BOC":
		return Location{"Isla Colon", "America/Panama"}
	case "BOD":
		return Location{"Bordeaux/Merignac", "Europe/Paris"}
	case "BOE":
		return Location{"Boundji", "Africa/Brazzaville"}
	case "BOG":
		return Location{"Bogota", "America/Bogota"}
	case "BOH":
		return Location{"Bournemouth", "Europe/London"}
	case "BOI":
		return Location{"Boise", "America/Boise"}
	case "BOJ":
		return Location{"Burgas", "Europe/Sofia"}
	case "BOL":
		return Location{"Ballykelly", "Europe/London"}
	case "BOM":
		return Location{"Mumbai", "Asia/Kolkata"}
	case "BON":
		return Location{"Kralendijk", "America/Kralendijk"}
	case "BOO":
		return Location{"Bodo", "Europe/Oslo"}
	case "BOP":
		return Location{"Bouar", "Africa/Bangui"}
	case "BOR":
		return Location{"Belfort", "Europe/Paris"}
	case "BOS":
		return Location{"Boston", "America/New_York"}
	case "BOU":
		return Location{"Bourges", "Europe/Paris"}
	case "BOW":
		return Location{"Bartow", "America/New_York"}
	case "BOX":
		return Location{"", "Australia/Darwin"}
	case "BOY":
		return Location{"Bobo Dioulasso", "Africa/Ouagadougou"}
	case "BOZ":
		return Location{"Bozoum", "Africa/Bangui"}
	case "BPC":
		return Location{"Bamenda", "Africa/Douala"}
	case "BPF":
		return Location{"Batuna Mission Station", "Pacific/Guadalcanal"}
	case "BPG":
		return Location{"Barra Do Garcas", "America/Cuiaba"}
	case "BPH":
		return Location{"", "Asia/Manila"}
	case "BPI":
		return Location{"Big Piney", "America/Denver"}
	case "BPL":
		return Location{"Bole", "Asia/Shanghai"}
	case "BPM":
		return Location{"Hyderabad", "Asia/Kolkata"}
	case "BPN":
		return Location{"Balikpapan-Borneo Island", "Asia/Makassar"}
	case "BPR":
		return Location{"Borongan City", "Asia/Manila"}
	case "BPS":
		return Location{"Porto Seguro", "America/Bahia"}
	case "BPT":
		return Location{"Beaumont/Port Arthur", "America/Chicago"}
	case "BPX":
		return Location{"Bangda", "Asia/Shanghai"}
	case "BPY":
		return Location{"", "Indian/Antananarivo"}
	case "BQA":
		return Location{"Baler", "Asia/Manila"}
	case "BQB":
		return Location{"Busselton", "Australia/Perth"}
	case "BQD":
		return Location{"Budardalur", "Atlantic/Reykjavik"}
	case "BQE":
		return Location{"Bubaque", "Africa/Bissau"}
	case "BQG":
		return Location{"Bogorodskoye", "Asia/Vladivostok"}
	case "BQH":
		return Location{"London", "Europe/London"}
	case "BQK":
		return Location{"Brunswick", "America/New_York"}
	case "BQL":
		return Location{"", "Australia/Brisbane"}
	case "BQN":
		return Location{"Aguadilla", "America/Puerto_Rico"}
	case "BQO":
		return Location{"Bouna", "Africa/Abidjan"}
	case "BQQ":
		return Location{"Barra", "America/Bahia"}
	case "BQS":
		return Location{"Blagoveschensk", "Asia/Yakutsk"}
	case "BQT":
		return Location{"Brest", "Europe/Minsk"}
	case "BQU":
		return Location{"Bequia", "America/St_Vincent"}
	case "BQW":
		return Location{"", "Australia/Perth"}
	case "BRA":
		return Location{"Barreiras", "America/Bahia"}
	case "BRB":
		return Location{"", "America/Fortaleza"}
	case "BRC":
		return Location{"San Carlos de Bariloche", "America/Argentina/Salta"}
	case "BRD":
		return Location{"Brainerd", "America/Chicago"}
	case "BRE":
		return Location{"Bremen", "Europe/Berlin"}
	case "BRI":
		return Location{"Bari", "Europe/Rome"}
	case "BRK":
		return Location{"", "Australia/Sydney"}
	case "BRL":
		return Location{"Burlington", "America/Chicago"}
	case "BRM":
		return Location{"Barquisimeto", "America/Caracas"}
	case "BRN":
		return Location{"Bern", "Europe/Zurich"}
	case "BRO":
		return Location{"Brownsville", "America/Chicago"}
	case "BRQ":
		return Location{"Brno", "Europe/Prague"}
	case "BRR":
		return Location{"Eoligarry", "Europe/London"}
	case "BRS":
		return Location{"Bristol", "Europe/London"}
	case "BRT":
		return Location{"", "Australia/Darwin"}
	case "BRU":
		return Location{"Brussels", "Europe/Brussels"}
	case "BRV":
		return Location{"Bremerhaven", "Europe/Berlin"}
	case "BRW":
		return Location{"Barrow", "America/Anchorage"}
	case "BRX":
		return Location{"Barahona", "America/Santo_Domingo"}
	case "BRY":
		return Location{"Bardstown", "America/New_York"}
	case "BSA":
		return Location{"Bosaso", "Africa/Mogadishu"}
	case "BSB":
		return Location{"Brasilia", "America/Sao_Paulo"}
	case "BSC":
		return Location{"Bahia Solano", "America/Bogota"}
	case "BSD":
		return Location{"", "Asia/Shanghai"}
	case "BSE":
		return Location{"Sematan", "Asia/Kuching"}
	case "BSF":
		return Location{"Camp Pohakuloa", "Pacific/Honolulu"}
	case "BSG":
		return Location{"", "Africa/Malabo"}
	case "BSJ":
		return Location{"", "Australia/Melbourne"}
	case "BSK":
		return Location{"Biskra", "Africa/Algiers"}
	case "BSL":
		return Location{"Bale/Mulhouse", "Europe/Paris"}
	case "BSM":
		return Location{"Amol", "Asia/Tehran"}
	case "BSN":
		return Location{"Bossangoa", "Africa/Bangui"}
	case "BSO":
		return Location{"Basco", "Asia/Manila"}
	case "BSQ":
		return Location{"Bisbee", "America/Hermosillo"}
	case "BSR":
		return Location{"Basrah", "Asia/Baghdad"}
	case "BSS":
		return Location{"Balsas", "America/Fortaleza"}
	case "BST":
		return Location{"Bost", "Asia/Kabul"}
	case "BSU":
		return Location{"Basankusu", "Africa/Kinshasa"}
	case "BSW":
		return Location{"Boswell Bay", "America/Anchorage"}
	case "BSX":
		return Location{"Pathein", "Asia/Yangon"}
	case "BSY":
		return Location{"", "Africa/Mogadishu"}
	case "BSZ":
		return Location{"Bishkek", "Asia/Bishkek"}
	case "BTA":
		return Location{"Bertoua", "Africa/Douala"}
	case "BTB":
		return Location{"Betou", "Africa/Brazzaville"}
	case "BTC":
		return Location{"Batticaloa", "Asia/Colombo"}
	case "BTD":
		return Location{"", "Australia/Darwin"}
	case "BTE":
		return Location{"Bonthe", "Africa/Freetown"}
	case "BTF":
		return Location{"Bountiful", "America/Denver"}
	case "BTG":
		return Location{"Batangafo", "Africa/Bangui"}
	case "BTH":
		return Location{"Batam Island", "Asia/Jakarta"}
	case "BTI":
		return Location{"Barter Island Lrrs", "America/Anchorage"}
	case "BTJ":
		return Location{"Banda Aceh-Sumatra Island", "Asia/Jakarta"}
	case "BTK":
		return Location{"Bratsk", "Asia/Irkutsk"}
	case "BTL":
		return Location{"Battle Creek", "America/Detroit"}
	case "BTM":
		return Location{"Butte", "America/Denver"}
	case "BTN":
		return Location{"Bennettsville", "America/New_York"}
	case "BTO":
		return Location{"Botopasi", "America/Paramaribo"}
	case "BTP":
		return Location{"Butler", "America/New_York"}
	case "BTQ":
		return Location{"Butare", "Africa/Kigali"}
	case "BTR":
		return Location{"Baton Rouge", "America/Chicago"}
	case "BTS":
		return Location{"Bratislava", "Europe/Bratislava"}
	case "BTT":
		return Location{"Bettles", "America/Anchorage"}
	case "BTU":
		return Location{"Bintulu", "Asia/Kuching"}
	case "BTV":
		return Location{"Burlington", "America/New_York"}
	case "BTW":
		return Location{"Batu Licin-Borneo Island", "Asia/Makassar"}
	case "BTX":
		return Location{"", "Australia/Brisbane"}
	case "BTY":
		return Location{"Beatty", "America/Los_Angeles"}
	case "BTZ":
		return Location{"Yarom", "Asia/Bangkok"}
	case "BUA":
		return Location{"Buka Island", "Pacific/Bougainville"}
	case "BUB":
		return Location{"Burwell", "America/Chicago"}
	case "BUC":
		return Location{"", "Australia/Brisbane"}
	case "BUD":
		return Location{"Budapest", "Europe/Budapest"}
	case "BUF":
		return Location{"Buffalo", "America/New_York"}
	case "BUG":
		return Location{"Benguela", "Africa/Luanda"}
	case "BUI":
		return Location{"Bokondini-Papua Island", "Asia/Jayapura"}
	case "BUJ":
		return Location{"", "Africa/Algiers"}
	case "BUL":
		return Location{"Bulolo", "Pacific/Port_Moresby"}
	case "BUM":
		return Location{"Butler", "America/Chicago"}
	case "BUN":
		return Location{"Buenaventura", "America/Bogota"}
	case "BUO":
		return Location{"Burao", "Africa/Mogadishu"}
	case "BUP":
		return Location{"", "Asia/Kolkata"}
	case "BUQ":
		return Location{"Bulawayo", "Africa/Harare"}
	case "BUR":
		return Location{"Burbank", "America/Los_Angeles"}
	case "BUS":
		return Location{"Batumi", "Asia/Tbilisi"}
	case "BUW":
		return Location{"Bau Bau-Butung Island", "Asia/Makassar"}
	case "BUX":
		return Location{"", "Africa/Lubumbashi"}
	case "BUY":
		return Location{"", "Australia/Perth"}
	case "BUZ":
		return Location{"Bushehr", "Asia/Tehran"}
	case "BVA":
		return Location{"Beauvais/Tille", "Europe/Paris"}
	case "BVB":
		return Location{"Boa Vista", "America/Boa_Vista"}
	case "BVC":
		return Location{"Rabil", "Atlantic/Cape_Verde"}
	case "BVE":
		return Location{"Nespouls", "Europe/Paris"}
	case "BVG":
		return Location{"Berlevag", "Europe/Oslo"}
	case "BVH":
		return Location{"Vilhena", "America/Porto_Velho"}
	case "BVI":
		return Location{"", "Australia/Brisbane"}
	case "BVJ":
		return Location{"Bovanenkovo", "Asia/Yekaterinburg"}
	case "BVK":
		return Location{"Itenes", "America/La_Paz"}
	case "BVL":
		return Location{"Baures", "America/La_Paz"}
	case "BVM":
		return Location{"Belmonte", "America/Bahia"}
	case "BVO":
		return Location{"Bartlesville", "America/Chicago"}
	case "BVR":
		return Location{"Brava Island", "Atlantic/Cape_Verde"}
	case "BVS":
		return Location{"Breves", "America/Belem"}
	case "BVU":
		return Location{"Beluga", "America/Anchorage"}
	case "BVV":
		return Location{"Iturup Island", "Asia/Magadan"}
	case "BVX":
		return Location{"Batesville", "America/Chicago"}
	case "BVY":
		return Location{"Beverly", "America/New_York"}
	case "BVZ":
		return Location{"", "Australia/Perth"}
	case "BWA":
		return Location{"Bhairawa", "Asia/Kathmandu"}
	case "BWB":
		return Location{"", "Australia/Perth"}
	case "BWC":
		return Location{"Brawley", "America/Los_Angeles"}
	case "BWD":
		return Location{"Brownwood", "America/Chicago"}
	case "BWE":
		return Location{"", "Europe/Berlin"}
	case "BWF":
		return Location{"Barrow-in-Furness", "Europe/London"}
	case "BWG":
		return Location{"Bowling Green", "America/Chicago"}
	case "BWH":
		return Location{"Butterworth", "Asia/Kuala_Lumpur"}
	case "BWI":
		return Location{"Baltimore", "America/New_York"}
	case "BWK":
		return Location{"Brac Island", "Europe/Zagreb"}
	case "BWL":
		return Location{"Blackwell", "America/Chicago"}
	case "BWM":
		return Location{"Bowman", "America/Denver"}
	case "BWN":
		return Location{"Bandar Seri Begawan", "Asia/Brunei"}
	case "BWO":
		return Location{"Balakovo", "Europe/Saratov"}
	case "BWQ":
		return Location{"", "Australia/Sydney"}
	case "BWT":
		return Location{"Burnie", "Australia/Hobart"}
	case "BWU":
		return Location{"Sydney", "Australia/Sydney"}
	case "BWW":
		return Location{"Cayo Santa Maria", "America/Havana"}
	case "BXA":
		return Location{"Bogalusa", "America/Chicago"}
	case "BXB":
		return Location{"Babo-Papua Island", "Asia/Jayapura"}
	case "BXD":
		return Location{"Bade-Papua Island", "Asia/Jayapura"}
	case "BXE":
		return Location{"Bakel", "Africa/Dakar"}
	case "BXF":
		return Location{"Bellburn", "Australia/Perth"}
	case "BXG":
		return Location{"", "Australia/Melbourne"}
	case "BXH":
		return Location{"Balkhash", "Asia/Almaty"}
	case "BXI":
		return Location{"Boundiali", "Africa/Abidjan"}
	case "BXJ":
		return Location{"Aima Ata", "Asia/Almaty"}
	case "BXK":
		return Location{"Buckeye", "America/Phoenix"}
	case "BXN":
		return Location{"Bodrum", "Europe/Istanbul"}
	case "BXO":
		return Location{"Buochs", "Europe/Zurich"}
	case "BXP":
		return Location{"Biala Podlaska", "Europe/Warsaw"}
	case "BXR":
		return Location{"", "Asia/Tehran"}
	case "BXS":
		return Location{"Borrego Springs", "America/Los_Angeles"}
	case "BXT":
		return Location{"Bontang-Borneo Island", "Asia/Makassar"}
	case "BXU":
		return Location{"Butuan City", "Asia/Manila"}
	case "BXV":
		return Location{"Breiddalsvik", "Atlantic/Reykjavik"}
	case "BXY":
		return Location{"Baikonur", "Asia/Qyzylorda"}
	case "BYC":
		return Location{"Yacuiba", "America/La_Paz"}
	case "BYD":
		return Location{"Al-Bayda", "Asia/Aden"}
	case "BYF":
		return Location{"Albert/Bray", "Europe/Paris"}
	case "BYG":
		return Location{"Buffalo", "America/Denver"}
	case "BYH":
		return Location{"Blytheville", "America/Chicago"}
	case "BYI":
		return Location{"Burley", "America/Boise"}
	case "BYJ":
		return Location{"Beja", "Europe/Lisbon"}
	case "BYK":
		return Location{"", "Africa/Abidjan"}
	case "BYM":
		return Location{"Bayamo", "America/Havana"}
	case "BYN":
		return Location{"Bayankhongor", "Asia/Ulaanbaatar"}
	case "BYO":
		return Location{"Bonito", "America/Campo_Grande"}
	case "BYP":
		return Location{"", "Australia/Perth"}
	case "BYQ":
		return Location{"Bunju Island", "Asia/Makassar"}
	case "BYR":
		return Location{"Laeso", "Europe/Copenhagen"}
	case "BYS":
		return Location{"Fort Irwin/Barstow", "America/Los_Angeles"}
	case "BYT":
		return Location{"Bantry", "Europe/Dublin"}
	case "BYU":
		return Location{"Bayreuth", "Europe/Berlin"}
	case "BYW":
		return Location{"Blakely Island", "America/Los_Angeles"}
	case "BZA":
		return Location{"Bonanza", "America/Managua"}
	case "BZC":
		return Location{"Cabo Frio", "America/Sao_Paulo"}
	case "BZD":
		return Location{"", "Australia/Sydney"}
	case "BZE":
		return Location{"Belize City", "America/Belize"}
	case "BZG":
		return Location{"Bydgoszcz", "Europe/Warsaw"}
	case "BZI":
		return Location{"Balikesir", "Europe/Istanbul"}
	case "BZK":
		return Location{"Bryansk", "Europe/Moscow"}
	case "BZL":
		return Location{"Barisal", "Asia/Dhaka"}
	case "BZN":
		return Location{"Bozeman", "America/Denver"}
	case "BZO":
		return Location{"Bolzano", "Europe/Rome"}
	case "BZP":
		return Location{"Lakefield National Park", "Australia/Brisbane"}
	case "BZR":
		return Location{"Beziers/Vias", "Europe/Paris"}
	case "BZU":
		return Location{"", "Africa/Lubumbashi"}
	case "BZV":
		return Location{"Brazzaville", "Africa/Brazzaville"}
	case "BZX":
		return Location{"Bazhong", "Asia/Shanghai"}
	case "BZY":
		return Location{"Strymba", "Europe/Chisinau"}
	case "BZZ":
		return Location{"Brize Norton", "Europe/London"}
	case "CAA":
		return Location{"Catacamas", "America/Tegucigalpa"}
	case "CAB":
		return Location{"Cabinda", "Africa/Luanda"}
	case "CAC":
		return Location{"Cascavel", "America/Sao_Paulo"}
	case "CAD":
		return Location{"Cadillac", "America/Detroit"}
	case "CAE":
		return Location{"Columbia", "America/New_York"}
	case "CAF":
		return Location{"Carauari", "America/Manaus"}
	case "CAG":
		return Location{"Cagliari", "Europe/Rome"}
	case "CAH":
		return Location{"Ca Mau City", "Asia/Ho_Chi_Minh"}
	case "CAI":
		return Location{"Cairo", "Africa/Cairo"}
	case "CAJ":
		return Location{"Canaima", "America/Caracas"}
	case "CAK":
		return Location{"Akron", "America/New_York"}
	case "CAL":
		return Location{"Campbeltown", "Europe/London"}
	case "CAM":
		return Location{"Camiri", "America/La_Paz"}
	case "CAN":
		return Location{"Guangzhou", "Asia/Shanghai"}
	case "CAO":
		return Location{"Clayton", "America/Denver"}
	case "CAP":
		return Location{"Cap Haitien", "America/Port-au-Prince"}
	case "CAQ":
		return Location{"Caucasia", "America/Bogota"}
	case "CAR":
		return Location{"Caribou", "America/New_York"}
	case "CAS":
		return Location{"Casablanca", "Africa/Casablanca"}
	case "CAT":
		return Location{"Cascais", "Europe/Lisbon"}
	case "CAU":
		return Location{"Caruaru", "America/Recife"}
	case "CAV":
		return Location{"Cazombo", "Africa/Luanda"}
	case "CAW":
		return Location{"Campos Dos Goytacazes", "America/Sao_Paulo"}
	case "CAX":
		return Location{"Carlisle", "Europe/London"}
	case "CAY":
		return Location{"Cayenne / Rochambeau", "America/Cayenne"}
	case "CAZ":
		return Location{"", "Australia/Sydney"}
	case "CBB":
		return Location{"Cochabamba", "America/La_Paz"}
	case "CBD":
		return Location{"", "Asia/Kolkata"}
	case "CBE":
		return Location{"Cumberland", "America/New_York"}
	case "CBF":
		return Location{"Council Bluffs", "America/Chicago"}
	case "CBG":
		return Location{"Cambridge", "Europe/London"}
	case "CBH":
		return Location{"Bechar", "Africa/Algiers"}
	case "CBI":
		return Location{"", "Australia/Hobart"}
	case "CBJ":
		return Location{"Cabo Rojo", "America/Santo_Domingo"}
	case "CBK":
		return Location{"Colby", "America/Chicago"}
	case "CBL":
		return Location{"", "America/Caracas"}
	case "CBM":
		return Location{"Columbus", "America/Chicago"}
	case "CBN":
		return Location{"Cirebon-Java Island", "Asia/Jakarta"}
	case "CBO":
		return Location{"Cotabato City", "Asia/Manila"}
	case "CBP":
		return Location{"", "Europe/Lisbon"}
	case "CBQ":
		return Location{"Calabar", "Africa/Lagos"}
	case "CBR":
		return Location{"Canberra", "Australia/Sydney"}
	case "CBS":
		return Location{"Cabimas", "America/Caracas"}
	case "CBT":
		return Location{"Catumbela", "Africa/Luanda"}
	case "CBU":
		return Location{"Cottbus", "Europe/Berlin"}
	case "CBV":
		return Location{"Coban", "America/Guatemala"}
	case "CBW":
		return Location{"Campo Mourao", "America/Sao_Paulo"}
	case "CBX":
		return Location{"", "Australia/Sydney"}
	case "CBY":
		return Location{"Canobie", "Australia/Brisbane"}
	case "CCB":
		return Location{"Upland", "America/Los_Angeles"}
	case "CCC":
		return Location{"Cayo Coco", "America/Havana"}
	case "CCE":
		return Location{"New Administrative Capital", "Africa/Cairo"}
	case "CCF":
		return Location{"Carcassonne/Salvaza", "Europe/Paris"}
	case "CCG":
		return Location{"Crane", "America/Chicago"}
	case "CCH":
		return Location{"Chile Chico", "America/Santiago"}
	case "CCI":
		return Location{"Concordia", "America/Sao_Paulo"}
	case "CCJ":
		return Location{"Calicut", "Asia/Kolkata"}
	case "CCK":
		return Location{"Cocos (Keeling) Islands", "Indian/Cocos"}
	case "CCL":
		return Location{"", "Australia/Brisbane"}
	case "CCM":
		return Location{"Criciuma", "America/Sao_Paulo"}
	case "CCN":
		return Location{"Chakcharan", "Asia/Kabul"}
	case "CCO":
		return Location{"Puerto Lopez", "America/Bogota"}
	case "CCP":
		return Location{"Concepcion", "America/Santiago"}
	case "CCR":
		return Location{"Concord", "America/Los_Angeles"}
	case "CCS":
		return Location{"Caracas", "America/Caracas"}
	case "CCT":
		return Location{"Colonia Catriel", "America/Argentina/Salta"}
	case "CCU":
		return Location{"Kolkata", "Asia/Kolkata"}
	case "CCV":
		return Location{"Craig Cove", "Pacific/Efate"}
	case "CCW":
		return Location{"", "Australia/Adelaide"}
	case "CCX":
		return Location{"Caceres", "America/Cuiaba"}
	case "CCY":
		return Location{"Charles City", "America/Chicago"}
	case "CCZ":
		return Location{"", "America/Nassau"}
	case "CDA":
		return Location{"", "Australia/Darwin"}
	case "CDB":
		return Location{"Cold Bay", "America/Nome"}
	case "CDC":
		return Location{"Cedar City", "America/Denver"}
	case "CDD":
		return Location{"Cauquira", "America/Tegucigalpa"}
	case "CDE":
		return Location{"Chengde", "Asia/Shanghai"}
	case "CDF":
		return Location{"Cortina D'Ampezzo", "Europe/Rome"}
	case "CDG":
		return Location{"Paris", "Europe/Paris"}
	case "CDH":
		return Location{"Camden", "America/Chicago"}
	case "CDI":
		return Location{"Cachoeiro Do Itapemirim", "America/Sao_Paulo"}
	case "CDJ":
		return Location{"Conceicao Do Araguaia", "America/Araguaina"}
	case "CDK":
		return Location{"Cedar Key", "America/New_York"}
	case "CDL":
		return Location{"Candle", "America/Anchorage"}
	case "CDN":
		return Location{"Camden", "America/New_York"}
	case "CDO":
		return Location{"Cradock", "Africa/Johannesburg"}
	case "CDP":
		return Location{"", "Asia/Kolkata"}
	case "CDQ":
		return Location{"", "Australia/Brisbane"}
	case "CDR":
		return Location{"Chadron", "America/Denver"}
	case "CDS":
		return Location{"Childress", "America/Chicago"}
	case "CDT":
		return Location{"Castellón", "Europe/Madrid"}
	case "CDU":
		return Location{"", "Australia/Sydney"}
	case "CDV":
		return Location{"Cordova", "America/Anchorage"}
	case "CDW":
		return Location{"Caldwell", "America/New_York"}
	case "CDY":
		return Location{"Mapun", "Asia/Manila"}
	case "CEA":
		return Location{"Wichita", "America/Chicago"}
	case "CEB":
		return Location{"Lapu-Lapu City", "Asia/Manila"}
	case "CEC":
		return Location{"Crescent City", "America/Los_Angeles"}
	case "CED":
		return Location{"", "Australia/Adelaide"}
	case "CEE":
		return Location{"Cherepovets", "Europe/Moscow"}
	case "CEF":
		return Location{"Springfield/Chicopee", "America/New_York"}
	case "CEG":
		return Location{"Hawarden", "Europe/London"}
	case "CEH":
		return Location{"", "Africa/Blantyre"}
	case "CEI":
		return Location{"Chiang Rai", "Asia/Bangkok"}
	case "CEK":
		return Location{"Chelyabinsk", "Asia/Yekaterinburg"}
	case "CEL":
		return Location{"Canela", "America/Sao_Paulo"}
	case "CEM":
		return Location{"Central", "America/Anchorage"}
	case "CEN":
		return Location{"Ciudad Obregon", "America/Hermosillo"}
	case "CEO":
		return Location{"Waco Kungo", "Africa/Luanda"}
	case "CEP":
		return Location{"Concepcion", "America/La_Paz"}
	case "CEQ":
		return Location{"Cannes/Mandelieu", "Europe/Paris"}
	case "CER":
		return Location{"Cherbourg/Maupertus", "Europe/Paris"}
	case "CES":
		return Location{"", "Australia/Sydney"}
	case "CET":
		return Location{"Cholet/Le Pontreau", "Europe/Paris"}
	case "CEU":
		return Location{"Clemson", "America/New_York"}
	case "CEV":
		return Location{"Connersville", "America/Indiana/Indianapolis"}
	case "CEW":
		return Location{"Crestview", "America/Chicago"}
	case "CEX":
		return Location{"Chena Hot Springs", "America/Anchorage"}
	case "CEY":
		return Location{"Murray", "America/Chicago"}
	case "CEZ":
		return Location{"Cortez", "America/Denver"}
	case "CFB":
		return Location{"Cabo Frio", "America/Sao_Paulo"}
	case "CFC":
		return Location{"Caçador", "America/Sao_Paulo"}
	case "CFD":
		return Location{"Bryan", "America/Chicago"}
	case "CFE":
		return Location{"Clermont-Ferrand/Auvergne", "Europe/Paris"}
	case "CFF":
		return Location{"Cafunfo", "Africa/Luanda"}
	case "CFG":
		return Location{"Cienfuegos", "America/Havana"}
	case "CFH":
		return Location{"Clifton Hills Station", "Australia/Adelaide"}
	case "CFI":
		return Location{"", "Australia/Darwin"}
	case "CFK":
		return Location{"Chlef", "Africa/Algiers"}
	case "CFN":
		return Location{"Donegal", "Europe/Dublin"}
	case "CFO":
		return Location{"Confresa", "America/Cuiaba"}
	case "CFR":
		return Location{"Caen/Carpiquet", "Europe/Paris"}
	case "CFS":
		return Location{"Coffs Harbour", "Australia/Sydney"}
	case "CFT":
		return Location{"Clifton/Morenci", "America/Phoenix"}
	case "CFU":
		return Location{"Kerkyra Island", "Europe/Athens"}
	case "CFV":
		return Location{"Coffeyville", "America/Chicago"}
	case "CGB":
		return Location{"Cuiaba", "America/Cuiaba"}
	case "CGC":
		return Location{"Cape Gloucester", "Pacific/Port_Moresby"}
	case "CGD":
		return Location{"Changde", "Asia/Shanghai"}
	case "CGE":
		return Location{"Cambridge", "America/New_York"}
	case "CGF":
		return Location{"Cleveland", "America/New_York"}
	case "CGH":
		return Location{"Sao Paulo", "America/Sao_Paulo"}
	case "CGI":
		return Location{"Cape Girardeau", "America/Chicago"}
	case "CGJ":
		return Location{"Kasompe", "Africa/Lusaka"}
	case "CGK":
		return Location{"Jakarta", "Asia/Jakarta"}
	case "CGM":
		return Location{"", "Asia/Manila"}
	case "CGN":
		return Location{"Cologne", "Europe/Berlin"}
	case "CGO":
		return Location{"Zhengzhou", "Asia/Shanghai"}
	case "CGP":
		return Location{"Chittagong", "Asia/Dhaka"}
	case "CGQ":
		return Location{"Changchun", "Asia/Shanghai"}
	case "CGR":
		return Location{"Campo Grande", "America/Campo_Grande"}
	case "CGS":
		return Location{"College Park", "America/New_York"}
	case "CGV":
		return Location{"", "Australia/Eucla"}
	case "CGY":
		return Location{"Laguindingan", "Asia/Manila"}
	case "CGZ":
		return Location{"Casa Grande", "America/Phoenix"}
	case "CHA":
		return Location{"Chattanooga", "America/New_York"}
	case "CHB":
		return Location{"Chilas", "Asia/Karachi"}
	case "CHC":
		return Location{"Christchurch", "Pacific/Auckland"}
	case "CHF":
		return Location{"Jinhae", "Asia/Seoul"}
	case "CHG":
		return Location{"Chaoyang", "Asia/Shanghai"}
	case "CHH":
		return Location{"Chachapoyas", "America/Lima"}
	case "CHJ":
		return Location{"Chipinge", "Africa/Harare"}
	case "CHK":
		return Location{"Chickasha", "America/Chicago"}
	case "CHL":
		return Location{"Challis", "America/Boise"}
	case "CHM":
		return Location{"Chimbote", "America/Lima"}
	case "CHN":
		return Location{"Jeon Ju", "Asia/Seoul"}
	case "CHO":
		return Location{"Charlottesville", "America/New_York"}
	case "CHQ":
		return Location{"Souda", "Europe/Athens"}
	case "CHR":
		return Location{"Chateauroux/Deols", "Europe/Paris"}
	case "CHS":
		return Location{"Charleston", "America/New_York"}
	case "CHT":
		return Location{"Waitangi", "Pacific/Chatham"}
	case "CHU":
		return Location{"Chuathbaluk", "America/Anchorage"}
	case "CHV":
		return Location{"Chaves", "Europe/Lisbon"}
	case "CHX":
		return Location{"Changuinola", "America/Panama"}
	case "CHY":
		return Location{"", "Pacific/Guadalcanal"}
	case "CIA":
		return Location{"Roma", "Europe/Rome"}
	case "CIC":
		return Location{"Chico", "America/Los_Angeles"}
	case "CID":
		return Location{"Cedar Rapids", "America/Chicago"}
	case "CIE":
		return Location{"", "Australia/Perth"}
	case "CIF":
		return Location{"Chifeng", "Asia/Shanghai"}
	case "CIG":
		return Location{"Craig", "America/Denver"}
	case "CIH":
		return Location{"Changzhi", "Asia/Shanghai"}
	case "CII":
		return Location{"Aydin", "Europe/Istanbul"}
	case "CIJ":
		return Location{"Cobija", "America/La_Paz"}
	case "CIK":
		return Location{"Chalkyitsik", "America/Anchorage"}
	case "CIM":
		return Location{"Cimitarra", "America/Bogota"}
	case "CIN":
		return Location{"Carroll", "America/Chicago"}
	case "CIO":
		return Location{"Concepcion", "America/Asuncion"}
	case "CIP":
		return Location{"Chipata", "Africa/Lusaka"}
	case "CIQ":
		return Location{"Chiquimula", "America/Guatemala"}
	case "CIR":
		return Location{"Cairo", "America/Chicago"}
	case "CIS":
		return Location{"Abariringa", "Pacific/Kanton"}
	case "CIT":
		return Location{"Shymkent", "Asia/Almaty"}
	case "CIU":
		return Location{"Sault Ste Marie", "America/Detroit"}
	case "CIW":
		return Location{"Canouan", "America/St_Vincent"}
	case "CIX":
		return Location{"Chiclayo", "America/Lima"}
	case "CIY":
		return Location{"Comiso", "Europe/Rome"}
	case "CIZ":
		return Location{"Coari", "America/Manaus"}
	case "CJA":
		return Location{"Cajamarca", "America/Lima"}
	case "CJB":
		return Location{"Coimbatore", "Asia/Kolkata"}
	case "CJC":
		return Location{"Calama", "America/Santiago"}
	case "CJF":
		return Location{"", "Australia/Perth"}
	case "CJJ":
		return Location{"Cheongju", "Asia/Seoul"}
	case "CJL":
		return Location{"Chitral", "Asia/Karachi"}
	case "CJM":
		return Location{"", "Asia/Bangkok"}
	case "CJS":
		return Location{"Ciudad Juarez", "America/Ciudad_Juarez"}
	case "CJT":
		return Location{"", "America/Mexico_City"}
	case "CJU":
		return Location{"Jeju City", "Asia/Seoul"}
	case "CKA":
		return Location{"Cherokee", "America/Chicago"}
	case "CKB":
		return Location{"Clarksburg", "America/New_York"}
	case "CKC":
		return Location{"Cherkasy", "Europe/Kyiv"}
	case "CKG":
		return Location{"Chongqing", "Asia/Shanghai"}
	case "CKH":
		return Location{"Chokurdah", "Asia/Srednekolymsk"}
	case "CKI":
		return Location{"", "Australia/Darwin"}
	case "CKK":
		return Location{"Ash Flat", "America/Chicago"}
	case "CKL":
		return Location{"Moscow", "Europe/Moscow"}
	case "CKM":
		return Location{"Clarksdale", "America/Chicago"}
	case "CKN":
		return Location{"Crookston", "America/Chicago"}
	case "CKO":
		return Location{"Cornelio Procopio", "America/Sao_Paulo"}
	case "CKS":
		return Location{"Carajas", "America/Belem"}
	case "CKT":
		return Location{"Sarakhs", "Asia/Tehran"}
	case "CKV":
		return Location{"Clarksville", "America/Chicago"}
	case "CKW":
		return Location{"Christmas Creek mine", "Australia/Perth"}
	case "CKY":
		return Location{"Conakry", "Africa/Conakry"}
	case "CKZ":
		return Location{"Canakkale", "Europe/Istanbul"}
	case "CLA":
		return Location{"Comilla", "Asia/Dhaka"}
	case "CLB":
		return Location{"Castlebar", "Europe/Dublin"}
	case "CLD":
		return Location{"Carlsbad", "America/Los_Angeles"}
	case "CLE":
		return Location{"Cleveland", "America/New_York"}
	case "CLH":
		return Location{"", "Australia/Sydney"}
	case "CLI":
		return Location{"Clintonville", "America/Chicago"}
	case "CLJ":
		return Location{"Cluj-Napoca", "Europe/Bucharest"}
	case "CLK":
		return Location{"Clinton", "America/Chicago"}
	case "CLL":
		return Location{"College Station", "America/Chicago"}
	case "CLM":
		return Location{"Port Angeles", "America/Los_Angeles"}
	case "CLN":
		return Location{"Carolina", "America/Fortaleza"}
	case "CLO":
		return Location{"Cali", "America/Bogota"}
	case "CLP":
		return Location{"Clarks Point", "America/Anchorage"}
	case "CLQ":
		return Location{"Colima", "America/Mexico_City"}
	case "CLR":
		return Location{"Calipatria", "America/Los_Angeles"}
	case "CLS":
		return Location{"Chehalis", "America/Los_Angeles"}
	case "CLT":
		return Location{"Charlotte", "America/New_York"}
	case "CLU":
		return Location{"Columbus", "America/Indiana/Indianapolis"}
	case "CLV":
		return Location{"Caldas Novas", "America/Sao_Paulo"}
	case "CLW":
		return Location{"Clearwater", "America/New_York"}
	case "CLX":
		return Location{"Clorinda", "America/Argentina/Cordoba"}
	case "CLY":
		return Location{"Calvi/Sainte-Catherine", "Europe/Paris"}
	case "CLZ":
		return Location{"Guarico", "America/Caracas"}
	case "CMA":
		return Location{"", "Australia/Brisbane"}
	case "CMB":
		return Location{"Colombo", "Asia/Colombo"}
	case "CMC":
		return Location{"Camocim", "America/Fortaleza"}
	case "CMD":
		return Location{"", "Australia/Sydney"}
	case "CME":
		return Location{"Ciudad del Carmen", "America/Merida"}
	case "CMF":
		return Location{"Chambery/Aix-les-Bains", "Europe/Paris"}
	case "CMG":
		return Location{"Corumba", "America/Campo_Grande"}
	case "CMH":
		return Location{"Columbus", "America/New_York"}
	case "CMI":
		return Location{"Champaign/Urbana", "America/Chicago"}
	case "CMJ":
		return Location{"Qimei", "Asia/Taipei"}
	case "CMK":
		return Location{"Club Makokola", "Africa/Blantyre"}
	case "CML":
		return Location{"", "Australia/Brisbane"}
	case "CMM":
		return Location{"Carmelita", "America/Guatemala"}
	case "CMN":
		return Location{"Casablanca", "Africa/Casablanca"}
	case "CMO":
		return Location{"Obbia", "Africa/Mogadishu"}
	case "CMP":
		return Location{"Santana Do Araguaia", "America/Belem"}
	case "CMQ":
		return Location{"", "Australia/Brisbane"}
	case "CMR":
		return Location{"Colmar/Houssen", "Europe/Paris"}
	case "CMS":
		return Location{"Scusciuban", "Africa/Mogadishu"}
	case "CMU":
		return Location{"Kundiawa", "Pacific/Port_Moresby"}
	case "CMV":
		return Location{"", "Pacific/Auckland"}
	case "CMW":
		return Location{"Camaguey", "America/Havana"}
	case "CMX":
		return Location{"Hancock", "America/Detroit"}
	case "CMY":
		return Location{"Sparta", "America/Chicago"}
	case "CNA":
		return Location{"", "America/Hermosillo"}
	case "CNB":
		return Location{"", "Australia/Sydney"}
	case "CNC":
		return Location{"", "Australia/Brisbane"}
	case "CND":
		return Location{"Constanta", "Europe/Bucharest"}
	case "CNF":
		return Location{"Belo Horizonte", "America/Sao_Paulo"}
	case "CNG":
		return Location{"Cognac/Chateaubernard", "Europe/Paris"}
	case "CNH":
		return Location{"Claremont", "America/New_York"}
	case "CNI":
		return Location{"Changhai", "Asia/Shanghai"}
	case "CNJ":
		return Location{"Cloncurry", "Australia/Brisbane"}
	case "CNK":
		return Location{"Concordia", "America/Chicago"}
	case "CNL":
		return Location{"Sindal", "Europe/Copenhagen"}
	case "CNM":
		return Location{"Carlsbad", "America/Denver"}
	case "CNN":
		return Location{"Mattanur", "Asia/Kolkata"}
	case "CNO":
		return Location{"Chino", "America/Los_Angeles"}
	case "CNP":
		return Location{"Neerlerit Inaat", "America/Scoresbysund"}
	case "CNQ":
		return Location{"Corrientes", "America/Argentina/Cordoba"}
	case "CNR":
		return Location{"Chanaral", "America/Santiago"}
	case "CNS":
		return Location{"Cairns", "Australia/Brisbane"}
	case "CNU":
		return Location{"Chanute", "America/Chicago"}
	case "CNV":
		return Location{"Canavieiras", "America/Bahia"}
	case "CNW":
		return Location{"Waco", "America/Chicago"}
	case "CNX":
		return Location{"Chiang Mai", "Asia/Bangkok"}
	case "CNY":
		return Location{"Moab", "America/Denver"}
	case "COC":
		return Location{"Concordia", "America/Argentina/Cordoba"}
	case "COD":
		return Location{"Cody", "America/Denver"}
	case "COE":
		return Location{"Coeur d'Alene", "America/Los_Angeles"}
	case "COF":
		return Location{"Cocoa Beach", "America/New_York"}
	case "COG":
		return Location{"Condoto", "America/Bogota"}
	case "COH":
		return Location{"", "Asia/Kolkata"}
	case "COI":
		return Location{"Merritt Island", "America/New_York"}
	case "COJ":
		return Location{"", "Australia/Sydney"}
	case "COK":
		return Location{"Cochin", "Asia/Kolkata"}
	case "COL":
		return Location{"Coll Island", "Europe/London"}
	case "COM":
		return Location{"Coleman", "America/Chicago"}
	case "CON":
		return Location{"Concord", "America/New_York"}
	case "COO":
		return Location{"Cotonou", "Africa/Porto-Novo"}
	case "COQ":
		return Location{"", "Asia/Choibalsan"}
	case "COR":
		return Location{"Cordoba", "America/Argentina/Cordoba"}
	case "COS":
		return Location{"Colorado Springs", "America/Denver"}
	case "COT":
		return Location{"Cotulla", "America/Chicago"}
	case "COU":
		return Location{"Columbia", "America/Chicago"}
	case "COV":
		return Location{"Tarsus", "Europe/Istanbul"}
	case "COW":
		return Location{"Coquimbo", "America/Santiago"}
	case "COX":
		return Location{"Andros", "America/Nassau"}
	case "COY":
		return Location{"", "Australia/Perth"}
	case "COZ":
		return Location{"Costanza", "America/Santo_Domingo"}
	case "CPA":
		return Location{"Harper", "Africa/Monrovia"}
	case "CPB":
		return Location{"Capurgana", "America/Bogota"}
	case "CPC":
		return Location{"Chapelco/San Martin de los Andes", "America/Argentina/Salta"}
	case "CPD":
		return Location{"", "Australia/Adelaide"}
	case "CPE":
		return Location{"Campeche", "America/Merida"}
	case "CPF":
		return Location{"Tjepu-Java Island", "Asia/Jakarta"}
	case "CPH":
		return Location{"Copenhagen", "Europe/Copenhagen"}
	case "CPL":
		return Location{"Chaparral", "America/Bogota"}
	case "CPM":
		return Location{"Compton", "America/Los_Angeles"}
	case "CPO":
		return Location{"Copiapo", "America/Santiago"}
	case "CPP":
		return Location{"Pica", "America/Santiago"}
	case "CPQ":
		return Location{"Campinas", "America/Sao_Paulo"}
	case "CPR":
		return Location{"Casper", "America/Denver"}
	case "CPS":
		return Location{"Cahokia/St Louis", "America/Chicago"}
	case "CPT":
		return Location{"Cape Town", "Africa/Johannesburg"}
	case "CPU":
		return Location{"Cururupu", "America/Fortaleza"}
	case "CPV":
		return Location{"Campina Grande", "America/Fortaleza"}
	case "CPX":
		return Location{"Culebra Island", "America/Puerto_Rico"}
	case "CQA":
		return Location{"Canarana", "America/Cuiaba"}
	case "CQD":
		return Location{"Shahrekord", "Asia/Tehran"}
	case "CQF":
		return Location{"Marck", "Europe/Paris"}
	case "CQM":
		return Location{"Ciudad Real", "Europe/Madrid"}
	case "CQS":
		return Location{"Costa Marques", "America/Porto_Velho"}
	case "CQW":
		return Location{"Wulong", "Asia/Shanghai"}
	case "CRA":
		return Location{"Craiova", "Europe/Bucharest"}
	case "CRB":
		return Location{"", "Australia/Sydney"}
	case "CRC":
		return Location{"Cartago", "America/Bogota"}
	case "CRD":
		return Location{"Comodoro Rivadavia", "America/Argentina/Catamarca"}
	case "CRE":
		return Location{"North Myrtle Beach", "America/New_York"}
	case "CRF":
		return Location{"Carnot", "Africa/Bangui"}
	case "CRG":
		return Location{"Jacksonville", "America/New_York"}
	case "CRI":
		return Location{"Colonel Hill", "America/Nassau"}
	case "CRK":
		return Location{"Angeles City", "Asia/Manila"}
	case "CRL":
		return Location{"Brussels", "Europe/Brussels"}
	case "CRM":
		return Location{"Catarman", "Asia/Manila"}
	case "CRO":
		return Location{"Corcoran", "America/Los_Angeles"}
	case "CRP":
		return Location{"Corpus Christi", "America/Chicago"}
	case "CRQ":
		return Location{"Caravelas", "America/Bahia"}
	case "CRR":
		return Location{"Ceres", "America/Argentina/Cordoba"}
	case "CRS":
		return Location{"Corsicana", "America/Chicago"}
	case "CRT":
		return Location{"Crossett", "America/Chicago"}
	case "CRU":
		return Location{"Carriacou Island", "America/Grenada"}
	case "CRV":
		return Location{"Crotone", "Europe/Rome"}
	case "CRW":
		return Location{"Charleston", "America/New_York"}
	case "CRX":
		return Location{"Corinth", "America/Chicago"}
	case "CRZ":
		return Location{"Turkmenabat", "Asia/Ashgabat"}
	case "CSA":
		return Location{"Colonsay", "Europe/London"}
	case "CSB":
		return Location{"Caransebes", "Europe/Bucharest"}
	case "CSC":
		return Location{"Canas", "America/Costa_Rica"}
	case "CSF":
		return Location{"Creil", "Europe/Paris"}
	case "CSG":
		return Location{"Columbus", "America/New_York"}
	case "CSH":
		return Location{"Solovetsky Islands", "Europe/Moscow"}
	case "CSI":
		return Location{"", "Australia/Sydney"}
	case "CSK":
		return Location{"Cap Skirring", "Africa/Dakar"}
	case "CSL":
		return Location{"Cabo San Lucas", "America/Mazatlan"}
	case "CSM":
		return Location{"Clinton", "America/Chicago"}
	case "CSN":
		return Location{"Carson City", "America/Los_Angeles"}
	case "CSO":
		return Location{"Magdeburg", "Europe/Berlin"}
	case "CSQ":
		return Location{"Creston", "America/Chicago"}
	case "CSS":
		return Location{"Cassilandia", "America/Campo_Grande"}
	case "CSU":
		return Location{"Santa Cruz Do Sul", "America/Sao_Paulo"}
	case "CSV":
		return Location{"Crossville", "America/Chicago"}
	case "CSX":
		return Location{"Changsha", "Asia/Shanghai"}
	case "CSY":
		return Location{"Cheboksary", "Europe/Moscow"}
	case "CSZ":
		return Location{"Coronel Suarez", "America/Argentina/Buenos_Aires"}
	case "CTA":
		return Location{"Catania", "Europe/Rome"}
	case "CTB":
		return Location{"Cut Bank", "America/Denver"}
	case "CTC":
		return Location{"Catamarca", "America/Argentina/Catamarca"}
	case "CTD":
		return Location{"Chitre", "America/Panama"}
	case "CTF":
		return Location{"Coatepeque", "America/Guatemala"}
	case "CTG":
		return Location{"Cartagena", "America/Bogota"}
	case "CTH":
		return Location{"Coatesville", "America/New_York"}
	case "CTI":
		return Location{"Cuito Cuanavale", "Africa/Luanda"}
	case "CTL":
		return Location{"Charleville", "Australia/Brisbane"}
	case "CTM":
		return Location{"Chetumal", "America/Cancun"}
	case "CTN":
		return Location{"", "Australia/Brisbane"}
	case "CTP":
		return Location{"Carutapera", "America/Fortaleza"}
	case "CTQ":
		return Location{"Santa Vitoria Do Palmar", "America/Sao_Paulo"}
	case "CTS":
		return Location{"Chitose / Tomakomai", "Asia/Tokyo"}
	case "CTT":
		return Location{"Le Castellet", "Europe/Paris"}
	case "CTU":
		return Location{"Chengdu", "Asia/Shanghai"}
	case "CTY":
		return Location{"Cross City", "America/New_York"}
	case "CTZ":
		return Location{"Clinton", "America/New_York"}
	case "CUA":
		return Location{"Ciudad Constitucion", "America/Mazatlan"}
	case "CUB":
		return Location{"Columbia", "America/New_York"}
	case "CUC":
		return Location{"Cucuta", "America/Bogota"}
	case "CUD":
		return Location{"", "Australia/Brisbane"}
	case "CUE":
		return Location{"Cuenca", "America/Guayaquil"}
	case "CUF":
		return Location{"Cuneo", "Europe/Rome"}
	case "CUG":
		return Location{"", "Australia/Sydney"}
	case "CUH":
		return Location{"Cushing", "America/Chicago"}
	case "CUL":
		return Location{"Culiacan", "America/Mazatlan"}
	case "CUM":
		return Location{"", "America/Caracas"}
	case "CUN":
		return Location{"Cancun", "America/Cancun"}
	case "CUO":
		return Location{"Caruru", "America/Bogota"}
	case "CUP":
		return Location{"Carupano", "America/Caracas"}
	case "CUQ":
		return Location{"", "Australia/Brisbane"}
	case "CUR":
		return Location{"Willemstad", "America/Curacao"}
	case "CUT":
		return Location{"Cutral-Co", "America/Argentina/Salta"}
	case "CUU":
		return Location{"Chihuahua", "America/Chihuahua"}
	case "CUV":
		return Location{"Casigua El Cubo", "America/Caracas"}
	case "CUY":
		return Location{"", "Australia/Perth"}
	case "CUZ":
		return Location{"Cusco", "America/Lima"}
	case "CVC":
		return Location{"", "Australia/Adelaide"}
	case "CVE":
		return Location{"Covenas", "America/Bogota"}
	case "CVF":
		return Location{"Courcheval", "Europe/Paris"}
	case "CVG":
		return Location{"Hebron", "America/New_York"}
	case "CVH":
		return Location{"Lafontaine", "America/Argentina/Salta"}
	case "CVJ":
		return Location{"", "America/Mexico_City"}
	case "CVM":
		return Location{"Ciudad Victoria", "America/Monterrey"}
	case "CVN":
		return Location{"Clovis", "America/Denver"}
	case "CVO":
		return Location{"Corvallis", "America/Los_Angeles"}
	case "CVQ":
		return Location{"", "Australia/Perth"}
	case "CVS":
		return Location{"Clovis", "America/Denver"}
	case "CVT":
		return Location{"Coventry", "Europe/London"}
	case "CVU":
		return Location{"Corvo", "Atlantic/Azores"}
	case "CWA":
		return Location{"Mosinee", "America/Chicago"}
	case "CWB":
		return Location{"Curitiba", "America/Sao_Paulo"}
	case "CWC":
		return Location{"Chernivtsi", "Europe/Kyiv"}
	case "CWE":
		return Location{"El Cairo", "Africa/Cairo"}
	case "CWF":
		return Location{"Lake Charles", "America/Chicago"}
	case "CWI":
		return Location{"Clinton", "America/Chicago"}
	case "CWK":
		return Location{"Chitrakoot", "Asia/Kolkata"}
	case "CWL":
		return Location{"Cardiff", "Europe/London"}
	case "CWR":
		return Location{"", "Australia/Adelaide"}
	case "CWT":
		return Location{"", "Australia/Sydney"}
	case "CWW":
		return Location{"", "Australia/Melbourne"}
	case "CWX":
		return Location{"Willcox", "America/Phoenix"}
	case "CXA":
		return Location{"", "America/Caracas"}
	case "CXB":
		return Location{"Cox's Bazar", "Asia/Dhaka"}
	case "CXF":
		return Location{"Coldfoot", "America/Anchorage"}
	case "CXI":
		return Location{"Banana", "Pacific/Kiritimati"}
	case "CXJ":
		return Location{"Caxias Do Sul", "America/Sao_Paulo"}
	case "CXL":
		return Location{"Calexico", "America/Los_Angeles"}
	case "CXM":
		return Location{"Camaxilo", "Africa/Luanda"}
	case "CXN":
		return Location{"Candala", "Africa/Mogadishu"}
	case "CXO":
		return Location{"Houston", "America/Chicago"}
	case "CXP":
		return Location{"Cilacap-Java Island", "Asia/Jakarta"}
	case "CXQ":
		return Location{"", "Australia/Perth"}
	case "CXR":
		return Location{"Nha Trang", "Asia/Ho_Chi_Minh"}
	case "CXT":
		return Location{"", "Australia/Brisbane"}
	case "CXY":
		return Location{"Cat Cay", "America/Nassau"}
	case "CYA":
		return Location{"Les Cayes", "America/Port-au-Prince"}
	case "CYB":
		return Location{"Cayman Brac", "America/Cayman"}
	case "CYF":
		return Location{"Chefornak", "America/Nome"}
	case "CYG":
		return Location{"", "Australia/Melbourne"}
	case "CYI":
		return Location{"Chiayi City", "Asia/Taipei"}
	case "CYL":
		return Location{"Coyoles", "America/Tegucigalpa"}
	case "CYO":
		return Location{"Cayo Largo del Sur", "America/Havana"}
	case "CYP":
		return Location{"Calbayog City", "Asia/Manila"}
	case "CYR":
		return Location{"Colonia", "America/Montevideo"}
	case "CYS":
		return Location{"Cheyenne", "America/Denver"}
	case "CYT":
		return Location{"Yakataga", "America/Anchorage"}
	case "CYU":
		return Location{"Cuyo", "Asia/Manila"}
	case "CYW":
		return Location{"Celaya", "America/Mexico_City"}
	case "CYX":
		return Location{"Cherskiy", "Asia/Srednekolymsk"}
	case "CYZ":
		return Location{"Cauayan City", "Asia/Manila"}
	case "CZA":
		return Location{"", "America/Merida"}
	case "CZE":
		return Location{"Coro", "America/Caracas"}
	case "CZF":
		return Location{"Cape Romanzof", "America/Nome"}
	case "CZL":
		return Location{"Constantine", "Africa/Algiers"}
	case "CZM":
		return Location{"Cozumel", "America/Cancun"}
	case "CZS":
		return Location{"Cruzeiro Do Sul", "America/Rio_Branco"}
	case "CZT":
		return Location{"Carrizo Springs", "America/Chicago"}
	case "CZU":
		return Location{"Corozal", "America/Bogota"}
	case "CZW":
		return Location{"Czestochowa", "Europe/Warsaw"}
	case "CZX":
		return Location{"Changzhou", "Asia/Shanghai"}
	case "CZY":
		return Location{"", "Australia/Brisbane"}
	case "DAA":
		return Location{"Fort Belvoir", "America/New_York"}
	case "DAB":
		return Location{"Daytona Beach", "America/New_York"}
	case "DAC":
		return Location{"Dhaka", "Asia/Dhaka"}
	case "DAD":
		return Location{"Da Nang", "Asia/Ho_Chi_Minh"}
	case "DAG":
		return Location{"Daggett", "America/Los_Angeles"}
	case "DAK":
		return Location{"", "Africa/Cairo"}
	case "DAL":
		return Location{"Dallas", "America/Chicago"}
	case "DAM":
		return Location{"Damascus", "Asia/Damascus"}
	case "DAN":
		return Location{"Danville", "America/New_York"}
	case "DAO":
		return Location{"Dabo", "Pacific/Port_Moresby"}
	case "DAP":
		return Location{"Darchula", "Asia/Kathmandu"}
	case "DAR":
		return Location{"Dar es Salaam", "Africa/Dar_es_Salaam"}
	case "DAS":
		return Location{"Great Bear Lake", "America/Edmonton"}
	case "DAT":
		return Location{"Datong", "Asia/Shanghai"}
	case "DAU":
		return Location{"Daru", "Pacific/Port_Moresby"}
	case "DAV":
		return Location{"David", "America/Panama"}
	case "DAX":
		return Location{"Dazhou", "Asia/Shanghai"}
	case "DAY":
		return Location{"Dayton", "America/New_York"}
	case "DAZ":
		return Location{"Darwaz", "Asia/Dushanbe"}
	case "DBA":
		return Location{"Dalbandin", "Asia/Karachi"}
	case "DBB":
		return Location{"El Alamein", "Africa/Cairo"}
	case "DBD":
		return Location{"", "Asia/Kolkata"}
	case "DBM":
		return Location{"Debra Marcos", "Africa/Addis_Ababa"}
	case "DBN":
		return Location{"Dublin", "America/New_York"}
	case "DBO":
		return Location{"Dubbo", "Australia/Sydney"}
	case "DBP":
		return Location{"Debepare", "Pacific/Port_Moresby"}
	case "DBQ":
		return Location{"Dubuque", "America/Chicago"}
	case "DBR":
		return Location{"Darbhanga", "Asia/Kolkata"}
	case "DBT":
		return Location{"Debre Tabor", "Africa/Addis_Ababa"}
	case "DBV":
		return Location{"Dubrovnik", "Europe/Zagreb"}
	case "DBY":
		return Location{"", "Australia/Brisbane"}
	case "DCA":
		return Location{"Washington", "America/New_York"}
	case "DCF":
		return Location{"Canefield", "America/Dominica"}
	case "DCI":
		return Location{"Decimomannu", "Europe/Rome"}
	case "DCM":
		return Location{"Castres/Mazamet", "Europe/Paris"}
	case "DCN":
		return Location{"", "Australia/Perth"}
	case "DCT":
		return Location{"", "America/Nassau"}
	case "DCU":
		return Location{"Decatur", "America/Chicago"}
	case "DCY":
		return Location{"Daocheng County", "Asia/Shanghai"}
	case "DDC":
		return Location{"Dodge City", "America/Chicago"}
	case "DDD":
		return Location{"Kudahuvadhoo", "Indian/Maldives"}
	case "DDG":
		return Location{"Dandong", "Asia/Shanghai"}
	case "DDN":
		return Location{"", "Australia/Brisbane"}
	case "DDR":
		return Location{"", "Asia/Shanghai"}
	case "DDU":
		return Location{"Dadu", "Asia/Karachi"}
	case "DEA":
		return Location{"Dera Ghazi Khan", "Asia/Karachi"}
	case "DEB":
		return Location{"Debrecen", "Europe/Budapest"}
	case "DEC":
		return Location{"Decatur", "America/Chicago"}
	case "DED":
		return Location{"Dehradun", "Asia/Kolkata"}
	case "DEE":
		return Location{"Kunashir Island", "Asia/Magadan"}
	case "DEF":
		return Location{"", "Asia/Tehran"}
	case "DEH":
		return Location{"Decorah", "America/Chicago"}
	case "DEI":
		return Location{"Denis Island", "Indian/Mahe"}
	case "DEL":
		return Location{"New Delhi", "Asia/Kolkata"}
	case "DEM":
		return Location{"Dembidollo", "Africa/Addis_Ababa"}
	case "DEN":
		return Location{"Denver", "America/Denver"}
	case "DEP":
		return Location{"Daporijo", "Asia/Kolkata"}
	case "DER":
		return Location{"Derim", "Pacific/Port_Moresby"}
	case "DES":
		return Location{"Desroches Island", "Indian/Mahe"}
	case "DET":
		return Location{"Detroit", "America/Detroit"}
	case "DEZ":
		return Location{"Deir ez-Zor", "Asia/Damascus"}
	case "DFI":
		return Location{"Defiance", "America/New_York"}
	case "DFP":
		return Location{"Drumduff", "Australia/Brisbane"}
	case "DFW":
		return Location{"Dallas-Fort Worth", "America/Chicago"}
	case "DGA":
		return Location{"Dangriga", "America/Belize"}
	case "DGD":
		return Location{"", "Australia/Perth"}
	case "DGE":
		return Location{"Mudgee", "Australia/Sydney"}
	case "DGF":
		return Location{"Douglas Lake", "America/Vancouver"}
	case "DGH":
		return Location{"Deoghar", "Asia/Kolkata"}
	case "DGL":
		return Location{"Douglas", "America/Hermosillo"}
	case "DGN":
		return Location{"Dahlgren", "America/New_York"}
	case "DGO":
		return Location{"Durango", "America/Monterrey"}
	case "DGP":
		return Location{"Daugavpils", "Europe/Riga"}
	case "DGR":
		return Location{"", "Pacific/Auckland"}
	case "DGT":
		return Location{"Dumaguete City", "Asia/Manila"}
	case "DGU":
		return Location{"Dedougou", "Africa/Ouagadougou"}
	case "DGW":
		return Location{"Douglas", "America/Denver"}
	case "DGX":
		return Location{"St. Athan", "Europe/London"}
	case "DHA":
		return Location{"", "Asia/Riyadh"}
	case "DHD":
		return Location{"", "Australia/Brisbane"}
	case "DHF":
		return Location{"", "Asia/Dubai"}
	case "DHI":
		return Location{"Dhangarhi", "Asia/Kathmandu"}
	case "DHM":
		return Location{"", "Asia/Kolkata"}
	case "DHN":
		return Location{"Dothan", "America/Chicago"}
	case "DHR":
		return Location{"Den Helder", "Europe/Amsterdam"}
	case "DHT":
		return Location{"Dalhart", "America/Chicago"}
	case "DIA":
		return Location{"Doha", "Asia/Qatar"}
	case "DIB":
		return Location{"Dibrugarh", "Asia/Kolkata"}
	case "DIE":
		return Location{"", "Indian/Antananarivo"}
	case "DIG":
		return Location{"Shangri-La", "Asia/Shanghai"}
	case "DIJ":
		return Location{"Dijon/Longvic", "Europe/Paris"}
	case "DIK":
		return Location{"Dickinson", "America/Denver"}
	case "DIL":
		return Location{"Dili", "Asia/Dili"}
	case "DIM":
		return Location{"Dimbokro", "Africa/Abidjan"}
	case "DIN":
		return Location{"Dien Bien Phu", "Asia/Bangkok"}
	case "DIP":
		return Location{"Diapaga", "Africa/Ouagadougou"}
	case "DIQ":
		return Location{"Divinopolis", "America/Sao_Paulo"}
	case "DIR":
		return Location{"Dire Dawa", "Africa/Addis_Ababa"}
	case "DIS":
		return Location{"Dolisie", "Africa/Brazzaville"}
	case "DIU":
		return Location{"Diu", "Asia/Kolkata"}
	case "DIY":
		return Location{"Diyarbakir", "Europe/Istanbul"}
	case "DJA":
		return Location{"Djougou", "Africa/Porto-Novo"}
	case "DJB":
		return Location{"Jambi-Sumatra Island", "Asia/Jakarta"}
	case "DJE":
		return Location{"Djerba", "Africa/Tunis"}
	case "DJG":
		return Location{"Djanet", "Africa/Algiers"}
	case "DJJ":
		return Location{"Jayapura-Papua Island", "Asia/Jayapura"}
	case "DJM":
		return Location{"Djambala", "Africa/Brazzaville"}
	case "DJO":
		return Location{"", "Africa/Abidjan"}
	case "DJR":
		return Location{"", "Australia/Brisbane"}
	case "DJU":
		return Location{"Djupivogur", "Atlantic/Reykjavik"}
	case "DKI":
		return Location{"", "Australia/Brisbane"}
	case "DKK":
		return Location{"Dunkirk", "America/New_York"}
	case "DKR":
		return Location{"Dakar", "Africa/Dakar"}
	case "DKS":
		return Location{"Dikson", "Asia/Krasnoyarsk"}
	case "DKV":
		return Location{"", "Australia/Darwin"}
	case "DLA":
		return Location{"Douala", "Africa/Douala"}
	case "DLC":
		return Location{"Dalian", "Asia/Shanghai"}
	case "DLD":
		return Location{"Dagali", "Europe/Oslo"}
	case "DLE":
		return Location{"Dole/Tavaux", "Europe/Paris"}
	case "DLF":
		return Location{"Del Rio", "America/Chicago"}
	case "DLG":
		return Location{"Dillingham", "America/Anchorage"}
	case "DLH":
		return Location{"Duluth", "America/Chicago"}
	case "DLI":
		return Location{"Dalat", "Asia/Ho_Chi_Minh"}
	case "DLK":
		return Location{"Dulkaninna", "Australia/Adelaide"}
	case "DLL":
		return Location{"Dillon", "America/New_York"}
	case "DLM":
		return Location{"Dalaman", "Europe/Istanbul"}
	case "DLN":
		return Location{"Dillon", "America/Denver"}
	case "DLS":
		return Location{"The Dalles", "America/Los_Angeles"}
	case "DLU":
		return Location{"Xiaguan", "Asia/Shanghai"}
	case "DLV":
		return Location{"", "Australia/Darwin"}
	case "DLY":
		return Location{"Dillon's Bay", "Pacific/Efate"}
	case "DLZ":
		return Location{"Dalanzadgad", "Asia/Ulaanbaatar"}
	case "DMA":
		return Location{"Tucson", "America/Phoenix"}
	case "DMB":
		return Location{"Taraz", "Asia/Almaty"}
	case "DMD":
		return Location{"", "Australia/Brisbane"}
	case "DME":
		return Location{"Moscow", "Europe/Moscow"}
	case "DMK":
		return Location{"Bangkok", "Asia/Bangkok"}
	case "DMM":
		return Location{"Ad Dammam", "Asia/Riyadh"}
	case "DMN":
		return Location{"Deming", "America/Denver"}
	case "DMO":
		return Location{"Sedalia", "America/Chicago"}
	case "DMT":
		return Location{"Diamantino", "America/Cuiaba"}
	case "DMU":
		return Location{"Dimapur", "Asia/Kolkata"}
	case "DNA":
		return Location{"", "Asia/Tokyo"}
	case "DNB":
		return Location{"", "Australia/Brisbane"}
	case "DND":
		return Location{"Dundee", "Europe/London"}
	case "DNH":
		return Location{"Dunhuang", "Asia/Shanghai"}
	case "DNK":
		return Location{"Dnipropetrovsk", "Europe/Kyiv"}
	case "DNL":
		return Location{"Augusta", "America/New_York"}
	case "DNN":
		return Location{"Dalton", "America/New_York"}
	case "DNO":
		return Location{"Dianopolis", "America/Araguaina"}
	case "DNP":
		return Location{"Dang", "Asia/Kathmandu"}
	case "DNQ":
		return Location{"", "Australia/Sydney"}
	case "DNR":
		return Location{"Dinard/Pleurtuit/Saint-Malo", "Europe/Paris"}
	case "DNS":
		return Location{"Denison", "America/Chicago"}
	case "DNV":
		return Location{"Danville", "America/Chicago"}
	case "DNX":
		return Location{"Dinder", "Africa/Khartoum"}
	case "DNZ":
		return Location{"Denizli", "Europe/Istanbul"}
	case "DOB":
		return Location{"Dobo-Kobror Island", "Asia/Jayapura"}
	case "DOD":
		return Location{"Dodoma", "Africa/Dar_es_Salaam"}
	case "DOE":
		return Location{"Djumu-Djomoe", "America/Paramaribo"}
	case "DOG":
		return Location{"Dongola", "Africa/Khartoum"}
	case "DOH":
		return Location{"Doha", "Asia/Qatar"}
	case "DOK":
		return Location{"Donetsk", "Europe/Kyiv"}
	case "DOL":
		return Location{"Deauville", "Europe/Paris"}
	case "DOM":
		return Location{"Marigot", "America/Dominica"}
	case "DON":
		return Location{"Dos Lagunas", "America/Guatemala"}
	case "DOP":
		return Location{"Dolpa", "Asia/Kathmandu"}
	case "DOR":
		return Location{"Dori", "Africa/Ouagadougou"}
	case "DOU":
		return Location{"Dourados", "America/Campo_Grande"}
	case "DOV":
		return Location{"Dover", "America/New_York"}
	case "DOX":
		return Location{"", "Australia/Perth"}
	case "DOY":
		return Location{"Dongying", "Asia/Shanghai"}
	case "DPA":
		return Location{"Chicago/West Chicago", "America/Chicago"}
	case "DPE":
		return Location{"Dieppe", "Europe/Paris"}
	case "DPG":
		return Location{"Dugway Proving Ground", "America/Denver"}
	case "DPL":
		return Location{"Dipolog City", "Asia/Manila"}
	case "DPO":
		return Location{"Devonport", "Australia/Hobart"}
	case "DPS":
		return Location{"Denpasar-Bali Island", "Asia/Makassar"}
	case "DQA":
		return Location{"Daqing Shi", "Asia/Shanghai"}
	case "DQH":
		return Location{"Deadhorse", "America/Anchorage"}
	case "DQM":
		return Location{"Duqm", "Asia/Muscat"}
	case "DRA":
		return Location{"Mercury", "America/Los_Angeles"}
	case "DRB":
		return Location{"", "Australia/Perth"}
	case "DRD":
		return Location{"", "Australia/Brisbane"}
	case "DRE":
		return Location{"Drummond Island", "America/Detroit"}
	case "DRG":
		return Location{"Deering", "America/Nome"}
	case "DRI":
		return Location{"De Ridder", "America/Chicago"}
	case "DRJ":
		return Location{"Drietabbetje", "America/Paramaribo"}
	case "DRK":
		return Location{"Puntarenas", "America/Costa_Rica"}
	case "DRN":
		return Location{"", "Australia/Brisbane"}
	case "DRO":
		return Location{"Durango", "America/Denver"}
	case "DRP":
		return Location{"Daraga", "Asia/Manila"}
	case "DRR":
		return Location{"", "Australia/Brisbane"}
	case "DRS":
		return Location{"Dresden", "Europe/Berlin"}
	case "DRT":
		return Location{"Del Rio", "America/Chicago"}
	case "DRW":
		return Location{"Darwin", "Australia/Darwin"}
	case "DRY":
		return Location{"", "Australia/Perth"}
	case "DSA":
		return Location{"Doncaster", "Europe/London"}
	case "DSC":
		return Location{"Dschang", "Africa/Douala"}
	case "DSD":
		return Location{"Grande Anse", "America/Guadeloupe"}
	case "DSE":
		return Location{"Dessie", "Africa/Addis_Ababa"}
	case "DSI":
		return Location{"Destin", "America/Chicago"}
	case "DSK":
		return Location{"Dera Ismael Khan", "Asia/Karachi"}
	case "DSM":
		return Location{"Des Moines", "America/Chicago"}
	case "DSN":
		return Location{"Ordos", "Asia/Shanghai"}
	case "DSO":
		return Location{"", "Asia/Pyongyang"}
	case "DSS":
		return Location{"Diass", "Africa/Dakar"}
	case "DSV":
		return Location{"Dansville", "America/New_York"}
	case "DTA":
		return Location{"Delta", "America/Denver"}
	case "DTB":
		return Location{"Siborong-Borong", "Asia/Jakarta"}
	case "DTD":
		return Location{"Datadawai-Borneo Island", "Asia/Makassar"}
	case "DTE":
		return Location{"Daet", "Asia/Manila"}
	case "DTI":
		return Location{"Diamantina", "America/Sao_Paulo"}
	case "DTL":
		return Location{"Detroit Lakes", "America/Chicago"}
	case "DTM":
		return Location{"Dortmund", "Europe/Berlin"}
	case "DTN":
		return Location{"Shreveport", "America/Chicago"}
	case "DTR":
		return Location{"Decatur", "America/Los_Angeles"}
	case "DTW":
		return Location{"Detroit", "America/Detroit"}
	case "DUA":
		return Location{"Durant", "America/Chicago"}
	case "DUB":
		return Location{"Dublin", "Europe/Dublin"}
	case "DUC":
		return Location{"Duncan", "America/Chicago"}
	case "DUD":
		return Location{"Dunedin", "Pacific/Auckland"}
	case "DUE":
		return Location{"Chitato", "Africa/Luanda"}
	case "DUF":
		return Location{"Corolla", "America/New_York"}
	case "DUG":
		return Location{"Douglas Bisbee", "America/Phoenix"}
	case "DUJ":
		return Location{"Dubois", "America/New_York"}
	case "DUK":
		return Location{"Mubatuba", "Africa/Johannesburg"}
	case "DUM":
		return Location{"Dumai-Sumatra Island", "Asia/Jakarta"}
	case "DUQ":
		return Location{"Duncan", "America/Vancouver"}
	case "DUR":
		return Location{"Durban", "Africa/Johannesburg"}
	case "DUS":
		return Location{"Dusseldorf", "Europe/Berlin"}
	case "DUT":
		return Location{"Unalaska", "America/Nome"}
	case "DVK":
		return Location{"Diavik", "America/Edmonton"}
	case "DVL":
		return Location{"Devils Lake", "America/Chicago"}
	case "DVN":
		return Location{"Davenport", "America/Chicago"}
	case "DVO":
		return Location{"Davao City", "Asia/Manila"}
	case "DVP":
		return Location{"", "Australia/Brisbane"}
	case "DVR":
		return Location{"", "Australia/Darwin"}
	case "DVT":
		return Location{"Phoenix", "America/Phoenix"}
	case "DWA":
		return Location{"Dwangwa", "Africa/Blantyre"}
	case "DWB":
		return Location{"Soalala", "Indian/Antananarivo"}
	case "DWC":
		return Location{"Jebel Ali", "Asia/Dubai"}
	case "DWD":
		return Location{"Dawadmi", "Asia/Riyadh"}
	case "DWH":
		return Location{"Houston", "America/Chicago"}
	case "DWN":
		return Location{"Oklahoma City", "America/Chicago"}
	case "DXB":
		return Location{"Dubai", "Asia/Dubai"}
	case "DXD":
		return Location{"New Dixie", "Australia/Brisbane"}
	case "DXE":
		return Location{"Madison", "America/Chicago"}
	case "DXN":
		return Location{"Noida", "Asia/Kolkata"}
	case "DXR":
		return Location{"Danbury", "America/New_York"}
	case "DYA":
		return Location{"", "Australia/Brisbane"}
	case "DYG":
		return Location{"Dayong", "Asia/Shanghai"}
	case "DYL":
		return Location{"Doylestown", "America/New_York"}
	case "DYR":
		return Location{"Anadyr", "Asia/Anadyr"}
	case "DYS":
		return Location{"Abilene", "America/Chicago"}
	case "DYU":
		return Location{"Dushanbe", "Asia/Dushanbe"}
	case "DYW":
		return Location{"Daly Waters", "Australia/Darwin"}
	case "DZA":
		return Location{"Dzaoudzi", "Indian/Mayotte"}
	case "DZH":
		return Location{"Dazhou", "Asia/Shanghai"}
	case "DZN":
		return Location{"Zhezkazgan", "Asia/Almaty"}
	case "DZO":
		return Location{"Durazno", "America/Montevideo"}
	case "EAA":
		return Location{"Eagle", "America/Anchorage"}
	case "EAB":
		return Location{"Abs", "Asia/Aden"}
	case "EAE":
		return Location{"Sangafa", "Pacific/Efate"}
	case "EAM":
		return Location{"Nejran", "Asia/Riyadh"}
	case "EAN":
		return Location{"Wheatland", "America/Denver"}
	case "EAR":
		return Location{"Kearney", "America/Chicago"}
	case "EAS":
		return Location{"Hondarribia", "Europe/Madrid"}
	case "EAT":
		return Location{"Wenatchee", "America/Los_Angeles"}
	case "EAU":
		return Location{"Eau Claire", "America/Chicago"}
	case "EBA":
		return Location{"Marina  Di Campo", "Europe/Rome"}
	case "EBB":
		return Location{"Kampala", "Africa/Kampala"}
	case "EBD":
		return Location{"Al-Ubayyid", "Africa/Khartoum"}
	case "EBG":
		return Location{"El Bagre", "America/Bogota"}
	case "EBH":
		return Location{"El Bayadh", "Africa/Algiers"}
	case "EBJ":
		return Location{"Esbjerg", "Europe/Copenhagen"}
	case "EBL":
		return Location{"Arbil", "Asia/Baghdad"}
	case "EBM":
		return Location{"El Borma", "Africa/Tunis"}
	case "EBS":
		return Location{"Webster City", "America/Chicago"}
	case "EBU":
		return Location{"Saint-Etienne/Boutheon", "Europe/Paris"}
	case "EBW":
		return Location{"Ebolowa", "Africa/Douala"}
	case "ECG":
		return Location{"Elizabeth City", "America/New_York"}
	case "ECH":
		return Location{"", "Australia/Melbourne"}
	case "ECN":
		return Location{"Nicosia", "Asia/Famagusta"}
	case "ECP":
		return Location{"Panama City Beach", "America/Chicago"}
	case "ECS":
		return Location{"Newcastle", "America/Denver"}
	case "ECV":
		return Location{"Madrid", "Europe/Madrid"}
	case "EDB":
		return Location{"El Debba", "Africa/Khartoum"}
	case "EDE":
		return Location{"Edenton", "America/New_York"}
	case "EDF":
		return Location{"Anchorage", "America/Anchorage"}
	case "EDI":
		return Location{"Edinburgh", "Europe/London"}
	case "EDK":
		return Location{"El Dorado", "America/Chicago"}
	case "EDL":
		return Location{"Eldoret", "Africa/Nairobi"}
	case "EDM":
		return Location{"La Roche-sur-Yon/Les Ajoncs", "Europe/Paris"}
	case "EDO":
		return Location{"Edremit", "Europe/Istanbul"}
	case "EDR":
		return Location{"", "Australia/Brisbane"}
	case "EDW":
		return Location{"Edwards", "America/Los_Angeles"}
	case "EED":
		return Location{"Needles", "America/Los_Angeles"}
	case "EEK":
		return Location{"Eek", "America/Nome"}
	case "EEN":
		return Location{"Keene", "America/New_York"}
	case "EFD":
		return Location{"Houston", "America/Chicago"}
	case "EFK":
		return Location{"Newport", "America/New_York"}
	case "EFL":
		return Location{"Kefallinia Island", "Europe/Athens"}
	case "EFW":
		return Location{"Jefferson", "America/Chicago"}
	case "EGC":
		return Location{"Bergerac/Roumaniere", "Europe/Paris"}
	case "EGE":
		return Location{"Eagle", "America/Denver"}
	case "EGH":
		return Location{"", "Africa/Cairo"}
	case "EGI":
		return Location{"Crestview", "America/Chicago"}
	case "EGM":
		return Location{"Sege", "Pacific/Guadalcanal"}
	case "EGN":
		return Location{"Geneina", "Africa/Khartoum"}
	case "EGO":
		return Location{"Belgorod", "Europe/Moscow"}
	case "EGS":
		return Location{"Egilsstadir", "Atlantic/Reykjavik"}
	case "EGV":
		return Location{"Eagle River", "America/Chicago"}
	case "EGX":
		return Location{"Egegik", "America/Anchorage"}
	case "EHL":
		return Location{"El Bolson", "America/Argentina/Salta"}
	case "EHM":
		return Location{"Cape Newenham", "America/Nome"}
	case "EHU":
		return Location{"Ezhou", "Asia/Shanghai"}
	case "EIB":
		return Location{"Eisenach", "Europe/Berlin"}
	case "EIE":
		return Location{"Yeniseysk", "Asia/Krasnoyarsk"}
	case "EIH":
		return Location{"Einasleigh", "Australia/Brisbane"}
	case "EIK":
		return Location{"Yeysk", "Europe/Moscow"}
	case "EIL":
		return Location{"Fairbanks", "America/Anchorage"}
	case "EIN":
		return Location{"Eindhoven", "Europe/Amsterdam"}
	case "EIS":
		return Location{"Road Town", "America/Tortola"}
	case "EIY":
		return Location{"Sapir", "Asia/Amman"}
	case "EJA":
		return Location{"Barrancabermeja", "America/Bogota"}
	case "EJH":
		return Location{"Al Wajh", "Asia/Riyadh"}
	case "EJN":
		return Location{"", "Asia/Shanghai"}
	case "EKA":
		return Location{"Eureka", "America/Los_Angeles"}
	case "EKB":
		return Location{"Ekibastuz", "Asia/Almaty"}
	case "EKI":
		return Location{"Elkhart", "America/Indiana/Indianapolis"}
	case "EKN":
		return Location{"Elkins", "America/New_York"}
	case "EKO":
		return Location{"Elko", "America/Los_Angeles"}
	case "EKS":
		return Location{"Shakhtersk", "Asia/Sakhalin"}
	case "EKT":
		return Location{"Eskilstuna", "Europe/Stockholm"}
	case "EKX":
		return Location{"Elizabethtown", "America/New_York"}
	case "ELA":
		return Location{"Eagle Lake", "America/Chicago"}
	case "ELB":
		return Location{"El Banco", "America/Bogota"}
	case "ELC":
		return Location{"Elcho Island", "Australia/Darwin"}
	case "ELD":
		return Location{"El Dorado", "America/Chicago"}
	case "ELF":
		return Location{"El Fasher", "Africa/Khartoum"}
	case "ELG":
		return Location{"", "Africa/Algiers"}
	case "ELH":
		return Location{"North Eleuthera", "America/Nassau"}
	case "ELI":
		return Location{"Elim", "America/Nome"}
	case "ELK":
		return Location{"Elk City", "America/Chicago"}
	case "ELL":
		return Location{"Ellisras", "Africa/Johannesburg"}
	case "ELM":
		return Location{"Elmira/Corning", "America/New_York"}
	case "ELN":
		return Location{"Ellensburg", "America/Los_Angeles"}
	case "ELO":
		return Location{"El Dorado", "America/Argentina/Cordoba"}
	case "ELP":
		return Location{"El Paso", "America/Denver"}
	case "ELQ":
		return Location{"", "Asia/Riyadh"}
	case "ELS":
		return Location{"East London", "Africa/Johannesburg"}
	case "ELT":
		return Location{"", "Africa/Cairo"}
	case "ELU":
		return Location{"Guemar", "Africa/Algiers"}
	case "ELY":
		return Location{"Ely", "America/Los_Angeles"}
	case "ELZ":
		return Location{"Wellsville", "America/New_York"}
	case "EMA":
		return Location{"Nottingham", "Europe/London"}
	case "EMD":
		return Location{"Emerald", "Australia/Brisbane"}
	case "EME":
		return Location{"Emden", "Europe/Berlin"}
	case "EMG":
		return Location{"Empangeni", "Africa/Johannesburg"}
	case "EMK":
		return Location{"Emmonak", "America/Nome"}
	case "EML":
		return Location{"", "Europe/Zurich"}
	case "EMM":
		return Location{"Kemmerer", "America/Denver"}
	case "EMN":
		return Location{"Nema", "Africa/Nouakchott"}
	case "EMP":
		return Location{"Emporia", "America/Chicago"}
	case "EMT":
		return Location{"El Monte", "America/Los_Angeles"}
	case "EMX":
		return Location{"El Maiten", "America/Argentina/Catamarca"}
	case "ENA":
		return Location{"Kenai", "America/Anchorage"}
	case "ENB":
		return Location{"Eneabba", "Australia/Perth"}
	case "ENC":
		return Location{"Nancy/Essey", "Europe/Paris"}
	case "END":
		return Location{"Enid", "America/Chicago"}
	case "ENE":
		return Location{"Ende-Flores Island", "Asia/Makassar"}
	case "ENF":
		return Location{"Enontekio", "Europe/Helsinki"}
	case "ENH":
		return Location{"Enshi", "Asia/Shanghai"}
	case "ENI":
		return Location{"El Nido", "Asia/Manila"}
	case "ENK":
		return Location{"Enniskillen", "Europe/London"}
	case "ENL":
		return Location{"Centralia", "America/Chicago"}
	case "ENN":
		return Location{"Nenana", "America/Anchorage"}
	case "ENO":
		return Location{"Encarnacion", "America/Asuncion"}
	case "ENS":
		return Location{"Enschede", "Europe/Amsterdam"}
	case "ENT":
		return Location{"Eniwetok Atoll", "Pacific/Majuro"}
	case "ENU":
		return Location{"Enegu", "Africa/Lagos"}
	case "ENV":
		return Location{"Wendover", "America/Denver"}
	case "ENW":
		return Location{"Kenosha", "America/Chicago"}
	case "ENY":
		return Location{"Yan'an", "Asia/Shanghai"}
	case "EOH":
		return Location{"Medellin", "America/Bogota"}
	case "EOI":
		return Location{"Eday", "Europe/London"}
	case "EOK":
		return Location{"Keokuk", "America/Chicago"}
	case "EOR":
		return Location{"Bolivar", "America/Caracas"}
	case "EOS":
		return Location{"Neosho", "America/Chicago"}
	case "EOZ":
		return Location{"", "America/Caracas"}
	case "EPA":
		return Location{"El Palomar", "America/Argentina/Buenos_Aires"}
	case "EPH":
		return Location{"Ephrata", "America/Los_Angeles"}
	case "EPL":
		return Location{"Epinal/Mirecourt", "Europe/Paris"}
	case "EPR":
		return Location{"", "Australia/Perth"}
	case "EPS":
		return Location{"Arroyo Barril", "America/Santo_Domingo"}
	case "EPU":
		return Location{"Parnu", "Europe/Tallinn"}
	case "EQS":
		return Location{"Esquel", "America/Argentina/Catamarca"}
	case "ERA":
		return Location{"Erigavo", "Africa/Mogadishu"}
	case "ERB":
		return Location{"", "Australia/Adelaide"}
	case "ERC":
		return Location{"Erzincan", "Europe/Istanbul"}
	case "ERD":
		return Location{"Berdyansk", "Europe/Kyiv"}
	case "ERF":
		return Location{"Erfurt", "Europe/Berlin"}
	case "ERG":
		return Location{"Erbogachen", "Asia/Irkutsk"}
	case "ERH":
		return Location{"Errachidia", "Africa/Casablanca"}
	case "ERI":
		return Location{"Erie", "America/New_York"}
	case "ERL":
		return Location{"Erenhot", "Asia/Shanghai"}
	case "ERM":
		return Location{"Erechim", "America/Sao_Paulo"}
	case "ERN":
		return Location{"Eirunepe", "America/Eirunepe"}
	case "ERR":
		return Location{"Errol", "America/New_York"}
	case "ERS":
		return Location{"Windhoek", "Africa/Windhoek"}
	case "ERV":
		return Location{"Kerrville", "America/Chicago"}
	case "ERZ":
		return Location{"Erzurum", "Europe/Istanbul"}
	case "ESB":
		return Location{"Ankara", "Europe/Istanbul"}
	case "ESC":
		return Location{"Escanaba", "America/Detroit"}
	case "ESD":
		return Location{"Eastsound", "America/Los_Angeles"}
	case "ESE":
		return Location{"", "America/Tijuana"}
	case "ESF":
		return Location{"Alexandria", "America/Chicago"}
	case "ESG":
		return Location{"Mariscal Estigarribia", "America/Asuncion"}
	case "ESH":
		return Location{"Brighton", "Europe/London"}
	case "ESI":
		return Location{"Espinosa", "America/Sao_Paulo"}
	case "ESK":
		return Location{"Eskisehir", "Europe/Istanbul"}
	case "ESL":
		return Location{"Elista", "Europe/Moscow"}
	case "ESM":
		return Location{"Tachina", "America/Guayaquil"}
	case "ESN":
		return Location{"Easton", "America/New_York"}
	case "ESR":
		return Location{"El Salvador", "America/Santiago"}
	case "ESS":
		return Location{"", "Europe/Berlin"}
	case "EST":
		return Location{"Estherville", "America/Chicago"}
	case "ESU":
		return Location{"Essaouira", "Africa/Casablanca"}
	case "ESW":
		return Location{"Easton", "America/Los_Angeles"}
	case "ETB":
		return Location{"West Bend", "America/Chicago"}
	case "ETD":
		return Location{"Etadunna", "Australia/Adelaide"}
	case "ETE":
		return Location{"Metema", "Africa/Addis_Ababa"}
	case "ETH":
		return Location{"Eilat", "Asia/Amman"}
	case "ETM":
		return Location{"Eilat", "Asia/Amman"}
	case "ETN":
		return Location{"Eastland", "America/Chicago"}
	case "ETR":
		return Location{"Santa Rosa", "America/Guayaquil"}
	case "ETS":
		return Location{"Enterprise", "America/Chicago"}
	case "ETZ":
		return Location{"Metz / Nancy", "Europe/Paris"}
	case "EUA":
		return Location{"Eua Island", "Pacific/Tongatapu"}
	case "EUC":
		return Location{"", "Australia/Eucla"}
	case "EUF":
		return Location{"Eufaula", "America/Chicago"}
	case "EUG":
		return Location{"Eugene", "America/Los_Angeles"}
	case "EUM":
		return Location{"Neumunster", "Europe/Berlin"}
	case "EUN":
		return Location{"El Aaiun", "Africa/El_Aaiun"}
	case "EUQ":
		return Location{"San Jose", "Asia/Manila"}
	case "EUX":
		return Location{"Sint Eustatius", "America/Kralendijk"}
	case "EVD":
		return Location{"Eva Downs", "Australia/Darwin"}
	case "EVE":
		return Location{"Evenes", "Europe/Oslo"}
	case "EVG":
		return Location{"", "Europe/Stockholm"}
	case "EVH":
		return Location{"", "Australia/Sydney"}
	case "EVM":
		return Location{"Eveleth", "America/Chicago"}
	case "EVN":
		return Location{"Yerevan", "Asia/Yerevan"}
	case "EVV":
		return Location{"Evansville", "America/Chicago"}
	case "EVW":
		return Location{"Evanston", "America/Denver"}
	case "EVX":
		return Location{"Evreux/Fauville", "Europe/Paris"}
	case "EWB":
		return Location{"New Bedford", "America/New_York"}
	case "EWI":
		return Location{"Enarotali-Papua Island", "Asia/Jayapura"}
	case "EWK":
		return Location{"Newton", "America/Chicago"}
	case "EWN":
		return Location{"New Bern", "America/New_York"}
	case "EWO":
		return Location{"Ewo", "Africa/Brazzaville"}
	case "EWR":
		return Location{"Newark", "America/New_York"}
	case "EXM":
		return Location{"", "Australia/Perth"}
	case "EXT":
		return Location{"Exeter", "Europe/London"}
	case "EYL":
		return Location{"Yelimane", "Africa/Bamako"}
	case "EYP":
		return Location{"El Yopal", "America/Bogota"}
	case "EYR":
		return Location{"Yerington", "America/Los_Angeles"}
	case "EYS":
		return Location{"Eliye Springs", "Africa/Nairobi"}
	case "EYW":
		return Location{"Key West", "America/New_York"}
	case "EZE":
		return Location{"Ezeiza", "America/Argentina/Buenos_Aires"}
	case "EZS":
		return Location{"Elazig", "Europe/Istanbul"}
	case "EZV":
		return Location{"", "Asia/Yekaterinburg"}
	case "FAA":
		return Location{"", "Africa/Conakry"}
	case "FAB":
		return Location{"Farnborough", "Europe/London"}
	case "FAC":
		return Location{"", "Pacific/Tahiti"}
	case "FAE":
		return Location{"Vagar", "Atlantic/Faroe"}
	case "FAF":
		return Location{"Fort Eustis", "America/New_York"}
	case "FAG":
		return Location{"Fagurholsmyri", "Atlantic/Reykjavik"}
	case "FAH":
		return Location{"Farah", "Asia/Kabul"}
	case "FAI":
		return Location{"Fairbanks", "America/Anchorage"}
	case "FAJ":
		return Location{"Fajardo", "America/Puerto_Rico"}
	case "FAM":
		return Location{"Farmington", "America/Chicago"}
	case "FAN":
		return Location{"Farsund", "Europe/Oslo"}
	case "FAO":
		return Location{"Faro", "Europe/Lisbon"}
	case "FAQ":
		return Location{"Frieda River", "Pacific/Port_Moresby"}
	case "FAR":
		return Location{"Fargo", "America/Chicago"}
	case "FAS":
		return Location{"Faskrudsfjordur", "Atlantic/Reykjavik"}
	case "FAT":
		return Location{"Fresno", "America/Los_Angeles"}
	case "FAU":
		return Location{"Fahud", "Asia/Muscat"}
	case "FAV":
		return Location{"", "Pacific/Tahiti"}
	case "FAY":
		return Location{"Fayetteville", "America/New_York"}
	case "FAZ":
		return Location{"Fasa", "Asia/Tehran"}
	case "FBA":
		return Location{"Fonte Boa", "America/Manaus"}
	case "FBD":
		return Location{"Faizabad", "Asia/Kabul"}
	case "FBE":
		return Location{"Francisco Beltrao", "America/Sao_Paulo"}
	case "FBG":
		return Location{"Fort Bragg", "America/New_York"}
	case "FBK":
		return Location{"Fairbanks/Ft Wainwright", "America/Anchorage"}
	case "FBL":
		return Location{"Faribault", "America/Chicago"}
	case "FBM":
		return Location{"Lubumbashi", "Africa/Lubumbashi"}
	case "FBR":
		return Location{"Fort Bridger", "America/Denver"}
	case "FBY":
		return Location{"Fairbury", "America/Chicago"}
	case "FCA":
		return Location{"Kalispell", "America/Denver"}
	case "FCB":
		return Location{"Ficksburg", "Africa/Johannesburg"}
	case "FCH":
		return Location{"Fresno", "America/Los_Angeles"}
	case "FCM":
		return Location{"Minneapolis", "America/Chicago"}
	case "FCO":
		return Location{"Rome", "Europe/Rome"}
	case "FCS":
		return Location{"Fort Carson", "America/Denver"}
	case "FCY":
		return Location{"Forrest City", "America/Chicago"}
	case "FDA":
		return Location{"Fundacion", "America/Bogota"}
	case "FDE":
		return Location{"Forde", "Europe/Oslo"}
	case "FDF":
		return Location{"Fort-de-France", "America/Martinique"}
	case "FDH":
		return Location{"Friedrichshafen", "Europe/Berlin"}
	case "FDK":
		return Location{"Frederick", "America/New_York"}
	case "FDO":
		return Location{"San Fernando", "America/Argentina/Buenos_Aires"}
	case "FDR":
		return Location{"Frederick", "America/Chicago"}
	case "FDU":
		return Location{"", "Africa/Kinshasa"}
	case "FDY":
		return Location{"Findlay", "America/New_York"}
	case "FEB":
		return Location{"Sanfebagar", "Asia/Kathmandu"}
	case "FEC":
		return Location{"Feira De Santana", "America/Bahia"}
	case "FEG":
		return Location{"Fergana", "Asia/Tashkent"}
	case "FEJ":
		return Location{"Feijo", "America/Rio_Branco"}
	case "FEK":
		return Location{"Ferkessedougou", "Africa/Abidjan"}
	case "FEL":
		return Location{"Furstenfeldbruck", "Europe/Berlin"}
	case "FEN":
		return Location{"Fernando De Noronha", "America/Noronha"}
	case "FEP":
		return Location{"Freeport", "America/Chicago"}
	case "FET":
		return Location{"Fremont", "America/Chicago"}
	case "FEZ":
		return Location{"Fes", "Africa/Casablanca"}
	case "FFA":
		return Location{"Kill Devil Hills", "America/New_York"}
	case "FFD":
		return Location{"Fairford", "Europe/London"}
	case "FFL":
		return Location{"Fairfield", "America/Chicago"}
	case "FFM":
		return Location{"Fergus Falls", "America/Chicago"}
	case "FFO":
		return Location{"Dayton", "America/New_York"}
	case "FFT":
		return Location{"Frankfort", "America/New_York"}
	case "FFU":
		return Location{"Futaleufu", "America/Santiago"}
	case "FGD":
		return Location{"Fderik", "Africa/Nouakchott"}
	case "FGI":
		return Location{"Apia", "Pacific/Apia"}
	case "FGU":
		return Location{"", "Pacific/Tahiti"}
	case "FHB":
		return Location{"Fernandina Beach", "America/New_York"}
	case "FHU":
		return Location{"Fort Huachuca Sierra Vista", "America/Phoenix"}
	case "FHZ":
		return Location{"Fakahina", "Pacific/Tahiti"}
	case "FID":
		return Location{"Fishers Island", "America/New_York"}
	case "FIE":
		return Location{"Fair Isle", "Europe/London"}
	case "FIG":
		return Location{"", "Africa/Conakry"}
	case "FIH":
		return Location{"Kinshasa", "Africa/Kinshasa"}
	case "FIK":
		return Location{"Finke", "Australia/Darwin"}
	case "FIL":
		return Location{"Fillmore", "America/Denver"}
	case "FIN":
		return Location{"Buki", "Pacific/Port_Moresby"}
	case "FIZ":
		return Location{"", "Australia/Perth"}
	case "FJR":
		return Location{"", "Asia/Dubai"}
	case "FKB":
		return Location{"Baden-Baden", "Europe/Berlin"}
	case "FKI":
		return Location{"Kisangani", "Africa/Lubumbashi"}
	case "FKJ":
		return Location{"", "Asia/Tokyo"}
	case "FKL":
		return Location{"Franklin", "America/New_York"}
	case "FKN":
		return Location{"Franklin", "America/New_York"}
	case "FKQ":
		return Location{"Fakfak-Papua Island", "Asia/Jayapura"}
	case "FKS":
		return Location{"Sukagawa", "Asia/Tokyo"}
	case "FLA":
		return Location{"Florencia", "America/Bogota"}
	case "FLB":
		return Location{"Floriano", "America/Fortaleza"}
	case "FLD":
		return Location{"Fond du Lac", "America/Chicago"}
	case "FLF":
		return Location{"Flensburg", "Europe/Berlin"}
	case "FLG":
		return Location{"Flagstaff", "America/Phoenix"}
	case "FLI":
		return Location{"Flateyri", "Atlantic/Reykjavik"}
	case "FLL":
		return Location{"Fort Lauderdale", "America/New_York"}
	case "FLM":
		return Location{"Filadelfia", "America/Asuncion"}
	case "FLN":
		return Location{"Florianopolis", "America/Sao_Paulo"}
	case "FLO":
		return Location{"Florence", "America/New_York"}
	case "FLP":
		return Location{"Flippin", "America/Chicago"}
	case "FLR":
		return Location{"Firenze", "Europe/Rome"}
	case "FLS":
		return Location{"", "Australia/Hobart"}
	case "FLV":
		return Location{"Fort Leavenworth", "America/Chicago"}
	case "FLW":
		return Location{"Santa Cruz das Flores", "Atlantic/Azores"}
	case "FLX":
		return Location{"Fallon", "America/Los_Angeles"}
	case "FLY":
		return Location{"", "Australia/Sydney"}
	case "FLZ":
		return Location{"Sibolga-Sumatra Island", "Asia/Jakarta"}
	case "FMA":
		return Location{"Formosa", "America/Argentina/Cordoba"}
	case "FMC":
		return Location{"Five Mile", "America/Anchorage"}
	case "FME":
		return Location{"Fort Meade(Odenton)", "America/New_York"}
	case "FMG":
		return Location{"Brasilito", "America/Costa_Rica"}
	case "FMH":
		return Location{"Falmouth", "America/New_York"}
	case "FMI":
		return Location{"", "Africa/Lubumbashi"}
	case "FMM":
		return Location{"Memmingen", "Europe/Berlin"}
	case "FMN":
		return Location{"Farmington", "America/Denver"}
	case "FMO":
		return Location{"Munster", "Europe/Berlin"}
	case "FMS":
		return Location{"Fort Madison", "America/Chicago"}
	case "FMU":
		return Location{"Florence", "America/Los_Angeles"}
	case "FMY":
		return Location{"Fort Myers", "America/New_York"}
	case "FNA":
		return Location{"Freetown", "Africa/Freetown"}
	case "FNB":
		return Location{"", "Europe/Berlin"}
	case "FNC":
		return Location{"Funchal", "Atlantic/Madeira"}
	case "FND":
		return Location{"Funadhoo", "Indian/Maldives"}
	case "FNE":
		return Location{"Fane Mission", "Pacific/Port_Moresby"}
	case "FNG":
		return Location{"Fada N'gourma", "Africa/Ouagadougou"}
	case "FNH":
		return Location{"Fincha", "Africa/Addis_Ababa"}
	case "FNI":
		return Location{"Nimes/Garons", "Europe/Paris"}
	case "FNJ":
		return Location{"Pyongyang", "Asia/Pyongyang"}
	case "FNL":
		return Location{"Fort Collins/Loveland", "America/Denver"}
	case "FNT":
		return Location{"Flint", "America/Detroit"}
	case "FNU":
		return Location{"Oristano", "Europe/Rome"}
	case "FOB":
		return Location{"Fort Bragg", "America/Los_Angeles"}
	case "FOC":
		return Location{"Fuzhou", "Asia/Shanghai"}
	case "FOD":
		return Location{"Fort Dodge", "America/Chicago"}
	case "FOE":
		return Location{"Topeka", "America/Chicago"}
	case "FOG":
		return Location{"Foggia", "Europe/Rome"}
	case "FOK":
		return Location{"Westhampton Beach", "America/New_York"}
	case "FOM":
		return Location{"Foumban", "Africa/Douala"}
	case "FON":
		return Location{"La Fortuna/San Carlos", "America/Costa_Rica"}
	case "FOO":
		return Location{"Kornasoren-Numfoor Island", "Asia/Jayapura"}
	case "FOR":
		return Location{"Fortaleza", "America/Fortaleza"}
	case "FOS":
		return Location{"", "Australia/Perth"}
	case "FOT":
		return Location{"", "Australia/Sydney"}
	case "FOU":
		return Location{"Fougamou", "Africa/Libreville"}
	case "FPO":
		return Location{"Freeport", "America/Nassau"}
	case "FPR":
		return Location{"Fort Pierce", "America/New_York"}
	case "FPY":
		return Location{"Perry", "America/New_York"}
	case "FRA":
		return Location{"Frankfurt-am-Main", "Europe/Berlin"}
	case "FRB":
		return Location{"Forbes", "Australia/Sydney"}
	case "FRC":
		return Location{"Franca", "America/Sao_Paulo"}
	case "FRD":
		return Location{"Friday Harbor", "America/Los_Angeles"}
	case "FRE":
		return Location{"Fera Island", "Pacific/Guadalcanal"}
	case "FRF":
		return Location{"Oschersleben", "Europe/Berlin"}
	case "FRG":
		return Location{"Farmingdale", "America/New_York"}
	case "FRH":
		return Location{"French Lick", "America/Indiana/Indianapolis"}
	case "FRI":
		return Location{"Fort Riley(Junction City)", "America/Chicago"}
	case "FRJ":
		return Location{"Hyeres", "Europe/Paris"}
	case "FRK":
		return Location{"Fregate Island", "Indian/Mahe"}
	case "FRL":
		return Location{"Forli", "Europe/Rome"}
	case "FRM":
		return Location{"Fairmont", "America/Chicago"}
	case "FRN":
		return Location{"Fort Richardson (Anchorage)", "America/Anchorage"}
	case "FRO":
		return Location{"Floro", "Europe/Oslo"}
	case "FRR":
		return Location{"Front Royal", "America/New_York"}
	case "FRS":
		return Location{"San Benito", "America/Guatemala"}
	case "FRT":
		return Location{"Frutillar", "America/Santiago"}
	case "FRW":
		return Location{"Francistown", "Africa/Gaborone"}
	case "FRY":
		return Location{"Fryeburg", "America/New_York"}
	case "FRZ":
		return Location{"Fritzlar", "Europe/Berlin"}
	case "FSC":
		return Location{"Figari Sud-Corse", "Europe/Paris"}
	case "FSD":
		return Location{"Sioux Falls", "America/Chicago"}
	case "FSI":
		return Location{"Fort Sill", "America/Chicago"}
	case "FSK":
		return Location{"Fort Scott", "America/Chicago"}
	case "FSM":
		return Location{"Fort Smith", "America/Chicago"}
	case "FSP":
		return Location{"Saint-Pierre", "America/Miquelon"}
	case "FSS":
		return Location{"Kinloss", "Europe/London"}
	case "FST":
		return Location{"Fort Stockton", "America/Chicago"}
	case "FSU":
		return Location{"Fort Sumner", "America/Denver"}
	case "FSZ":
		return Location{"", "Asia/Tokyo"}
	case "FTA":
		return Location{"Futuna Island", "Pacific/Efate"}
	case "FTE":
		return Location{"El Calafate", "America/Argentina/Rio_Gallegos"}
	case "FTI":
		return Location{"Fitiuta Village", "Pacific/Pago_Pago"}
	case "FTK":
		return Location{"Fort Knox", "America/New_York"}
	case "FTU":
		return Location{"Tolanaro", "Indian/Antananarivo"}
	case "FTW":
		return Location{"Fort Worth", "America/Chicago"}
	case "FTX":
		return Location{"Owando", "Africa/Brazzaville"}
	case "FTY":
		return Location{"Atlanta", "America/New_York"}
	case "FUE":
		return Location{"Fuerteventura Island", "Atlantic/Canary"}
	case "FUG":
		return Location{"Fuyang", "Asia/Shanghai"}
	case "FUJ":
		return Location{"Goto", "Asia/Tokyo"}
	case "FUK":
		return Location{"Fukuoka", "Asia/Tokyo"}
	case "FUL":
		return Location{"Fullerton", "America/Los_Angeles"}
	case "FUN":
		return Location{"Funafuti", "Pacific/Funafuti"}
	case "FUO":
		return Location{"Foshan", "Asia/Shanghai"}
	case "FUT":
		return Location{"Futuna Island", "Pacific/Wallis"}
	case "FVL":
		return Location{"", "Australia/Perth"}
	case "FVM":
		return Location{"Fuvahmulah Island", "Indian/Maldives"}
	case "FWA":
		return Location{"Fort Wayne", "America/Indiana/Indianapolis"}
	case "FWH":
		return Location{"Fort Worth", "America/Chicago"}
	case "FWL":
		return Location{"Farewell", "America/Anchorage"}
	case "FXE":
		return Location{"Fort Lauderdale", "America/New_York"}
	case "FXO":
		return Location{"Cuamba", "Africa/Maputo"}
	case "FXY":
		return Location{"Forest City", "America/Chicago"}
	case "FYM":
		return Location{"Fayetteville", "America/Chicago"}
	case "FYT":
		return Location{"", "Africa/Ndjamena"}
	case "FYU":
		return Location{"Fort Yukon", "America/Anchorage"}
	case "FYV":
		return Location{"Fayetteville", "America/Chicago"}
	case "FZO":
		return Location{"Bristol", "Europe/London"}
	case "GAB":
		return Location{"Gabbs", "America/Los_Angeles"}
	case "GAC":
		return Location{"El Molino", "America/Tegucigalpa"}
	case "GAD":
		return Location{"Gadsden", "America/Chicago"}
	case "GAE":
		return Location{"Gabes", "Africa/Tunis"}
	case "GAF":
		return Location{"Gafsa", "Africa/Tunis"}
	case "GAG":
		return Location{"Gage", "America/Chicago"}
	case "GAH":
		return Location{"", "Australia/Brisbane"}
	case "GAI":
		return Location{"Gaithersburg", "America/New_York"}
	case "GAJ":
		return Location{"Yamagata", "Asia/Tokyo"}
	case "GAL":
		return Location{"Galena", "America/Anchorage"}
	case "GAM":
		return Location{"Gambell", "America/Nome"}
	case "GAN":
		return Location{"Gan", "Indian/Maldives"}
	case "GAO":
		return Location{"Guantanamo", "America/Havana"}
	case "GAP":
		return Location{"Gusap", "Pacific/Port_Moresby"}
	case "GAQ":
		return Location{"", "Africa/Bamako"}
	case "GAR":
		return Location{"Garaina", "Pacific/Port_Moresby"}
	case "GAS":
		return Location{"Garissa", "Africa/Nairobi"}
	case "GAT":
		return Location{"Avignon", "Europe/Paris"}
	case "GAU":
		return Location{"Guwahati", "Asia/Kolkata"}
	case "GAW":
		return Location{"Gangaw", "Asia/Yangon"}
	case "GAY":
		return Location{"", "Asia/Kolkata"}
	case "GBA":
		return Location{"Kemble", "Europe/London"}
	case "GBB":
		return Location{"Gabala", "Asia/Baku"}
	case "GBD":
		return Location{"Great Bend", "America/Chicago"}
	case "GBE":
		return Location{"Gaborone", "Africa/Gaborone"}
	case "GBF":
		return Location{"Negarbo", "Pacific/Port_Moresby"}
	case "GBG":
		return Location{"Galesburg", "America/Chicago"}
	case "GBH":
		return Location{"Galbraith Lake", "America/Anchorage"}
	case "GBI":
		return Location{"Kalaburagi", "Asia/Kolkata"}
	case "GBJ":
		return Location{"Grand Bourg", "America/Guadeloupe"}
	case "GBK":
		return Location{"Gbangbatok", "Africa/Freetown"}
	case "GBL":
		return Location{"", "Australia/Darwin"}
	case "GBP":
		return Location{"", "Australia/Brisbane"}
	case "GBR":
		return Location{"Great Barrington", "America/New_York"}
	case "GBT":
		return Location{"Gorgan", "Asia/Tehran"}
	case "GBU":
		return Location{"Khashm El Girba", "Africa/Khartoum"}
	case "GBV":
		return Location{"", "Australia/Perth"}
	case "GBW":
		return Location{"Ginbata", "Australia/Perth"}
	case "GBZ":
		return Location{"Claris", "Pacific/Auckland"}
	case "GCC":
		return Location{"Gillette", "America/Denver"}
	case "GCD":
		return Location{"Electric City", "America/Los_Angeles"}
	case "GCH":
		return Location{"", "Asia/Tehran"}
	case "GCI":
		return Location{"Saint Peter Port", "Europe/Guernsey"}
	case "GCJ":
		return Location{"Midrand", "Africa/Johannesburg"}
	case "GCK":
		return Location{"Garden City", "America/Chicago"}
	case "GCM":
		return Location{"Georgetown", "America/Cayman"}
	case "GCN":
		return Location{"Grand Canyon", "America/Phoenix"}
	case "GCW":
		return Location{"Peach Springs", "America/Phoenix"}
	case "GCY":
		return Location{"Greeneville", "America/New_York"}
	case "GDC":
		return Location{"Greenville", "America/New_York"}
	case "GDD":
		return Location{"Gordon Downs", "Australia/Perth"}
	case "GDE":
		return Location{"Gode", "Africa/Addis_Ababa"}
	case "GDG":
		return Location{"Magdagachi", "Asia/Yakutsk"}
	case "GDI":
		return Location{"Melle", "Africa/Bangui"}
	case "GDJ":
		return Location{"Gandajika", "Africa/Lubumbashi"}
	case "GDL":
		return Location{"Guadalajara", "America/Mexico_City"}
	case "GDM":
		return Location{"Gardner", "America/New_York"}
	case "GDN":
		return Location{"Gdansk", "Europe/Warsaw"}
	case "GDO":
		return Location{"", "America/Caracas"}
	case "GDP":
		return Location{"Guadalupe", "America/Fortaleza"}
	case "GDQ":
		return Location{"Gondar", "Africa/Addis_Ababa"}
	case "GDT":
		return Location{"Cockburn Town", "America/Grand_Turk"}
	case "GDV":
		return Location{"Glendive", "America/Denver"}
	case "GDW":
		return Location{"Gladwin", "America/Detroit"}
	case "GDX":
		return Location{"Magadan", "Asia/Magadan"}
	case "GDY":
		return Location{"Grundy", "America/New_York"}
	case "GDZ":
		return Location{"Gelendzhik", "Europe/Moscow"}
	case "GEA":
		return Location{"Noumea", "Pacific/Noumea"}
	case "GEB":
		return Location{"Gebe Island", "Asia/Jayapura"}
	case "GED":
		return Location{"Georgetown", "America/New_York"}
	case "GEE":
		return Location{"", "Australia/Hobart"}
	case "GEF":
		return Location{"Liangia", "Pacific/Guadalcanal"}
	case "GEG":
		return Location{"Spokane", "America/Los_Angeles"}
	case "GEL":
		return Location{"Santo Angelo", "America/Sao_Paulo"}
	case "GEO":
		return Location{"Georgetown", "America/Guyana"}
	case "GER":
		return Location{"Nueva Gerona", "America/Havana"}
	case "GES":
		return Location{"General Santos", "Asia/Manila"}
	case "GET":
		return Location{"", "Australia/Perth"}
	case "GEV":
		return Location{"Gallivare", "Europe/Stockholm"}
	case "GEX":
		return Location{"", "Australia/Melbourne"}
	case "GEY":
		return Location{"Greybull", "America/Denver"}
	case "GFD":
		return Location{"Greenfield", "America/Indiana/Indianapolis"}
	case "GFF":
		return Location{"Griffith", "Australia/Sydney"}
	case "GFK":
		return Location{"Grand Forks", "America/Chicago"}
	case "GFL":
		return Location{"Glens Falls", "America/New_York"}
	case "GFN":
		return Location{"", "Australia/Sydney"}
	case "GFO":
		return Location{"Bartica", "America/Guyana"}
	case "GFR":
		return Location{"Granville", "Europe/Paris"}
	case "GFY":
		return Location{"Grootfontein", "Africa/Windhoek"}
	case "GGB":
		return Location{"Agua Boa", "America/Cuiaba"}
	case "GGD":
		return Location{"", "Australia/Brisbane"}
	case "GGE":
		return Location{"Georgetown", "America/New_York"}
	case "GGF":
		return Location{"Almeirim", "America/Santarem"}
	case "GGG":
		return Location{"Longview", "America/Chicago"}
	case "GGH":
		return Location{"Cianorte", "America/Sao_Paulo"}
	case "GGM":
		return Location{"Kakamega", "Africa/Nairobi"}
	case "GGN":
		return Location{"Gagnoa", "Africa/Abidjan"}
	case "GGO":
		return Location{"Guiglo", "Africa/Abidjan"}
	case "GGS":
		return Location{"Gobernador Gregores", "America/Argentina/Rio_Gallegos"}
	case "GGT":
		return Location{"George Town", "America/Nassau"}
	case "GGW":
		return Location{"Glasgow", "America/Denver"}
	case "GHA":
		return Location{"Ghardaia", "Africa/Algiers"}
	case "GHB":
		return Location{"Governor's Harbour", "America/Nassau"}
	case "GHC":
		return Location{"", "America/Nassau"}
	case "GHF":
		return Location{"", "Europe/Berlin"}
	case "GHM":
		return Location{"Centerville", "America/Chicago"}
	case "GHT":
		return Location{"Ghat", "Africa/Tripoli"}
	case "GHU":
		return Location{"Gualeguaychu", "America/Argentina/Cordoba"}
	case "GHV":
		return Location{"Brasov", "Europe/Bucharest"}
	case "GIB":
		return Location{"Gibraltar", "Europe/Gibraltar"}
	case "GIC":
		return Location{"", "Australia/Brisbane"}
	case "GID":
		return Location{"Gitega", "Africa/Bujumbura"}
	case "GIF":
		return Location{"Winter Haven", "America/New_York"}
	case "GIG":
		return Location{"Rio De Janeiro", "America/Sao_Paulo"}
	case "GII":
		return Location{"Siguiri", "Africa/Conakry"}
	case "GIL":
		return Location{"Gilgit", "Asia/Karachi"}
	case "GIR":
		return Location{"Girardot", "America/Bogota"}
	case "GIS":
		return Location{"Gisborne", "Pacific/Auckland"}
	case "GIU":
		return Location{"Sigiriya", "Asia/Colombo"}
	case "GIY":
		return Location{"Giyani", "Africa/Johannesburg"}
	case "GIZ":
		return Location{"Jizan", "Asia/Riyadh"}
	case "GJA":
		return Location{"Guanaja", "America/Tegucigalpa"}
	case "GJL":
		return Location{"Jijel", "Africa/Algiers"}
	case "GJM":
		return Location{"Guajara-Mirim", "America/Porto_Velho"}
	case "GJR":
		return Location{"Gjogur", "Atlantic/Reykjavik"}
	case "GJT":
		return Location{"Grand Junction", "America/Denver"}
	case "GKA":
		return Location{"Goronka", "Pacific/Port_Moresby"}
	case "GKD":
		return Location{"Gokceada", "Europe/Istanbul"}
	case "GKE":
		return Location{"", "Europe/Amsterdam"}
	case "GKH":
		return Location{"Gorkha", "Asia/Kathmandu"}
	case "GKK":
		return Location{"Kooddoo", "Indian/Maldives"}
	case "GKL":
		return Location{"", "Australia/Brisbane"}
	case "GKN":
		return Location{"Gulkana", "America/Anchorage"}
	case "GKT":
		return Location{"Sevierville", "America/New_York"}
	case "GLA":
		return Location{"Glasgow", "Europe/London"}
	case "GLD":
		return Location{"Goodland", "America/Denver"}
	case "GLE":
		return Location{"Gainesville", "America/Chicago"}
	case "GLF":
		return Location{"Golfito", "America/Costa_Rica"}
	case "GLG":
		return Location{"", "Australia/Brisbane"}
	case "GLH":
		return Location{"Greenville", "America/Chicago"}
	case "GLI":
		return Location{"", "Australia/Sydney"}
	case "GLJ":
		return Location{"Garzon", "America/Bogota"}
	case "GLK":
		return Location{"Galcaio", "Africa/Mogadishu"}
	case "GLL":
		return Location{"Klanten", "Europe/Oslo"}
	case "GLM":
		return Location{"", "Australia/Brisbane"}
	case "GLO":
		return Location{"Staverton", "Europe/London"}
	case "GLR":
		return Location{"Gaylord", "America/Detroit"}
	case "GLS":
		return Location{"Galveston", "America/Chicago"}
	case "GLT":
		return Location{"Gladstone", "Australia/Brisbane"}
	case "GLU":
		return Location{"Gelephu", "Asia/Thimphu"}
	case "GLV":
		return Location{"Golovin", "America/Nome"}
	case "GLW":
		return Location{"Glasgow", "America/Chicago"}
	case "GLX":
		return Location{"Galela-Celebes Island", "Asia/Jayapura"}
	case "GLZ":
		return Location{"Breda", "Europe/Amsterdam"}
	case "GMA":
		return Location{"Gemena", "Africa/Kinshasa"}
	case "GMB":
		return Location{"Gambela", "Africa/Addis_Ababa"}
	case "GMD":
		return Location{"Ben Slimane", "Africa/Casablanca"}
	case "GME":
		return Location{"Gomel", "Europe/Minsk"}
	case "GMI":
		return Location{"Gasmata Island", "Pacific/Port_Moresby"}
	case "GML":
		return Location{"Kiev", "Europe/Kyiv"}
	case "GMM":
		return Location{"Gamboma", "Africa/Brazzaville"}
	case "GMN":
		return Location{"", "Pacific/Auckland"}
	case "GMO":
		return Location{"Gombe", "Africa/Lagos"}
	case "GMP":
		return Location{"Seoul", "Asia/Seoul"}
	case "GMR":
		return Location{"", "Pacific/Gambier"}
	case "GMS":
		return Location{"Uberlandia", "America/Sao_Paulo"}
	case "GMT":
		return Location{"Granite Mountain", "America/Anchorage"}
	case "GMU":
		return Location{"Greenville", "America/New_York"}
	case "GMZ":
		return Location{"Alajero", "Atlantic/Canary"}
	case "GNA":
		return Location{"Hrodna", "Europe/Minsk"}
	case "GNB":
		return Location{"Grenoble/Saint-Geoirs", "Europe/Paris"}
	case "GND":
		return Location{"Saint George's", "America/Grenada"}
	case "GNF":
		return Location{"Quincy", "America/Los_Angeles"}
	case "GNG":
		return Location{"Gooding", "America/Boise"}
	case "GNI":
		return Location{"Lyudao", "Asia/Taipei"}
	case "GNJ":
		return Location{"Ganja", "Asia/Baku"}
	case "GNM":
		return Location{"Guanambi", "America/Bahia"}
	case "GNN":
		return Location{"Ghinnir", "Africa/Addis_Ababa"}
	case "GNR":
		return Location{"General Roca", "America/Argentina/Salta"}
	case "GNS":
		return Location{"Gunung Sitoli-Nias Island", "Asia/Jakarta"}
	case "GNT":
		return Location{"Grants", "America/Denver"}
	case "GNV":
		return Location{"Gainesville", "America/New_York"}
	case "GNY":
		return Location{"Sanliurfa", "Europe/Istanbul"}
	case "GNZ":
		return Location{"Ghanzi", "Africa/Gaborone"}
	case "GOA":
		return Location{"Genova", "Europe/Rome"}
	case "GOB":
		return Location{"Goba", "Africa/Addis_Ababa"}
	case "GOG":
		return Location{"Gobabis", "Africa/Windhoek"}
	case "GOH":
		return Location{"Nuuk", "America/Nuuk"}
	case "GOI":
		return Location{"Dabolim", "Asia/Kolkata"}
	case "GOJ":
		return Location{"Nizhny Novgorod", "Europe/Moscow"}
	case "GOK":
		return Location{"Guthrie", "America/Chicago"}
	case "GOM":
		return Location{"Goma", "Africa/Kigali"}
	case "GON":
		return Location{"Groton (New London)", "America/New_York"}
	case "GOO":
		return Location{"", "Australia/Brisbane"}
	case "GOP":
		return Location{"Gorakhpur", "Asia/Kolkata"}
	case "GOQ":
		return Location{"Golmud", "Asia/Shanghai"}
	case "GOR":
		return Location{"Gore", "Africa/Addis_Ababa"}
	case "GOT":
		return Location{"Gothenburg", "Europe/Stockholm"}
	case "GOU":
		return Location{"Garoua", "Africa/Douala"}
	case "GOV":
		return Location{"Nhulunbuy", "Australia/Darwin"}
	case "GOX":
		return Location{"Mopa", "Asia/Kolkata"}
	case "GOZ":
		return Location{"Gorna Oryahovitsa", "Europe/Sofia"}
	case "GPA":
		return Location{"Patras", "Europe/Athens"}
	case "GPB":
		return Location{"Guarapuava", "America/Sao_Paulo"}
	case "GPI":
		return Location{"Guapi", "America/Bogota"}
	case "GPL":
		return Location{"Pococi", "America/Costa_Rica"}
	case "GPN":
		return Location{"", "Australia/Darwin"}
	case "GPO":
		return Location{"General Pico", "America/Argentina/Salta"}
	case "GPS":
		return Location{"Baltra", "Pacific/Galapagos"}
	case "GPT":
		return Location{"Gulfport", "America/Chicago"}
	case "GPZ":
		return Location{"Grand Rapids", "America/Chicago"}
	case "GQQ":
		return Location{"Galion", "America/New_York"}
	case "GRB":
		return Location{"Green Bay", "America/Chicago"}
	case "GRD":
		return Location{"Greenwood", "America/New_York"}
	case "GRE":
		return Location{"Greenville", "America/Chicago"}
	case "GRF":
		return Location{"Fort Lewis/Tacoma", "America/Los_Angeles"}
	case "GRI":
		return Location{"Grand Island", "America/Chicago"}
	case "GRJ":
		return Location{"George", "Africa/Johannesburg"}
	case "GRK":
		return Location{"Fort Cavazos/Killeen", "America/Chicago"}
	case "GRL":
		return Location{"Au", "Pacific/Port_Moresby"}
	case "GRM":
		return Location{"Grand Marais", "America/Chicago"}
	case "GRN":
		return Location{"Gordon", "America/Denver"}
	case "GRO":
		return Location{"Girona", "Europe/Madrid"}
	case "GRP":
		return Location{"Gurupi", "America/Araguaina"}
	case "GRQ":
		return Location{"Groningen", "Europe/Amsterdam"}
	case "GRR":
		return Location{"Grand Rapids", "America/Detroit"}
	case "GRS":
		return Location{"Grosetto", "Europe/Rome"}
	case "GRU":
		return Location{"Sao Paulo", "America/Sao_Paulo"}
	case "GRV":
		return Location{"Grozny", "Europe/Moscow"}
	case "GRW":
		return Location{"Santa Cruz da Graciosa", "Atlantic/Azores"}
	case "GRX":
		return Location{"Granada", "Europe/Madrid"}
	case "GRY":
		return Location{"Grimsey", "Atlantic/Reykjavik"}
	case "GRZ":
		return Location{"Graz", "Europe/Vienna"}
	case "GSA":
		return Location{"Long Miau", "Asia/Kuching"}
	case "GSB":
		return Location{"Goldsboro", "America/New_York"}
	case "GSC":
		return Location{"", "Australia/Perth"}
	case "GSE":
		return Location{"Gothenburg", "Europe/Stockholm"}
	case "GSH":
		return Location{"Goshen", "America/Indiana/Indianapolis"}
	case "GSI":
		return Location{"Grand-Santi", "America/Cayenne"}
	case "GSJ":
		return Location{"Puerto San Jose", "America/Guatemala"}
	case "GSM":
		return Location{"", "Asia/Tehran"}
	case "GSN":
		return Location{"Mount Gunson", "Australia/Adelaide"}
	case "GSO":
		return Location{"Greensboro", "America/New_York"}
	case "GSP":
		return Location{"Greenville", "America/New_York"}
	case "GSQ":
		return Location{"", "Africa/Cairo"}
	case "GSR":
		return Location{"Gardo", "Africa/Mogadishu"}
	case "GSS":
		return Location{"Belfast", "Africa/Johannesburg"}
	case "GST":
		return Location{"Gustavus", "America/Juneau"}
	case "GSU":
		return Location{"Gedaref", "Africa/Khartoum"}
	case "GSV":
		return Location{"Saratov", "Europe/Saratov"}
	case "GTA":
		return Location{"Gatokae", "Pacific/Guadalcanal"}
	case "GTB":
		return Location{"Fort Drum", "America/New_York"}
	case "GTE":
		return Location{"Groote Eylandt", "Australia/Darwin"}
	case "GTF":
		return Location{"Great Falls", "America/Denver"}
	case "GTG":
		return Location{"Grantsburg", "America/Chicago"}
	case "GTI":
		return Location{"Rugen", "Europe/Berlin"}
	case "GTN":
		return Location{"Glentanner Station", "Pacific/Auckland"}
	case "GTO":
		return Location{"Gorontalo-Celebes Island", "Asia/Makassar"}
	case "GTP":
		return Location{"Grants Pass", "America/Los_Angeles"}
	case "GTR":
		return Location{"Columbus/W Point/Starkville", "America/Chicago"}
	case "GTS":
		return Location{"The Granites", "Australia/Darwin"}
	case "GTT":
		return Location{"", "Australia/Brisbane"}
	case "GTW":
		return Location{"Holesov", "Europe/Prague"}
	case "GUA":
		return Location{"Guatemala City", "America/Guatemala"}
	case "GUB":
		return Location{"Guerrero Negro", "America/Mazatlan"}
	case "GUC":
		return Location{"Gunnison", "America/Denver"}
	case "GUD":
		return Location{"Goundam", "Africa/Bamako"}
	case "GUF":
		return Location{"Gulf Shores", "America/Chicago"}
	case "GUH":
		return Location{"", "Australia/Sydney"}
	case "GUI":
		return Location{"", "America/Caracas"}
	case "GUJ":
		return Location{"Guaratingueta", "America/Sao_Paulo"}
	case "GUL":
		return Location{"", "Australia/Sydney"}
	case "GUM":
		return Location{"Hagatna", "Pacific/Guam"}
	case "GUP":
		return Location{"Gallup", "America/Denver"}
	case "GUQ":
		return Location{"Guanare", "America/Caracas"}
	case "GUR":
		return Location{"Gurney", "Pacific/Port_Moresby"}
	case "GUS":
		return Location{"Peru", "America/Indiana/Indianapolis"}
	case "GUT":
		return Location{"Gutersloh", "Europe/Berlin"}
	case "GUU":
		return Location{"Grundarfjordur", "Atlantic/Reykjavik"}
	case "GUV":
		return Location{"Mougulu", "Pacific/Port_Moresby"}
	case "GUW":
		return Location{"Atyrau", "Asia/Atyrau"}
	case "GUX":
		return Location{"", "Asia/Kolkata"}
	case "GUY":
		return Location{"Guymon", "America/Chicago"}
	case "GUZ":
		return Location{"Guarapari", "America/Sao_Paulo"}
	case "GVA":
		return Location{"Geneva", "Europe/Paris"}
	case "GVE":
		return Location{"Gordonsville", "America/New_York"}
	case "GVI":
		return Location{"Green River", "Pacific/Port_Moresby"}
	case "GVL":
		return Location{"Gainesville", "America/New_York"}
	case "GVN":
		return Location{"Sovetskaya Gavan", "Asia/Vladivostok"}
	case "GVP":
		return Location{"", "Australia/Brisbane"}
	case "GVR":
		return Location{"Governador Valadares", "America/Sao_Paulo"}
	case "GVT":
		return Location{"Greenville", "America/Chicago"}
	case "GVX":
		return Location{"Gavle / Sandviken", "Europe/Stockholm"}
	case "GWA":
		return Location{"Gwa", "Asia/Yangon"}
	case "GWD":
		return Location{"Gwadar", "Asia/Karachi"}
	case "GWE":
		return Location{"Gweru", "Africa/Harare"}
	case "GWL":
		return Location{"Gwalior", "Asia/Kolkata"}
	case "GWO":
		return Location{"Greenwood", "America/Chicago"}
	case "GWS":
		return Location{"Glenwood Springs", "America/Denver"}
	case "GWT":
		return Location{"Westerland", "Europe/Berlin"}
	case "GWY":
		return Location{"Galway", "Europe/Dublin"}
	case "GXF":
		return Location{"Sayun", "Asia/Aden"}
	case "GXG":
		return Location{"Negage", "Africa/Luanda"}
	case "GXQ":
		return Location{"Coyhaique", "America/Santiago"}
	case "GXX":
		return Location{"Yagoua", "Africa/Douala"}
	case "GXY":
		return Location{"Greeley", "America/Denver"}
	case "GYA":
		return Location{"Guayaramerin", "America/La_Paz"}
	case "GYD":
		return Location{"Baku", "Asia/Baku"}
	case "GYE":
		return Location{"Guayaquil", "America/Guayaquil"}
	case "GYG":
		return Location{"Magan", "Asia/Yakutsk"}
	case "GYI":
		return Location{"Gisenyi", "Africa/Kigali"}
	case "GYL":
		return Location{"", "Australia/Perth"}
	case "GYM":
		return Location{"Guaymas", "America/Hermosillo"}
	case "GYN":
		return Location{"Goiania", "America/Sao_Paulo"}
	case "GYP":
		return Location{"", "Australia/Brisbane"}
	case "GYR":
		return Location{"Goodyear", "America/Phoenix"}
	case "GYS":
		return Location{"Guangyuan", "Asia/Shanghai"}
	case "GYU":
		return Location{"Guyuan", "Asia/Shanghai"}
	case "GYY":
		return Location{"Gary", "America/Chicago"}
	case "GYZ":
		return Location{"Gruyere", "Australia/Perth"}
	case "GZG":
		return Location{"Garze", "Asia/Shanghai"}
	case "GZO":
		return Location{"Gizo", "Pacific/Guadalcanal"}
	case "GZP":
		return Location{"Gazipasa", "Europe/Istanbul"}
	case "GZT":
		return Location{"Gaziantep", "Europe/Istanbul"}
	case "GZW":
		return Location{"Qazvin", "Asia/Tehran"}
	case "HAA":
		return Location{"Hasvik", "Europe/Oslo"}
	case "HAB":
		return Location{"Hamilton", "America/Chicago"}
	case "HAC":
		return Location{"Hachijojima", "Asia/Tokyo"}
	case "HAD":
		return Location{"Halmstad", "Europe/Stockholm"}
	case "HAF":
		return Location{"Half Moon Bay", "America/Los_Angeles"}
	case "HAH":
		return Location{"Moroni", "Indian/Comoro"}
	case "HAI":
		return Location{"Three Rivers", "America/Detroit"}
	case "HAJ":
		return Location{"Hannover", "Europe/Berlin"}
	case "HAK":
		return Location{"Haikou", "Asia/Shanghai"}
	case "HAM":
		return Location{"Hamburg", "Europe/Berlin"}
	case "HAN":
		return Location{"Hanoi", "Asia/Bangkok"}
	case "HAO":
		return Location{"Hamilton", "America/New_York"}
	case "HAQ":
		return Location{"Haa Dhaalu Atoll", "Indian/Maldives"}
	case "HAR":
		return Location{"Harrisburg", "America/New_York"}
	case "HAS":
		return Location{"", "Asia/Riyadh"}
	case "HAT":
		return Location{"", "Australia/Brisbane"}
	case "HAU":
		return Location{"Karmoy", "Europe/Oslo"}
	case "HAV":
		return Location{"Havana", "America/Havana"}
	case "HAW":
		return Location{"Haverfordwest", "Europe/London"}
	case "HAY":
		return Location{"Aguachica", "America/Bogota"}
	case "HBA":
		return Location{"Hobart", "Australia/Hobart"}
	case "HBE":
		return Location{"Alexandria", "Africa/Cairo"}
	case "HBG":
		return Location{"Hattiesburg", "America/Chicago"}
	case "HBK":
		return Location{"Holbrook", "America/Phoenix"}
	case "HBQ":
		return Location{"Haibei", "Asia/Shanghai"}
	case "HBR":
		return Location{"Hobart", "America/Chicago"}
	case "HBT":
		return Location{"King Khaled Military City", "Asia/Riyadh"}
	case "HBU":
		return Location{"", "Asia/Hovd"}
	case "HBX":
		return Location{"Hubli", "Asia/Kolkata"}
	case "HCC":
		return Location{"Hudson", "America/New_York"}
	case "HCM":
		return Location{"Eil", "Africa/Mogadishu"}
	case "HCN":
		return Location{"Hengchung", "Asia/Taipei"}
	case "HCQ":
		return Location{"", "Australia/Perth"}
	case "HCR":
		return Location{"Holy Cross", "America/Anchorage"}
	case "HCW":
		return Location{"Cheraw", "America/New_York"}
	case "HCZ":
		return Location{"Chenzhou", "Asia/Shanghai"}
	case "HDD":
		return Location{"Hyderabad", "Asia/Karachi"}
	case "HDE":
		return Location{"Holdrege", "America/Chicago"}
	case "HDF":
		return Location{"Heringsdorf", "Europe/Berlin"}
	case "HDG":
		return Location{"Handan", "Asia/Shanghai"}
	case "HDH":
		return Location{"Mokuleia", "Pacific/Honolulu"}
	case "HDK":
		return Location{"Kulhudhuffushi", "Indian/Maldives"}
	case "HDM":
		return Location{"Hamadan", "Asia/Tehran"}
	case "HDN":
		return Location{"Hayden", "America/Denver"}
	case "HDO":
		return Location{"", "Asia/Kolkata"}
	case "HDR":
		return Location{"Havadarya", "Asia/Tehran"}
	case "HDS":
		return Location{"Hoedspruit", "Africa/Johannesburg"}
	case "HDY":
		return Location{"Hat Yai", "Asia/Bangkok"}
	case "HEA":
		return Location{"", "Asia/Kabul"}
	case "HED":
		return Location{"Herendeen Bay", "America/Anchorage"}
	case "HEE":
		return Location{"Helena/West Helena", "America/Chicago"}
	case "HEG":
		return Location{"Heglig Oilfield", "Africa/Khartoum"}
	case "HEH":
		return Location{"Heho", "Asia/Yangon"}
	case "HEI":
		return Location{"Busum", "Europe/Berlin"}
	case "HEK":
		return Location{"Heihe", "Asia/Shanghai"}
	case "HEL":
		return Location{"Helsinki", "Europe/Helsinki"}
	case "HEM":
		return Location{"Helsinki", "Europe/Helsinki"}
	case "HER":
		return Location{"Heraklion", "Europe/Athens"}
	case "HES":
		return Location{"Hermiston", "America/Los_Angeles"}
	case "HET":
		return Location{"Hohhot", "Asia/Shanghai"}
	case "HEV":
		return Location{"Gibraleon", "Europe/Madrid"}
	case "HEW":
		return Location{"Athens", "Europe/Athens"}
	case "HEZ":
		return Location{"Natchez", "America/Chicago"}
	case "HFA":
		return Location{"Haifa", "Asia/Jerusalem"}
	case "HFD":
		return Location{"Hartford", "America/New_York"}
	case "HFE":
		return Location{"Hefei", "Asia/Shanghai"}
	case "HFF":
		return Location{"Camp Mackall", "America/New_York"}
	case "HFN":
		return Location{"Hornafjordur", "Atlantic/Reykjavik"}
	case "HFS":
		return Location{"", "Europe/Stockholm"}
	case "HFT":
		return Location{"Hammerfest", "Europe/Oslo"}
	case "HGA":
		return Location{"Hargeisa", "Africa/Mogadishu"}
	case "HGD":
		return Location{"", "Australia/Brisbane"}
	case "HGE":
		return Location{"", "America/Caracas"}
	case "HGH":
		return Location{"Hangzhou", "Asia/Shanghai"}
	case "HGI":
		return Location{"Itanagar", "Asia/Kolkata"}
	case "HGL":
		return Location{"Helgoland", "Europe/Berlin"}
	case "HGN":
		return Location{"", "Asia/Bangkok"}
	case "HGO":
		return Location{"", "Africa/Abidjan"}
	case "HGR":
		return Location{"Hagerstown", "America/New_York"}
	case "HGS":
		return Location{"Freetown", "Africa/Freetown"}
	case "HGU":
		return Location{"Mount Hagen", "Pacific/Port_Moresby"}
	case "HGZ":
		return Location{"Hogatza", "America/Anchorage"}
	case "HHE":
		return Location{"", "Asia/Tokyo"}
	case "HHH":
		return Location{"Hilton Head Island", "America/New_York"}
	case "HHI":
		return Location{"Wahiawa", "Pacific/Honolulu"}
	case "HHN":
		return Location{"Hahn", "Europe/Berlin"}
	case "HHQ":
		return Location{"Hua Hin", "Asia/Bangkok"}
	case "HHR":
		return Location{"Hawthorne", "America/Los_Angeles"}
	case "HHZ":
		return Location{"Hikueru Atoll", "Pacific/Tahiti"}
	case "HIA":
		return Location{"Huai'an", "Asia/Shanghai"}
	case "HIB":
		return Location{"Hibbing", "America/Chicago"}
	case "HID":
		return Location{"Horn Island", "Australia/Brisbane"}
	case "HIE":
		return Location{"Whitefield", "America/New_York"}
	case "HIF":
		return Location{"Ogden", "America/Denver"}
	case "HIG":
		return Location{"", "Australia/Brisbane"}
	case "HII":
		return Location{"Lake Havasu City", "America/Phoenix"}
	case "HIJ":
		return Location{"Hiroshima", "Asia/Tokyo"}
	case "HIM":
		return Location{"Polonnaruwa Town", "Asia/Colombo"}
	case "HIN":
		return Location{"Sacheon", "Asia/Seoul"}
	case "HIO":
		return Location{"Portland", "America/Los_Angeles"}
	case "HIP":
		return Location{"", "Australia/Brisbane"}
	case "HIR":
		return Location{"Honiara", "Pacific/Guadalcanal"}
	case "HIW":
		return Location{"", "Asia/Tokyo"}
	case "HJJ":
		return Location{"Huaihua", "Asia/Shanghai"}
	case "HJR":
		return Location{"Khajuraho", "Asia/Kolkata"}
	case "HJT":
		return Location{"", "Asia/Ulaanbaatar"}
	case "HKA":
		return Location{"Blytheville", "America/Chicago"}
	case "HKD":
		return Location{"Hakodate", "Asia/Tokyo"}
	case "HKG":
		return Location{"Hong Kong", "Asia/Hong_Kong"}
	case "HKK":
		return Location{"", "Pacific/Auckland"}
	case "HKN":
		return Location{"Hoskins", "Pacific/Port_Moresby"}
	case "HKS":
		return Location{"Jackson", "America/Chicago"}
	case "HKT":
		return Location{"Phuket", "Asia/Bangkok"}
	case "HKV":
		return Location{"Haskovo", "Europe/Sofia"}
	case "HKY":
		return Location{"Hickory", "America/New_York"}
	case "HLA":
		return Location{"Johannesburg", "Africa/Johannesburg"}
	case "HLB":
		return Location{"Batesville", "America/Indiana/Indianapolis"}
	case "HLC":
		return Location{"Hill City", "America/Chicago"}
	case "HLD":
		return Location{"Hailar", "Asia/Shanghai"}
	case "HLE":
		return Location{"Saint Helena", "Atlantic/St_Helena"}
	case "HLF":
		return Location{"", "Europe/Stockholm"}
	case "HLG":
		return Location{"Wheeling", "America/New_York"}
	case "HLH":
		return Location{"Ulanhot", "Asia/Shanghai"}
	case "HLI":
		return Location{"Hollister", "America/Los_Angeles"}
	case "HLJ":
		return Location{"Barysiai", "Europe/Vilnius"}
	case "HLL":
		return Location{"", "Australia/Perth"}
	case "HLM":
		return Location{"Holland", "America/Detroit"}
	case "HLN":
		return Location{"Helena", "America/Denver"}
	case "HLO":
		return Location{"Onundarfjordur", "Atlantic/Reykjavik"}
	case "HLP":
		return Location{"Jakarta", "Asia/Jakarta"}
	case "HLR":
		return Location{"Fort Cavazos(Killeen)", "America/Chicago"}
	case "HLS":
		return Location{"", "Australia/Hobart"}
	case "HLT":
		return Location{"", "Australia/Melbourne"}
	case "HLU":
		return Location{"Houailou", "Pacific/Noumea"}
	case "HLW":
		return Location{"Hluhluwe", "Africa/Johannesburg"}
	case "HLY":
		return Location{"Angelsey", "Europe/London"}
	case "HLZ":
		return Location{"Hamilton", "Pacific/Auckland"}
	case "HMA":
		return Location{"Khanty-Mansiysk", "Asia/Yekaterinburg"}
	case "HMB":
		return Location{"Sohag", "Africa/Cairo"}
	case "HME":
		return Location{"Hassi Messaoud", "Africa/Algiers"}
	case "HMG":
		return Location{"", "Australia/Darwin"}
	case "HMI":
		return Location{"Hami", "Asia/Shanghai"}
	case "HMJ":
		return Location{"Khmelnytskyi", "Europe/Kyiv"}
	case "HMN":
		return Location{"Alamogordo", "America/Denver"}
	case "HMO":
		return Location{"Hermosillo", "America/Hermosillo"}
	case "HMR":
		return Location{"Hamar", "Europe/Oslo"}
	case "HMT":
		return Location{"Hemet", "America/Los_Angeles"}
	case "HMV":
		return Location{"", "Europe/Stockholm"}
	case "HMY":
		return Location{"Seosan", "Asia/Seoul"}
	case "HNA":
		return Location{"", "Asia/Tokyo"}
	case "HNB":
		return Location{"Huntingburg", "America/Indiana/Vincennes"}
	case "HNC":
		return Location{"Hatteras", "America/New_York"}
	case "HND":
		return Location{"Tokyo", "Asia/Tokyo"}
	case "HNE":
		return Location{"Tahneta Pass Lodge", "America/Anchorage"}
	case "HNH":
		return Location{"Hoonah", "America/Juneau"}
	case "HNI":
		return Location{"Hechi", "Asia/Shanghai"}
	case "HNL":
		return Location{"Honolulu", "Pacific/Honolulu"}
	case "HNM":
		return Location{"Hana", "Pacific/Honolulu"}
	case "HNS":
		return Location{"Haines", "America/Juneau"}
	case "HNY":
		return Location{"Hengyang", "Asia/Shanghai"}
	case "HOA":
		return Location{"Hola", "Africa/Nairobi"}
	case "HOB":
		return Location{"Hobbs", "America/Denver"}
	case "HOD":
		return Location{"Hodeida", "Asia/Aden"}
	case "HOF":
		return Location{"", "Asia/Riyadh"}
	case "HOG":
		return Location{"Holguin", "America/Havana"}
	case "HOH":
		return Location{"Hohenems / Dornbirn", "Europe/Vienna"}
	case "HOI":
		return Location{"", "Pacific/Tahiti"}
	case "HOK":
		return Location{"", "Australia/Darwin"}
	case "HOM":
		return Location{"Homer", "America/Anchorage"}
	case "HON":
		return Location{"Huron", "America/Chicago"}
	case "HOP":
		return Location{"Fort Campbell/Hopkinsville", "America/Chicago"}
	case "HOQ":
		return Location{"Hof", "Europe/Berlin"}
	case "HOR":
		return Location{"Horta", "Atlantic/Azores"}
	case "HOS":
		return Location{"Chos Malal", "America/Argentina/Salta"}
	case "HOT":
		return Location{"Hot Springs", "America/Chicago"}
	case "HOU":
		return Location{"Houston", "America/Chicago"}
	case "HOV":
		return Location{"Orsta", "Europe/Oslo"}
	case "HOX":
		return Location{"Hommalinn", "Asia/Yangon"}
	case "HPA":
		return Location{"Lifuka", "Pacific/Tongatapu"}
	case "HPB":
		return Location{"Hooper Bay", "America/Nome"}
	case "HPH":
		return Location{"Haiphong", "Asia/Bangkok"}
	case "HPN":
		return Location{"White Plains", "America/New_York"}
	case "HPT":
		return Location{"Hampton", "America/Chicago"}
	case "HPY":
		return Location{"Baytown", "America/Chicago"}
	case "HQL":
		return Location{"Tashkurgan", "Asia/Shanghai"}
	case "HQM":
		return Location{"Hoquiam", "America/Los_Angeles"}
	case "HRB":
		return Location{"Harbin", "Asia/Shanghai"}
	case "HRE":
		return Location{"Harare", "Africa/Harare"}
	case "HRF":
		return Location{"Hoarafushi", "Indian/Maldives"}
	case "HRG":
		return Location{"Hurghada", "Africa/Cairo"}
	case "HRH":
		return Location{"Aligarth", "Asia/Kolkata"}
	case "HRI":
		return Location{"Mattala", "Asia/Colombo"}
	case "HRK":
		return Location{"Kharkiv", "Europe/Kyiv"}
	case "HRL":
		return Location{"Harlingen", "America/Chicago"}
	case "HRM":
		return Location{"", "Africa/Algiers"}
	case "HRO":
		return Location{"Harrison", "America/Chicago"}
	case "HRS":
		return Location{"Harrismith", "Africa/Johannesburg"}
	case "HRT":
		return Location{"Linton-On-Ouse", "Europe/London"}
	case "HRY":
		return Location{"", "Australia/Darwin"}
	case "HRZ":
		return Location{"Horizontina", "America/Sao_Paulo"}
	case "HSA":
		return Location{"Turkistan", "Asia/Almaty"}
	case "HSB":
		return Location{"Harrisburg", "America/Chicago"}
	case "HSC":
		return Location{"Shaoguan", "Asia/Shanghai"}
	case "HSG":
		return Location{"Saga", "Asia/Tokyo"}
	case "HSH":
		return Location{"Las Vegas", "America/Los_Angeles"}
	case "HSI":
		return Location{"Hastings", "America/Chicago"}
	case "HSK":
		return Location{"Monflorite/Alcala del Obispo", "Europe/Madrid"}
	case "HSL":
		return Location{"Huslia", "America/Anchorage"}
	case "HSM":
		return Location{"", "Australia/Melbourne"}
	case "HSN":
		return Location{"Zhoushan", "Asia/Shanghai"}
	case "HSP":
		return Location{"Hot Springs", "America/New_York"}
	case "HSR":
		return Location{"Rajkot", "Asia/Kolkata"}
	case "HSS":
		return Location{"", "Asia/Kolkata"}
	case "HST":
		return Location{"Homestead", "America/New_York"}
	case "HSV":
		return Location{"Huntsville", "America/Chicago"}
	case "HSZ":
		return Location{"Hsinchu City", "Asia/Taipei"}
	case "HTA":
		return Location{"Chita", "Asia/Chita"}
	case "HTG":
		return Location{"Khatanga", "Asia/Krasnoyarsk"}
	case "HTH":
		return Location{"Hawthorne", "America/Los_Angeles"}
	case "HTI":
		return Location{"Hamilton Island", "Australia/Lindeman"}
	case "HTL":
		return Location{"Houghton Lake", "America/Detroit"}
	case "HTN":
		return Location{"Hotan", "Asia/Shanghai"}
	case "HTO":
		return Location{"East Hampton", "America/New_York"}
	case "HTR":
		return Location{"Hateruma", "Asia/Tokyo"}
	case "HTS":
		return Location{"Huntington", "America/New_York"}
	case "HTU":
		return Location{"", "Australia/Melbourne"}
	case "HTV":
		return Location{"Huntsville", "America/Chicago"}
	case "HTW":
		return Location{"Chesapeake/Huntington Wva", "America/New_York"}
	case "HTY":
		return Location{"Hatay", "Europe/Istanbul"}
	case "HTZ":
		return Location{"Hato Corozal", "America/Bogota"}
	case "HUA":
		return Location{"Redstone Arsnl Huntsville", "America/Chicago"}
	case "HUB":
		return Location{"", "Australia/Darwin"}
	case "HUE":
		return Location{"Humera", "Africa/Addis_Ababa"}
	case "HUF":
		return Location{"Terre Haute", "America/Indiana/Indianapolis"}
	case "HUG":
		return Location{"Huehuetenango", "America/Guatemala"}
	case "HUH":
		return Location{"Fare", "Pacific/Tahiti"}
	case "HUI":
		return Location{"Hue", "Asia/Ho_Chi_Minh"}
	case "HUJ":
		return Location{"Hugo", "America/Chicago"}
	case "HUL":
		return Location{"Houlton", "America/Moncton"}
	case "HUM":
		return Location{"Houma", "America/Chicago"}
	case "HUN":
		return Location{"Hualien City", "Asia/Taipei"}
	case "HUO":
		return Location{"Holingol", "Asia/Shanghai"}
	case "HUQ":
		return Location{"", "Africa/Tripoli"}
	case "HUS":
		return Location{"Hughes", "America/Anchorage"}
	case "HUT":
		return Location{"Hutchinson", "America/Chicago"}
	case "HUU":
		return Location{"Huanuco", "America/Lima"}
	case "HUW":
		return Location{"Humaita", "America/Manaus"}
	case "HUX":
		return Location{"Huatulco", "America/Mexico_City"}
	case "HUY":
		return Location{"Grimsby", "Europe/London"}
	case "HUZ":
		return Location{"Huizhou", "Asia/Shanghai"}
	case "HVA":
		return Location{"", "Indian/Antananarivo"}
	case "HVB":
		return Location{"Hervey Bay", "Australia/Brisbane"}
	case "HVD":
		return Location{"Khovd", "Asia/Hovd"}
	case "HVE":
		return Location{"Hanksville", "America/Denver"}
	case "HVG":
		return Location{"Honningsvag", "Europe/Oslo"}
	case "HVK":
		return Location{"Holmavik", "Atlantic/Reykjavik"}
	case "HVM":
		return Location{"Hvammstangi", "Atlantic/Reykjavik"}
	case "HVN":
		return Location{"New Haven", "America/New_York"}
	case "HVR":
		return Location{"Havre", "America/Denver"}
	case "HVS":
		return Location{"Hartsville", "America/New_York"}
	case "HWD":
		return Location{"Hayward", "America/Los_Angeles"}
	case "HWK":
		return Location{"Hawker", "Australia/Adelaide"}
	case "HWN":
		return Location{"Hwange", "Africa/Harare"}
	case "HWO":
		return Location{"Hollywood", "America/New_York"}
	case "HXD":
		return Location{"Delingha", "Asia/Shanghai"}
	case "HXX":
		return Location{"", "Australia/Sydney"}
	case "HYA":
		return Location{"Hyannis", "America/New_York"}
	case "HYC":
		return Location{"High Wycombe", "Europe/London"}
	case "HYD":
		return Location{"Hyderabad", "Asia/Kolkata"}
	case "HYN":
		return Location{"Huangyan", "Asia/Shanghai"}
	case "HYR":
		return Location{"Hayward", "America/Chicago"}
	case "HYS":
		return Location{"Hays", "America/Chicago"}
	case "HYV":
		return Location{"", "Europe/Helsinki"}
	case "HZB":
		return Location{"Merville/Calonne", "Europe/Paris"}
	case "HZG":
		return Location{"Hanzhong", "Asia/Shanghai"}
	case "HZH":
		return Location{"", "Asia/Shanghai"}
	case "HZK":
		return Location{"Husavik", "Atlantic/Reykjavik"}
	case "HZL":
		return Location{"Hazleton", "America/New_York"}
	case "HZP":
		return Location{"Fort Mackay", "America/Edmonton"}
	case "IAA":
		return Location{"Igarka", "Asia/Krasnoyarsk"}
	case "IAB":
		return Location{"Wichita", "America/Chicago"}
	case "IAD":
		return Location{"Dulles", "America/New_York"}
	case "IAG":
		return Location{"Niagara Falls", "America/New_York"}
	case "IAH":
		return Location{"Houston", "America/Chicago"}
	case "IAM":
		return Location{"Amenas", "Africa/Algiers"}
	case "IAN":
		return Location{"Kiana", "America/Anchorage"}
	case "IAO":
		return Location{"Del Carmen", "Asia/Manila"}
	case "IAQ":
		return Location{"", "Asia/Tehran"}
	case "IAR":
		return Location{"", "Europe/Moscow"}
	case "IAS":
		return Location{"Iasi", "Europe/Bucharest"}
	case "IBA":
		return Location{"Ibadan", "Africa/Lagos"}
	case "IBB":
		return Location{"Isabela", "Pacific/Galapagos"}
	case "IBE":
		return Location{"Ibague", "America/Bogota"}
	case "IBP":
		return Location{"Iberia", "America/Lima"}
	case "IBR":
		return Location{"Omitama", "Asia/Tokyo"}
	case "IBZ":
		return Location{"Ibiza", "Europe/Madrid"}
	case "ICA":
		return Location{"", "America/Caracas"}
	case "ICC":
		return Location{"Isla de Coche", "America/Caracas"}
	case "ICI":
		return Location{"Cicia", "Pacific/Fiji"}
	case "ICK":
		return Location{"Nieuw Nickerie", "America/Paramaribo"}
	case "ICL":
		return Location{"Clarinda", "America/Chicago"}
	case "ICN":
		return Location{"Seoul", "Asia/Seoul"}
	case "ICR":
		return Location{"Nicaro", "America/Havana"}
	case "ICS":
		return Location{"Cascade", "America/Boise"}
	case "ICT":
		return Location{"Wichita", "America/Chicago"}
	case "ICY":
		return Location{"Icy Bay", "America/Anchorage"}
	case "IDA":
		return Location{"Idaho Falls", "America/Boise"}
	case "IDB":
		return Location{"Idre", "Europe/Stockholm"}
	case "IDF":
		return Location{"Idiofa", "Africa/Kinshasa"}
	case "IDG":
		return Location{"Ida Grove", "America/Chicago"}
	case "IDH":
		return Location{"Grangeville", "America/Los_Angeles"}
	case "IDI":
		return Location{"Indiana", "America/New_York"}
	case "IDK":
		return Location{"", "Australia/Adelaide"}
	case "IDO":
		return Location{"Cristalandia", "America/Araguaina"}
	case "IDP":
		return Location{"Independence", "America/Chicago"}
	case "IDR":
		return Location{"Indore", "Asia/Kolkata"}
	case "IDY":
		return Location{"Ile d'Yeu", "Europe/Paris"}
	case "IEG":
		return Location{"Babimost", "Europe/Warsaw"}
	case "IEJ":
		return Location{"", "Asia/Tokyo"}
	case "IES":
		return Location{"Riesa", "Europe/Berlin"}
	case "IEV":
		return Location{"Kiev", "Europe/Kyiv"}
	case "IFA":
		return Location{"Iowa Falls", "America/Chicago"}
	case "IFF":
		return Location{"", "Australia/Brisbane"}
	case "IFH":
		return Location{"Hesa", "Asia/Tehran"}
	case "IFJ":
		return Location{"Isafjordur", "Atlantic/Reykjavik"}
	case "IFL":
		return Location{"", "Australia/Brisbane"}
	case "IFN":
		return Location{"Isfahan", "Asia/Tehran"}
	case "IFO":
		return Location{"Ivano-Frankivsk", "Europe/Kyiv"}
	case "IFP":
		return Location{"Bullhead City", "America/Phoenix"}
	case "IFU":
		return Location{"Ifuru", "Indian/Maldives"}
	case "IGA":
		return Location{"Matthew Town", "America/Nassau"}
	case "IGB":
		return Location{"Ingeniero Jacobacci", "America/Argentina/Salta"}
	case "IGD":
		return Location{"Igdir", "Europe/Istanbul"}
	case "IGG":
		return Location{"Igiugig", "America/Anchorage"}
	case "IGH":
		return Location{"", "Australia/Brisbane"}
	case "IGL":
		return Location{"Izmir", "Europe/Istanbul"}
	case "IGM":
		return Location{"Kingman", "America/Phoenix"}
	case "IGN":
		return Location{"", "Asia/Manila"}
	case "IGO":
		return Location{"Chigorodo", "America/Bogota"}
	case "IGR":
		return Location{"Puerto Iguazu", "America/Argentina/Cordoba"}
	case "IGS":
		return Location{"Manching", "Europe/Berlin"}
	case "IGT":
		return Location{"Magas", "Europe/Moscow"}
	case "IGU":
		return Location{"Foz Do Iguacu", "America/Argentina/Cordoba"}
	case "IHC":
		return Location{"Inhaca", "Africa/Maputo"}
	case "IHN":
		return Location{"Qishn", "Asia/Aden"}
	case "IHO":
		return Location{"Ihosy", "Indian/Antananarivo"}
	case "IHR":
		return Location{"", "Asia/Tehran"}
	case "IIA":
		return Location{"Inis Meain", "Europe/Dublin"}
	case "IIL":
		return Location{"Ilam", "Asia/Tehran"}
	case "IJK":
		return Location{"Izhevsk", "Europe/Samara"}
	case "IJU":
		return Location{"Ijui", "America/Sao_Paulo"}
	case "IJX":
		return Location{"Jacksonville", "America/Chicago"}
	case "IKA":
		return Location{"Tehran", "Asia/Tehran"}
	case "IKB":
		return Location{"North Wilkesboro", "America/New_York"}
	case "IKI":
		return Location{"Iki", "Asia/Tokyo"}
	case "IKK":
		return Location{"Kankakee", "America/Chicago"}
	case "IKL":
		return Location{"Ikela", "Africa/Kinshasa"}
	case "IKO":
		return Location{"Nikolski", "America/Nome"}
	case "IKP":
		return Location{"", "Australia/Brisbane"}
	case "IKS":
		return Location{"Tiksi", "Asia/Yakutsk"}
	case "IKT":
		return Location{"Irkutsk", "Asia/Irkutsk"}
	case "IKU":
		return Location{"Tamchy", "Asia/Bishkek"}
	case "ILA":
		return Location{"Illaga-Papua Island", "Asia/Jayapura"}
	case "ILD":
		return Location{"Lleida", "Europe/Madrid"}
	case "ILE":
		return Location{"Killeen", "America/Chicago"}
	case "ILF":
		return Location{"Ilford", "America/Winnipeg"}
	case "ILG":
		return Location{"Wilmington", "America/New_York"}
	case "ILH":
		return Location{"", "Europe/Berlin"}
	case "ILI":
		return Location{"Iliamna", "America/Anchorage"}
	case "ILK":
		return Location{"Ilaka", "Indian/Antananarivo"}
	case "ILL":
		return Location{"Willmar", "America/Chicago"}
	case "ILM":
		return Location{"Wilmington", "America/New_York"}
	case "ILN":
		return Location{"Wilmington", "America/New_York"}
	case "ILO":
		return Location{"Iloilo City", "Asia/Manila"}
	case "ILP":
		return Location{"Ile des Pins", "Pacific/Noumea"}
	case "ILQ":
		return Location{"Ilo", "America/Lima"}
	case "ILR":
		return Location{"Ilorin", "Africa/Lagos"}
	case "ILS":
		return Location{"San Salvador", "America/El_Salvador"}
	case "ILU":
		return Location{"Kilaguni", "Africa/Nairobi"}
	case "ILY":
		return Location{"Port Ellen", "Europe/London"}
	case "ILZ":
		return Location{"Zilina", "Europe/Bratislava"}
	case "IMB":
		return Location{"Imbaimadai", "America/Guyana"}
	case "IMF":
		return Location{"Imphal", "Asia/Kolkata"}
	case "IMK":
		return Location{"Simikot", "Asia/Kathmandu"}
	case "IML":
		return Location{"Imperial", "America/Denver"}
	case "IMM":
		return Location{"Immokalee", "America/New_York"}
	case "IMO":
		return Location{"Zemio", "Africa/Bangui"}
	case "IMP":
		return Location{"Imperatriz", "America/Fortaleza"}
	case "IMQ":
		return Location{"Makou", "Asia/Tehran"}
	case "IMT":
		return Location{"Iron Mountain Kingsford", "America/Chicago"}
	case "INA":
		return Location{"Inta", "Europe/Moscow"}
	case "INC":
		return Location{"Yinchuan", "Asia/Shanghai"}
	case "IND":
		return Location{"Indianapolis", "America/Indiana/Indianapolis"}
	case "INF":
		return Location{"In Guezzam", "Africa/Algiers"}
	case "ING":
		return Location{"El Calafate", "America/Argentina/Rio_Gallegos"}
	case "INH":
		return Location{"Inhambabe", "Africa/Maputo"}
	case "INI":
		return Location{"Nis", "Europe/Belgrade"}
	case "INJ":
		return Location{"", "Australia/Brisbane"}
	case "INK":
		return Location{"Wink", "America/Chicago"}
	case "INL":
		return Location{"International Falls", "America/Chicago"}
	case "INM":
		return Location{"", "Australia/Adelaide"}
	case "INN":
		return Location{"Innsbruck", "Europe/Vienna"}
	case "INO":
		return Location{"Inongo", "Africa/Kinshasa"}
	case "INQ":
		return Location{"Inis Oirr", "Europe/Dublin"}
	case "INS":
		return Location{"Indian Springs", "America/Los_Angeles"}
	case "INT":
		return Location{"Winston Salem", "America/New_York"}
	case "INU":
		return Location{"Yaren District", "Pacific/Nauru"}
	case "INV":
		return Location{"Inverness", "Europe/London"}
	case "INW":
		return Location{"Winslow", "America/Phoenix"}
	case "INX":
		return Location{"Inanwatan Airport-Papua Island", "Asia/Jayapura"}
	case "INZ":
		return Location{"In Salah", "Africa/Algiers"}
	case "IOA":
		return Location{"Ioannina", "Europe/Athens"}
	case "IOM":
		return Location{"Castletown", "Europe/Isle_of_Man"}
	case "ION":
		return Location{"Impfondo", "Africa/Brazzaville"}
	case "IOR":
		return Location{"Inis Mor", "Europe/Dublin"}
	case "IOS":
		return Location{"Ilheus", "America/Bahia"}
	case "IOU":
		return Location{"Ile Ouen", "Pacific/Noumea"}
	case "IOW":
		return Location{"Iowa City", "America/Chicago"}
	case "IPA":
		return Location{"Ipota", "Pacific/Efate"}
	case "IPC":
		return Location{"Isla De Pascua", "Pacific/Easter"}
	case "IPE":
		return Location{"Ipil", "Asia/Manila"}
	case "IPG":
		return Location{"Santo Antonio Do Ica", "America/Bogota"}
	case "IPH":
		return Location{"Ipoh", "Asia/Kuala_Lumpur"}
	case "IPI":
		return Location{"Ipiales", "America/Bogota"}
	case "IPL":
		return Location{"Imperial", "America/Los_Angeles"}
	case "IPN":
		return Location{"Ipatinga", "America/Sao_Paulo"}
	case "IPT":
		return Location{"Williamsport", "America/New_York"}
	case "IPU":
		return Location{"Ipiau", "America/Bahia"}
	case "IPZ":
		return Location{"Perez Zeledon", "America/Costa_Rica"}
	case "IQA":
		return Location{"Hit", "Asia/Baghdad"}
	case "IQM":
		return Location{"Qiemo", "Asia/Shanghai"}
	case "IQN":
		return Location{"Qingyang", "Asia/Shanghai"}
	case "IQQ":
		return Location{"Iquique", "America/Santiago"}
	case "IQT":
		return Location{"Iquitos", "America/Lima"}
	case "IRA":
		return Location{"Kirakira", "Pacific/Guadalcanal"}
	case "IRC":
		return Location{"Circle", "America/Anchorage"}
	case "IRD":
		return Location{"Ishurdi", "Asia/Dhaka"}
	case "IRE":
		return Location{"Irece", "America/Bahia"}
	case "IRG":
		return Location{"", "Australia/Brisbane"}
	case "IRI":
		return Location{"Nduli", "Africa/Dar_es_Salaam"}
	case "IRJ":
		return Location{"La Rioja", "America/Argentina/La_Rioja"}
	case "IRK":
		return Location{"Kirksville", "America/Chicago"}
	case "IRM":
		return Location{"", "Asia/Yekaterinburg"}
	case "IRN":
		return Location{"Iriona", "America/Tegucigalpa"}
	case "IRO":
		return Location{"Birao", "Africa/Bangui"}
	case "IRP":
		return Location{"", "Africa/Lubumbashi"}
	case "IRS":
		return Location{"Sturgis", "America/Detroit"}
	case "IRZ":
		return Location{"Santa Isabel Do Rio Negro", "America/Manaus"}
	case "ISA":
		return Location{"Mount Isa", "Australia/Brisbane"}
	case "ISB":
		return Location{"Islamabad", "Asia/Karachi"}
	case "ISC":
		return Location{"St. Mary's", "Europe/London"}
	case "ISE":
		return Location{"Isparta", "Europe/Istanbul"}
	case "ISG":
		return Location{"Ishigaki", "Asia/Tokyo"}
	case "ISI":
		return Location{"", "Australia/Brisbane"}
	case "ISJ":
		return Location{"", "America/Cancun"}
	case "ISK":
		return Location{"Nasik", "Asia/Kolkata"}
	case "ISL":
		return Location{"Istanbul", "Europe/Istanbul"}
	case "ISM":
		return Location{"Orlando", "America/New_York"}
	case "ISN":
		return Location{"Williston", "America/Chicago"}
	case "ISO":
		return Location{"Kinston", "America/New_York"}
	case "ISP":
		return Location{"Islip", "America/New_York"}
	case "ISQ":
		return Location{"Manistique", "America/Detroit"}
	case "ISS":
		return Location{"Wiscasset", "America/New_York"}
	case "IST":
		return Location{"Arnavutkoy", "Europe/Istanbul"}
	case "ISU":
		return Location{"Sulaymaniyah", "Asia/Baghdad"}
	case "ISW":
		return Location{"Wisconsin Rapids", "America/Chicago"}
	case "ITA":
		return Location{"Itacoatiara", "America/Manaus"}
	case "ITB":
		return Location{"Itaituba", "America/Santarem"}
	case "ITE":
		return Location{"Itubera", "America/Bahia"}
	case "ITH":
		return Location{"Ithaca", "America/New_York"}
	case "ITI":
		return Location{"Cumaru Do Norte", "America/Belem"}
	case "ITM":
		return Location{"Osaka", "Asia/Tokyo"}
	case "ITO":
		return Location{"Hilo", "Pacific/Honolulu"}
	case "ITP":
		return Location{"Itaperuna", "America/Sao_Paulo"}
	case "ITQ":
		return Location{"Itaqui", "America/Sao_Paulo"}
	case "IUD":
		return Location{"Ar Rayyan", "Asia/Qatar"}
	case "IUE":
		return Location{"Alofi", "Pacific/Niue"}
	case "IVA":
		return Location{"Ambanja", "Indian/Antananarivo"}
	case "IVC":
		return Location{"Invercargill", "Pacific/Auckland"}
	case "IVG":
		return Location{"Berane", "Europe/Podgorica"}
	case "IVL":
		return Location{"Ivalo", "Europe/Helsinki"}
	case "IVR":
		return Location{"", "Australia/Sydney"}
	case "IVW":
		return Location{"Inverway", "Australia/Darwin"}
	case "IWA":
		return Location{"Ivanovo", "Europe/Moscow"}
	case "IWD":
		return Location{"Ironwood", "America/Menominee"}
	case "IWJ":
		return Location{"Masuda", "Asia/Tokyo"}
	case "IWK":
		return Location{"Iwakuni", "Asia/Tokyo"}
	case "IWO":
		return Location{"", "Asia/Tokyo"}
	case "IWS":
		return Location{"Houston", "America/Chicago"}
	case "IXA":
		return Location{"Agartala", "Asia/Kolkata"}
	case "IXB":
		return Location{"Siliguri", "Asia/Kolkata"}
	case "IXC":
		return Location{"Chandigarh", "Asia/Kolkata"}
	case "IXD":
		return Location{"Allahabad", "Asia/Kolkata"}
	case "IXE":
		return Location{"Mangalore", "Asia/Kolkata"}
	case "IXG":
		return Location{"", "Asia/Kolkata"}
	case "IXH":
		return Location{"", "Asia/Dhaka"}
	case "IXI":
		return Location{"Lilabari", "Asia/Kolkata"}
	case "IXJ":
		return Location{"Jammu", "Asia/Kolkata"}
	case "IXK":
		return Location{"", "Asia/Kolkata"}
	case "IXL":
		return Location{"Leh", "Asia/Kolkata"}
	case "IXM":
		return Location{"Madurai", "Asia/Kolkata"}
	case "IXN":
		return Location{"Khowai", "Asia/Dhaka"}
	case "IXP":
		return Location{"", "Asia/Kolkata"}
	case "IXQ":
		return Location{"", "Asia/Kolkata"}
	case "IXR":
		return Location{"Ranchi", "Asia/Kolkata"}
	case "IXS":
		return Location{"Silchar", "Asia/Kolkata"}
	case "IXT":
		return Location{"Pasighat", "Asia/Kolkata"}
	case "IXU":
		return Location{"Aurangabad", "Asia/Kolkata"}
	case "IXV":
		return Location{"", "Asia/Kolkata"}
	case "IXW":
		return Location{"", "Asia/Kolkata"}
	case "IXY":
		return Location{"Kandla", "Asia/Kolkata"}
	case "IXZ":
		return Location{"Port Blair", "Asia/Kolkata"}
	case "IYK":
		return Location{"Inyokern", "America/Los_Angeles"}
	case "IYS":
		return Location{"Wasilla", "America/Anchorage"}
	case "IZA":
		return Location{"Juiz De Fora", "America/Sao_Paulo"}
	case "IZO":
		return Location{"Izumo", "Asia/Tokyo"}
	case "IZT":
		return Location{"", "America/Mexico_City"}
	case "JAA":
		return Location{"", "Asia/Kabul"}
	case "JAB":
		return Location{"", "Australia/Darwin"}
	case "JAC":
		return Location{"Jackson", "America/Denver"}
	case "JAD":
		return Location{"Perth", "Australia/Perth"}
	case "JAE":
		return Location{"Jaen", "America/Lima"}
	case "JAF":
		return Location{"Jaffna", "Asia/Colombo"}
	case "JAG":
		return Location{"Jacobabad", "Asia/Karachi"}
	case "JAI":
		return Location{"Jaipur", "Asia/Kolkata"}
	case "JAK":
		return Location{"Jacmel", "America/Port-au-Prince"}
	case "JAL":
		return Location{"Xalapa", "America/Mexico_City"}
	case "JAM":
		return Location{"Yambol", "Europe/Sofia"}
	case "JAN":
		return Location{"Jackson", "America/Chicago"}
	case "JAP":
		return Location{"Puntarenas", "America/Costa_Rica"}
	case "JAR":
		return Location{"", "Asia/Tehran"}
	case "JAS":
		return Location{"Jasper", "America/Chicago"}
	case "JAU":
		return Location{"Jauja", "America/Lima"}
	case "JAV":
		return Location{"Ilulissat", "America/Nuuk"}
	case "JAW":
		return Location{"Araripina", "America/Recife"}
	case "JAX":
		return Location{"Jacksonville", "America/New_York"}
	case "JBQ":
		return Location{"La Isabela", "America/Santo_Domingo"}
	case "JBR":
		return Location{"Jonesboro", "America/Chicago"}
	case "JCB":
		return Location{"Joacaba", "America/Sao_Paulo"}
	case "JCI":
		return Location{"Olathe", "America/Chicago"}
	case "JCK":
		return Location{"", "Australia/Brisbane"}
	case "JCM":
		return Location{"Jacobina", "America/Bahia"}
	case "JCR":
		return Location{"Jacareacanga", "America/Santarem"}
	case "JCT":
		return Location{"Junction", "America/Chicago"}
	case "JDA":
		return Location{"John Day", "America/Los_Angeles"}
	case "JDF":
		return Location{"Juiz De Fora", "America/Sao_Paulo"}
	case "JDG":
		return Location{"", "Asia/Seoul"}
	case "JDH":
		return Location{"Jodhpur", "Asia/Kolkata"}
	case "JDN":
		return Location{"Jordan", "America/Denver"}
	case "JDO":
		return Location{"Juazeiro Do Norte", "America/Fortaleza"}
	case "JDR":
		return Location{"Sao Joao Del Rei", "America/Sao_Paulo"}
	case "JDZ":
		return Location{"Jingdezhen", "Asia/Shanghai"}
	case "JED":
		return Location{"Jeddah", "Asia/Riyadh"}
	case "JEE":
		return Location{"Jeremie", "America/Port-au-Prince"}
	case "JEF":
		return Location{"Jefferson City", "America/Chicago"}
	case "JEG":
		return Location{"Aasiaat", "America/Nuuk"}
	case "JEK":
		return Location{"Lower Zambezi National Park", "Africa/Harare"}
	case "JEQ":
		return Location{"Jequie", "America/Bahia"}
	case "JER":
		return Location{"Saint Helier", "Europe/Jersey"}
	case "JFK":
		return Location{"New York", "America/New_York"}
	case "JFN":
		return Location{"Ashtabula", "America/New_York"}
	case "JFR":
		return Location{"Paamiut", "America/Nuuk"}
	case "JGA":
		return Location{"Jamnagar", "Asia/Kolkata"}
	case "JGN":
		return Location{"Jiayuguan", "Asia/Shanghai"}
	case "JGS":
		return Location{"Ji'an", "Asia/Shanghai"}
	case "JHB":
		return Location{"Senai", "Asia/Kuala_Lumpur"}
	case "JHF":
		return Location{"São Roque", "America/Sao_Paulo"}
	case "JHG":
		return Location{"Jinghong", "Asia/Shanghai"}
	case "JHL":
		return Location{"Albian Village", "America/Edmonton"}
	case "JHM":
		return Location{"Lahaina", "Pacific/Honolulu"}
	case "JHQ":
		return Location{"", "Australia/Brisbane"}
	case "JHS":
		return Location{"Sisimiut", "America/Nuuk"}
	case "JHW":
		return Location{"Jamestown", "America/New_York"}
	case "JIA":
		return Location{"Juina", "America/Cuiaba"}
	case "JIB":
		return Location{"Djibouti City", "Africa/Djibouti"}
	case "JIC":
		return Location{"Jinchang", "Asia/Shanghai"}
	case "JIJ":
		return Location{"Jijiga", "Africa/Addis_Ababa"}
	case "JIK":
		return Location{"Ikaria Island", "Europe/Athens"}
	case "JIL":
		return Location{"Jilin", "Asia/Shanghai"}
	case "JIM":
		return Location{"Jimma", "Africa/Addis_Ababa"}
	case "JIN":
		return Location{"Jinja", "Africa/Kampala"}
	case "JIP":
		return Location{"Jipijapa", "America/Guayaquil"}
	case "JIQ":
		return Location{"Chongqing", "Asia/Shanghai"}
	case "JIR":
		return Location{"Jiri", "Asia/Kathmandu"}
	case "JIU":
		return Location{"Jiujiang", "Asia/Shanghai"}
	case "JIW":
		return Location{"Jiwani", "Asia/Karachi"}
	case "JJD":
		return Location{"Jijoca de Jericoacoara", "America/Fortaleza"}
	case "JJI":
		return Location{"Juanjui", "America/Lima"}
	case "JJM":
		return Location{"Meru-Kinna", "Africa/Nairobi"}
	case "JJN":
		return Location{"Quanzhou", "Asia/Shanghai"}
	case "JKG":
		return Location{"Jonkoping", "Europe/Stockholm"}
	case "JKH":
		return Location{"Chios Island", "Europe/Athens"}
	case "JKL":
		return Location{"Kalymnos Island", "Europe/Athens"}
	case "JKR":
		return Location{"Janakpur", "Asia/Kathmandu"}
	case "JKV":
		return Location{"Jacksonville", "America/Chicago"}
	case "JLD":
		return Location{"Landskrona", "Europe/Stockholm"}
	case "JLN":
		return Location{"Joplin", "America/Chicago"}
	case "JLR":
		return Location{"", "Asia/Kolkata"}
	case "JLS":
		return Location{"Jales", "America/Sao_Paulo"}
	case "JMK":
		return Location{"Mykonos Island", "Europe/Athens"}
	case "JMO":
		return Location{"Jomsom", "Asia/Kathmandu"}
	case "JMS":
		return Location{"Jamestown", "America/Chicago"}
	case "JMU":
		return Location{"Jiamusi", "Asia/Shanghai"}
	case "JNA":
		return Location{"Januaria", "America/Sao_Paulo"}
	case "JNB":
		return Location{"Johannesburg", "Africa/Johannesburg"}
	case "JNG":
		return Location{"Jining", "Asia/Shanghai"}
	case "JNI":
		return Location{"Junin", "America/Argentina/Buenos_Aires"}
	case "JNU":
		return Location{"Juneau", "America/Juneau"}
	case "JNX":
		return Location{"Naxos Island", "Europe/Athens"}
	case "JNZ":
		return Location{"Jinzhou", "Asia/Shanghai"}
	case "JOE":
		return Location{"Joensuu / Liperi", "Europe/Helsinki"}
	case "JOG":
		return Location{"Yogyakarta-Java Island", "Asia/Jakarta"}
	case "JOH":
		return Location{"Port St Johns", "Africa/Johannesburg"}
	case "JOI":
		return Location{"Joinville", "America/Sao_Paulo"}
	case "JOK":
		return Location{"Yoshkar-Ola", "Europe/Moscow"}
	case "JOL":
		return Location{"", "Asia/Manila"}
	case "JOM":
		return Location{"Njombe", "Africa/Dar_es_Salaam"}
	case "JOS":
		return Location{"Jos", "Africa/Lagos"}
	case "JOT":
		return Location{"Joliet", "America/Chicago"}
	case "JPA":
		return Location{"Joao Pessoa", "America/Fortaleza"}
	case "JPR":
		return Location{"Ji-Parana", "America/Porto_Velho"}
	case "JQA":
		return Location{"Uummannaq", "America/Nuuk"}
	case "JQE":
		return Location{"Jaque", "America/Panama"}
	case "JRF":
		return Location{"Kapolei", "Pacific/Honolulu"}
	case "JRG":
		return Location{"Jharsuguda", "Asia/Kolkata"}
	case "JRH":
		return Location{"Jorhat", "Asia/Kolkata"}
	case "JRN":
		return Location{"Juruena", "America/Cuiaba"}
	case "JRO":
		return Location{"Arusha", "Africa/Dar_es_Salaam"}
	case "JSA":
		return Location{"", "Asia/Kolkata"}
	case "JSH":
		return Location{"Crete Island", "Europe/Athens"}
	case "JSI":
		return Location{"Skiathos", "Europe/Athens"}
	case "JSM":
		return Location{"Chubut", "America/Argentina/Catamarca"}
	case "JSR":
		return Location{"Jashahor", "Asia/Dhaka"}
	case "JST":
		return Location{"Johnstown", "America/New_York"}
	case "JSU":
		return Location{"Maniitsoq", "America/Nuuk"}
	case "JSY":
		return Location{"Syros Island", "Europe/Athens"}
	case "JTC":
		return Location{"Bauru", "America/Sao_Paulo"}
	case "JTI":
		return Location{"Jatai", "America/Sao_Paulo"}
	case "JTR":
		return Location{"Santorini Island", "Europe/Athens"}
	case "JTY":
		return Location{"Astypalaia Island", "Europe/Athens"}
	case "JUA":
		return Location{"Juara", "America/Cuiaba"}
	case "JUB":
		return Location{"Juba", "Africa/Juba"}
	case "JUI":
		return Location{"Juist", "Europe/Berlin"}
	case "JUJ":
		return Location{"San Salvador de Jujuy", "America/Argentina/Jujuy"}
	case "JUL":
		return Location{"Juliaca", "America/Lima"}
	case "JUM":
		return Location{"Jumla", "Asia/Kathmandu"}
	case "JUN":
		return Location{"", "Australia/Brisbane"}
	case "JUR":
		return Location{"", "Australia/Perth"}
	case "JUT":
		return Location{"Jutigalpa", "America/Tegucigalpa"}
	case "JUV":
		return Location{"Upernavik", "America/Nuuk"}
	case "JUZ":
		return Location{"Quzhou", "Asia/Shanghai"}
	case "JVA":
		return Location{"Ankavandra", "Indian/Antananarivo"}
	case "JVI":
		return Location{"Manville", "America/New_York"}
	case "JVL":
		return Location{"Janesville", "America/Chicago"}
	case "JWA":
		return Location{"", "Africa/Gaborone"}
	case "JWN":
		return Location{"", "Asia/Tehran"}
	case "JWO":
		return Location{"", "Asia/Seoul"}
	case "JXA":
		return Location{"Jixi", "Asia/Shanghai"}
	case "JXN":
		return Location{"Jackson", "America/Detroit"}
	case "JYR":
		return Location{"", "Asia/Tehran"}
	case "JYV":
		return Location{"Jyvaskylan Maalaiskunta", "Europe/Helsinki"}
	case "JZH":
		return Location{"Jiuzhaigou", "Asia/Shanghai"}
	case "KAA":
		return Location{"Kasama", "Africa/Lusaka"}
	case "KAB":
		return Location{"Kariba", "Africa/Harare"}
	case "KAC":
		return Location{"Kamishly", "Asia/Damascus"}
	case "KAD":
		return Location{"Kaduna", "Africa/Lagos"}
	case "KAG":
		return Location{"", "Asia/Seoul"}
	case "KAJ":
		return Location{"Kajaani", "Europe/Helsinki"}
	case "KAL":
		return Location{"Kaltag", "America/Anchorage"}
	case "KAN":
		return Location{"Kano", "Africa/Lagos"}
	case "KAO":
		return Location{"Kuusamo", "Europe/Helsinki"}
	case "KAP":
		return Location{"Kapanga", "Africa/Lubumbashi"}
	case "KAR":
		return Location{"Kamarang", "America/Guyana"}
	case "KAT":
		return Location{"Kaitaia", "Pacific/Auckland"}
	case "KAU":
		return Location{"Kauhava", "Europe/Helsinki"}
	case "KAV":
		return Location{"", "America/Caracas"}
	case "KAW":
		return Location{"Kawthoung", "Asia/Yangon"}
	case "KAX":
		return Location{"", "Australia/Perth"}
	case "KAY":
		return Location{"Wakaya Island", "Pacific/Fiji"}
	case "KAZ":
		return Location{"Kao-Celebes Island", "Asia/Jayapura"}
	case "KBA":
		return Location{"Kabala", "Africa/Freetown"}
	case "KBB":
		return Location{"Kirkimbie", "Australia/Darwin"}
	case "KBG":
		return Location{"Kabalega Falls", "Africa/Kampala"}
	case "KBH":
		return Location{"Kalat", "Asia/Karachi"}
	case "KBI":
		return Location{"Kribi", "Africa/Douala"}
	case "KBJ":
		return Location{"", "Australia/Darwin"}
	case "KBK":
		return Location{"Kushinagar", "Asia/Kolkata"}
	case "KBL":
		return Location{"Kabul", "Asia/Kabul"}
	case "KBM":
		return Location{"", "Pacific/Port_Moresby"}
	case "KBN":
		return Location{"Kabinda", "Africa/Lubumbashi"}
	case "KBO":
		return Location{"Kabalo", "Africa/Lubumbashi"}
	case "KBP":
		return Location{"Kiev", "Europe/Kyiv"}
	case "KBQ":
		return Location{"Kasungu", "Africa/Blantyre"}
	case "KBR":
		return Location{"Kota Baharu", "Asia/Kuala_Lumpur"}
	case "KBS":
		return Location{"Bo", "Africa/Freetown"}
	case "KBU":
		return Location{"Laut Island", "Asia/Makassar"}
	case "KBV":
		return Location{"Krabi", "Asia/Bangkok"}
	case "KBY":
		return Location{"", "Australia/Adelaide"}
	case "KBZ":
		return Location{"", "Pacific/Auckland"}
	case "KCA":
		return Location{"Kuqa", "Asia/Shanghai"}
	case "KCB":
		return Location{"Kasikasima", "America/Paramaribo"}
	case "KCE":
		return Location{"", "Australia/Brisbane"}
	case "KCF":
		return Location{"Kadanwari", "Asia/Karachi"}
	case "KCH":
		return Location{"Kuching", "Asia/Kuching"}
	case "KCK":
		return Location{"Kirensk", "Asia/Irkutsk"}
	case "KCM":
		return Location{"Kahramanmaras", "Europe/Istanbul"}
	case "KCO":
		return Location{"", "Europe/Istanbul"}
	case "KCS":
		return Location{"", "Australia/Darwin"}
	case "KCT":
		return Location{"Galle", "Asia/Colombo"}
	case "KCU":
		return Location{"Masindi", "Africa/Kampala"}
	case "KCZ":
		return Location{"Nankoku", "Asia/Tokyo"}
	case "KDA":
		return Location{"Kolda", "Africa/Dakar"}
	case "KDB":
		return Location{"", "Australia/Perth"}
	case "KDC":
		return Location{"Kandi", "Africa/Porto-Novo"}
	case "KDD":
		return Location{"Khuzdar", "Asia/Karachi"}
	case "KDH":
		return Location{"", "Asia/Kabul"}
	case "KDI":
		return Location{"Kendari-Celebes Island", "Asia/Makassar"}
	case "KDJ":
		return Location{"N'Djole", "Africa/Libreville"}
	case "KDK":
		return Location{"Kodiak", "America/Anchorage"}
	case "KDL":
		return Location{"Kardla", "Europe/Tallinn"}
	case "KDM":
		return Location{"Huvadhu Atoll", "Indian/Maldives"}
	case "KDN":
		return Location{"Ndende", "Africa/Libreville"}
	case "KDO":
		return Location{"Kadhdhoo", "Indian/Maldives"}
	case "KDR":
		return Location{"Kandrian", "Pacific/Port_Moresby"}
	case "KDT":
		return Location{"Nakhon Pathom", "Asia/Bangkok"}
	case "KDU":
		return Location{"Skardu", "Asia/Karachi"}
	case "KDV":
		return Location{"Vunisea", "Pacific/Fiji"}
	case "KDX":
		return Location{"Kadugli", "Africa/Khartoum"}
	case "KDY":
		return Location{"Tyopliy Klyuch", "Asia/Khandyga"}
	case "KEA":
		return Location{"Kerki", "Asia/Ashgabat"}
	case "KEC":
		return Location{"Kasenga", "Africa/Lubumbashi"}
	case "KED":
		return Location{"Kaedi", "Africa/Dakar"}
	case "KEE":
		return Location{"Kelle", "Africa/Brazzaville"}
	case "KEF":
		return Location{"Reykjavik", "Atlantic/Reykjavik"}
	case "KEI":
		return Location{"Kepi-Papua Island", "Asia/Jayapura"}
	case "KEJ":
		return Location{"Kemerovo", "Asia/Novokuznetsk"}
	case "KEL":
		return Location{"Kiel", "Europe/Berlin"}
	case "KEM":
		return Location{"Kemi / Tornio", "Europe/Helsinki"}
	case "KEN":
		return Location{"Kenema", "Africa/Freetown"}
	case "KEO":
		return Location{"Odienne", "Africa/Abidjan"}
	case "KEP":
		return Location{"Nepalgunj", "Asia/Kathmandu"}
	case "KEQ":
		return Location{"Kebar-Papua Island", "Asia/Jayapura"}
	case "KER":
		return Location{"Kerman", "Asia/Tehran"}
	case "KET":
		return Location{"Kengtung", "Asia/Yangon"}
	case "KEU":
		return Location{"Keekorok", "Africa/Nairobi"}
	case "KEV":
		return Location{"Halli / Kuorevesi", "Europe/Helsinki"}
	case "KEW":
		return Location{"Keewaywin", "America/Winnipeg"}
	case "KEY":
		return Location{"Kericho", "Africa/Nairobi"}
	case "KFA":
		return Location{"Kiffa", "Africa/Nouakchott"}
	case "KFE":
		return Location{"Cloudbreak Village", "Australia/Perth"}
	case "KFG":
		return Location{"", "Australia/Darwin"}
	case "KFP":
		return Location{"False Pass", "America/Nome"}
	case "KFS":
		return Location{"Kastamonu", "Europe/Istanbul"}
	case "KFZ":
		return Location{"Kukes", "Europe/Tirane"}
	case "KGA":
		return Location{"Kananga", "Africa/Lubumbashi"}
	case "KGC":
		return Location{"", "Australia/Adelaide"}
	case "KGD":
		return Location{"Kaliningrad", "Europe/Kaliningrad"}
	case "KGE":
		return Location{"Kagau Island", "Pacific/Guadalcanal"}
	case "KGF":
		return Location{"Karaganda", "Asia/Almaty"}
	case "KGG":
		return Location{"Kedougou", "Africa/Dakar"}
	case "KGI":
		return Location{"Kalgoorlie", "Australia/Perth"}
	case "KGJ":
		return Location{"Karonga", "Africa/Blantyre"}
	case "KGK":
		return Location{"Koliganek", "America/Anchorage"}
	case "KGL":
		return Location{"Kigali", "Africa/Kigali"}
	case "KGN":
		return Location{"Kasongo", "Africa/Lubumbashi"}
	case "KGO":
		return Location{"Kirovograd", "Europe/Kyiv"}
	case "KGP":
		return Location{"Kogalym", "Asia/Yekaterinburg"}
	case "KGS":
		return Location{"Kos Island", "Europe/Athens"}
	case "KGU":
		return Location{"Keningau", "Asia/Kuching"}
	case "KGY":
		return Location{"", "Australia/Brisbane"}
	case "KHC":
		return Location{"Kerch", "Europe/Simferopol"}
	case "KHD":
		return Location{"", "Asia/Tehran"}
	case "KHE":
		return Location{"Kherson", "Europe/Kyiv"}
	case "KHG":
		return Location{"Kashgar", "Asia/Shanghai"}
	case "KHH":
		return Location{"Kaohsiung City", "Asia/Taipei"}
	case "KHI":
		return Location{"Karachi", "Asia/Karachi"}
	case "KHJ":
		return Location{"", "Europe/Helsinki"}
	case "KHK":
		return Location{"", "Asia/Tehran"}
	case "KHM":
		return Location{"Kanti", "Asia/Yangon"}
	case "KHN":
		return Location{"Nanchang", "Asia/Shanghai"}
	case "KHR":
		return Location{"", "Asia/Ulaanbaatar"}
	case "KHS":
		return Location{"Khasab", "Asia/Muscat"}
	case "KHT":
		return Location{"Khost", "Asia/Kabul"}
	case "KHV":
		return Location{"Khabarovsk", "Asia/Vladivostok"}
	case "KHW":
		return Location{"Khwai River Lodge", "Africa/Gaborone"}
	case "KHY":
		return Location{"Khoy", "Asia/Tehran"}
	case "KHZ":
		return Location{"Kauehi", "Pacific/Tahiti"}
	case "KIA":
		return Location{"Kaieteur Falls", "America/Guyana"}
	case "KIC":
		return Location{"King City", "America/Los_Angeles"}
	case "KID":
		return Location{"Kristianstad", "Europe/Stockholm"}
	case "KIE":
		return Location{"Kieta", "Pacific/Bougainville"}
	case "KIF":
		return Location{"Kingfisher Lake", "America/Winnipeg"}
	case "KIH":
		return Location{"Kish Island", "Asia/Tehran"}
	case "KIJ":
		return Location{"Niigata", "Asia/Tokyo"}
	case "KIK":
		return Location{"Kirkuk", "Asia/Baghdad"}
	case "KIM":
		return Location{"Kimberley", "Africa/Johannesburg"}
	case "KIN":
		return Location{"Kingston", "America/Jamaica"}
	case "KIP":
		return Location{"Wichita Falls", "America/Chicago"}
	case "KIR":
		return Location{"Killarney", "Europe/Dublin"}
	case "KIS":
		return Location{"Kisumu", "Africa/Nairobi"}
	case "KIT":
		return Location{"Kithira Island", "Europe/Athens"}
	case "KIW":
		return Location{"Kitwe", "Africa/Lusaka"}
	case "KIX":
		return Location{"Osaka", "Asia/Tokyo"}
	case "KIY":
		return Location{"Kilwa Masoko", "Africa/Dar_es_Salaam"}
	case "KJA":
		return Location{"Krasnoyarsk", "Asia/Krasnoyarsk"}
	case "KJB":
		return Location{"Kurnool", "Asia/Kolkata"}
	case "KJH":
		return Location{"Kaili", "Asia/Shanghai"}
	case "KJK":
		return Location{"Wevelgem", "Europe/Brussels"}
	case "KJP":
		return Location{"Kerama", "Asia/Tokyo"}
	case "KJT":
		return Location{"Kertajati", "Asia/Jakarta"}
	case "KKA":
		return Location{"Koyuk", "America/Anchorage"}
	case "KKC":
		return Location{"Khon Kaen", "Asia/Bangkok"}
	case "KKD":
		return Location{"Kokoda", "Pacific/Port_Moresby"}
	case "KKE":
		return Location{"Kerikeri", "Pacific/Auckland"}
	case "KKH":
		return Location{"Kongiganak", "America/Nome"}
	case "KKJ":
		return Location{"Kitakyushu", "Asia/Tokyo"}
	case "KKM":
		return Location{"", "Asia/Bangkok"}
	case "KKN":
		return Location{"Kirkenes", "Europe/Oslo"}
	case "KKO":
		return Location{"", "Pacific/Auckland"}
	case "KKP":
		return Location{"Koolburra", "Australia/Brisbane"}
	case "KKQ":
		return Location{"Krasnoselkup", "Asia/Yekaterinburg"}
	case "KKR":
		return Location{"", "Pacific/Tahiti"}
	case "KKS":
		return Location{"", "Asia/Tehran"}
	case "KKW":
		return Location{"", "Africa/Kinshasa"}
	case "KKX":
		return Location{"", "Asia/Tokyo"}
	case "KKY":
		return Location{"Kilkenny", "Europe/Dublin"}
	case "KKZ":
		return Location{"Kaoh Kong", "Asia/Phnom_Penh"}
	case "KLB":
		return Location{"Kalabo", "Africa/Lusaka"}
	case "KLC":
		return Location{"Kaolack", "Africa/Dakar"}
	case "KLD":
		return Location{"Tver", "Europe/Moscow"}
	case "KLE":
		return Location{"Kaele", "Africa/Douala"}
	case "KLF":
		return Location{"Kaluga", "Europe/Moscow"}
	case "KLG":
		return Location{"Kalskag", "America/Anchorage"}
	case "KLH":
		return Location{"", "Asia/Kolkata"}
	case "KLI":
		return Location{"", "Africa/Kinshasa"}
	case "KLJ":
		return Location{"Klaipeda", "Europe/Vilnius"}
	case "KLK":
		return Location{"Kalokol", "Africa/Nairobi"}
	case "KLM":
		return Location{"", "Asia/Tehran"}
	case "KLN":
		return Location{"Larsen Bay", "America/Anchorage"}
	case "KLO":
		return Location{"Kalibo", "Asia/Manila"}
	case "KLQ":
		return Location{"Keluang-Sumatra Island", "Asia/Jakarta"}
	case "KLR":
		return Location{"", "Europe/Stockholm"}
	case "KLS":
		return Location{"Kelso", "America/Los_Angeles"}
	case "KLU":
		return Location{"Klagenfurt am Worthersee", "Europe/Vienna"}
	case "KLV":
		return Location{"Karlovy Vary", "Europe/Prague"}
	case "KLW":
		return Location{"Klawock", "America/Sitka"}
	case "KLX":
		return Location{"Kalamata", "Europe/Athens"}
	case "KLY":
		return Location{"Kalima", "Africa/Lubumbashi"}
	case "KLZ":
		return Location{"Kleinsee", "Africa/Johannesburg"}
	case "KMA":
		return Location{"Kerema", "Pacific/Port_Moresby"}
	case "KME":
		return Location{"Kamembe", "Africa/Kigali"}
	case "KMG":
		return Location{"Kunming", "Asia/Shanghai"}
	case "KMH":
		return Location{"Kuruman", "Africa/Johannesburg"}
	case "KMI":
		return Location{"Miyazaki", "Asia/Tokyo"}
	case "KMJ":
		return Location{"Kumamoto", "Asia/Tokyo"}
	case "KMK":
		return Location{"Makabana", "Africa/Brazzaville"}
	case "KML":
		return Location{"", "Australia/Brisbane"}
	case "KMN":
		return Location{"Kamina", "Africa/Lubumbashi"}
	case "KMO":
		return Location{"Manokotak", "America/Anchorage"}
	case "KMP":
		return Location{"Keetmanshoop", "Africa/Windhoek"}
	case "KMQ":
		return Location{"Kanazawa", "Asia/Tokyo"}
	case "KMR":
		return Location{"Karimui", "Pacific/Port_Moresby"}
	case "KMS":
		return Location{"Kumasi", "Africa/Accra"}
	case "KMU":
		return Location{"", "Africa/Mogadishu"}
	case "KMV":
		return Location{"Kalemyo", "Asia/Yangon"}
	case "KMW":
		return Location{"Kostroma", "Europe/Moscow"}
	case "KMX":
		return Location{"", "Asia/Riyadh"}
	case "KMZ":
		return Location{"Kaoma", "Africa/Lusaka"}
	case "KNA":
		return Location{"Vina Del Mar", "America/Santiago"}
	case "KNB":
		return Location{"Kanab", "America/Denver"}
	case "KND":
		return Location{"Kindu", "Africa/Lubumbashi"}
	case "KNF":
		return Location{"Marham", "Europe/London"}
	case "KNG":
		return Location{"Kaimana-Papua Island", "Asia/Jayapura"}
	case "KNH":
		return Location{"Shang-I", "Asia/Taipei"}
	case "KNI":
		return Location{"", "Australia/Perth"}
	case "KNJ":
		return Location{"Kindamba", "Africa/Brazzaville"}
	case "KNK":
		return Location{"Kokhanok", "America/Anchorage"}
	case "KNM":
		return Location{"Kaniama", "Africa/Lubumbashi"}
	case "KNN":
		return Location{"Kankan", "Africa/Conakry"}
	case "KNO":
		return Location{"Medan-Sumatra Island", "Asia/Jakarta"}
	case "KNQ":
		return Location{"Kone", "Pacific/Noumea"}
	case "KNR":
		return Location{"Kangan", "Asia/Tehran"}
	case "KNS":
		return Location{"", "Australia/Hobart"}
	case "KNT":
		return Location{"Kennett", "America/Chicago"}
	case "KNU":
		return Location{"Kanpur", "Asia/Kolkata"}
	case "KNW":
		return Location{"New Stuyahok", "America/Anchorage"}
	case "KNX":
		return Location{"Kununurra", "Australia/Perth"}
	case "KNZ":
		return Location{"Kenieba", "Africa/Bamako"}
	case "KOA":
		return Location{"Kailua/Kona", "Pacific/Honolulu"}
	case "KOC":
		return Location{"Koumac", "Pacific/Noumea"}
	case "KOE":
		return Location{"Kupang-Timor Island", "Asia/Makassar"}
	case "KOF":
		return Location{"Komatipoort", "Africa/Johannesburg"}
	case "KOH":
		return Location{"", "Australia/Brisbane"}
	case "KOI":
		return Location{"Orkney Islands", "Europe/London"}
	case "KOJ":
		return Location{"Kagoshima", "Asia/Tokyo"}
	case "KOK":
		return Location{"Kokkola / Kruunupyy", "Europe/Helsinki"}
	case "KOO":
		return Location{"Kongolo", "Africa/Lubumbashi"}
	case "KOP":
		return Location{"", "Asia/Bangkok"}
	case "KOQ":
		return Location{"Kothen", "Europe/Berlin"}
	case "KOS":
		return Location{"Sihanukville", "Asia/Phnom_Penh"}
	case "KOT":
		return Location{"Kotlik", "America/Nome"}
	case "KOU":
		return Location{"Koulamoutou", "Africa/Libreville"}
	case "KOV":
		return Location{"Kokshetau", "Asia/Almaty"}
	case "KOW":
		return Location{"Ganzhou", "Asia/Shanghai"}
	case "KOX":
		return Location{"Kokonau-Papua Island", "Asia/Jayapura"}
	case "KPC":
		return Location{"Port Clarence", "America/Nome"}
	case "KPI":
		return Location{"Kapit", "Asia/Kuching"}
	case "KPM":
		return Location{"", "Pacific/Port_Moresby"}
	case "KPN":
		return Location{"Kipnuk", "America/Nome"}
	case "KPO":
		return Location{"Pohang", "Asia/Seoul"}
	case "KPP":
		return Location{"", "Australia/Brisbane"}
	case "KPS":
		return Location{"", "Australia/Sydney"}
	case "KPV":
		return Location{"Perryville", "America/Anchorage"}
	case "KPW":
		return Location{"Keperveem", "Asia/Anadyr"}
	case "KQH":
		return Location{"Kishangarh", "Asia/Kolkata"}
	case "KQT":
		return Location{"Kurgan-Tyube", "Asia/Dushanbe"}
	case "KRA":
		return Location{"", "Australia/Melbourne"}
	case "KRB":
		return Location{"", "Australia/Brisbane"}
	case "KRC":
		return Location{"Sungai Penuh", "Asia/Jakarta"}
	case "KRE":
		return Location{"Kirundo", "Africa/Bujumbura"}
	case "KRF":
		return Location{"Kramfors / Solleftea", "Europe/Stockholm"}
	case "KRG":
		return Location{"Karasabai", "America/Guyana"}
	case "KRH":
		return Location{"Redhill", "Europe/London"}
	case "KRI":
		return Location{"Kikori", "Pacific/Port_Moresby"}
	case "KRJ":
		return Location{"", "Pacific/Port_Moresby"}
	case "KRK":
		return Location{"Krakow", "Europe/Warsaw"}
	case "KRL":
		return Location{"Korla", "Asia/Shanghai"}
	case "KRM":
		return Location{"Karanambo", "America/Guyana"}
	case "KRN":
		return Location{"", "Europe/Stockholm"}
	case "KRO":
		return Location{"Kurgan", "Asia/Yekaterinburg"}
	case "KRP":
		return Location{"Karup", "Europe/Copenhagen"}
	case "KRQ":
		return Location{"Kramatorsk", "Europe/Kyiv"}
	case "KRR":
		return Location{"Krasnodar", "Europe/Moscow"}
	case "KRS":
		return Location{"Kjevik", "Europe/Oslo"}
	case "KRT":
		return Location{"Khartoum", "Africa/Khartoum"}
	case "KRW":
		return Location{"Krasnovodsk", "Asia/Ashgabat"}
	case "KRY":
		return Location{"Karamay", "Asia/Shanghai"}
	case "KRZ":
		return Location{"Kiri", "Africa/Kinshasa"}
	case "KSA":
		return Location{"Okat", "Pacific/Kosrae"}
	case "KSC":
		return Location{"Kosice", "Europe/Bratislava"}
	case "KSD":
		return Location{"Karlstad", "Europe/Stockholm"}
	case "KSE":
		return Location{"Kasese", "Africa/Kampala"}
	case "KSF":
		return Location{"Kassel", "Europe/Berlin"}
	case "KSH":
		return Location{"Kermanshah", "Asia/Tehran"}
	case "KSI":
		return Location{"Kissidougou", "Africa/Conakry"}
	case "KSJ":
		return Location{"Kasos Island", "Europe/Athens"}
	case "KSK":
		return Location{"", "Europe/Stockholm"}
	case "KSL":
		return Location{"Kassala", "Africa/Khartoum"}
	case "KSM":
		return Location{"St Mary's", "America/Nome"}
	case "KSN":
		return Location{"Kostanay", "Asia/Qostanay"}
	case "KSO":
		return Location{"Kastoria", "Europe/Athens"}
	case "KSQ":
		return Location{"Karshi", "Asia/Samarkand"}
	case "KSS":
		return Location{"Sikasso", "Africa/Bamako"}
	case "KST":
		return Location{"Kosti", "Africa/Khartoum"}
	case "KSU":
		return Location{"Kvernberget", "Europe/Oslo"}
	case "KSV":
		return Location{"", "Australia/Brisbane"}
	case "KSW":
		return Location{"Kiryat Shmona", "Asia/Jerusalem"}
	case "KSY":
		return Location{"Kars", "Europe/Istanbul"}
	case "KSZ":
		return Location{"Kotlas", "Europe/Moscow"}
	case "KTA":
		return Location{"Karratha", "Australia/Perth"}
	case "KTD":
		return Location{"", "Asia/Tokyo"}
	case "KTE":
		return Location{"Kerteh", "Asia/Kuala_Lumpur"}
	case "KTF":
		return Location{"", "Pacific/Auckland"}
	case "KTG":
		return Location{"Ketapang-Borneo Island", "Asia/Pontianak"}
	case "KTI":
		return Location{"Phnom Penh", "Asia/Phnom_Penh"}
	case "KTL":
		return Location{"Kitale", "Africa/Nairobi"}
	case "KTM":
		return Location{"Kathmandu", "Asia/Kathmandu"}
	case "KTN":
		return Location{"Ketchikan", "America/Sitka"}
	case "KTO":
		return Location{"Kato", "America/Guyana"}
	case "KTP":
		return Location{"Tinson Pen", "America/Jamaica"}
	case "KTQ":
		return Location{"", "Europe/Helsinki"}
	case "KTR":
		return Location{"", "Australia/Darwin"}
	case "KTS":
		return Location{"Brevig Mission", "America/Nome"}
	case "KTT":
		return Location{"Kittila", "Europe/Helsinki"}
	case "KTU":
		return Location{"", "Asia/Kolkata"}
	case "KTW":
		return Location{"Katowice", "Europe/Warsaw"}
	case "KTX":
		return Location{"Koutiala", "Africa/Bamako"}
	case "KTY":
		return Location{"Kalutara", "Asia/Colombo"}
	case "KUA":
		return Location{"Kuantan", "Asia/Kuala_Lumpur"}
	case "KUC":
		return Location{"Kuria", "Pacific/Tarawa"}
	case "KUD":
		return Location{"Kudat", "Asia/Kuching"}
	case "KUF":
		return Location{"Samara", "Europe/Samara"}
	case "KUG":
		return Location{"", "Australia/Brisbane"}
	case "KUH":
		return Location{"Kushiro", "Asia/Tokyo"}
	case "KUK":
		return Location{"Kasigluk", "America/Nome"}
	case "KUL":
		return Location{"Kuala Lumpur", "Asia/Kuala_Lumpur"}
	case "KUM":
		return Location{"", "Asia/Tokyo"}
	case "KUN":
		return Location{"Kaunas", "Europe/Vilnius"}
	case "KUO":
		return Location{"Kuopio / Siilinjarvi", "Europe/Helsinki"}
	case "KUQ":
		return Location{"Kuri", "Pacific/Port_Moresby"}
	case "KUS":
		return Location{"Kulusuk", "America/Nuuk"}
	case "KUT":
		return Location{"Kutaisi", "Asia/Tbilisi"}
	case "KUU":
		return Location{"", "Asia/Kolkata"}
	case "KUV":
		return Location{"Kunsan", "Asia/Seoul"}
	case "KVA":
		return Location{"Kavala", "Europe/Athens"}
	case "KVB":
		return Location{"Skovde", "Europe/Stockholm"}
	case "KVC":
		return Location{"King Cove", "America/Nome"}
	case "KVG":
		return Location{"Kavieng", "Pacific/Port_Moresby"}
	case "KVK":
		return Location{"Apatity", "Europe/Moscow"}
	case "KVL":
		return Location{"Kivalina", "America/Nome"}
	case "KVM":
		return Location{"Markovo", "Asia/Anadyr"}
	case "KVX":
		return Location{"Kirov", "Europe/Kirov"}
	case "KWA":
		return Location{"Kwajalein", "Pacific/Kwajalein"}
	case "KWE":
		return Location{"Guiyang", "Asia/Shanghai"}
	case "KWG":
		return Location{"Kryvyi Rih", "Europe/Kyiv"}
	case "KWH":
		return Location{"Khwahan", "Asia/Dushanbe"}
	case "KWI":
		return Location{"Kuwait City", "Asia/Kuwait"}
	case "KWJ":
		return Location{"Gwangju", "Asia/Seoul"}
	case "KWK":
		return Location{"Kwigillingok", "America/Nome"}
	case "KWL":
		return Location{"Guilin City", "Asia/Shanghai"}
	case "KWM":
		return Location{"Kowanyama", "Australia/Brisbane"}
	case "KWN":
		return Location{"Quinhagak", "America/Anchorage"}
	case "KWO":
		return Location{"Kawito", "Pacific/Port_Moresby"}
	case "KWS":
		return Location{"Kwailabesi", "Pacific/Guadalcanal"}
	case "KWT":
		return Location{"Kwethluk", "America/Anchorage"}
	case "KWZ":
		return Location{"", "Africa/Lubumbashi"}
	case "KXB":
		return Location{"Kolaka", "Asia/Makassar"}
	case "KXD":
		return Location{"Kondinskoye", "Asia/Yekaterinburg"}
	case "KXE":
		return Location{"Klerksdorp", "Africa/Johannesburg"}
	case "KXF":
		return Location{"Koro Island", "Pacific/Fiji"}
	case "KXK":
		return Location{"Komsomolsk-on-Amur", "Asia/Vladivostok"}
	case "KXU":
		return Location{"Katiu", "Pacific/Tahiti"}
	case "KYA":
		return Location{"Konya", "Europe/Istanbul"}
	case "KYD":
		return Location{"Orchid Island", "Asia/Taipei"}
	case "KYE":
		return Location{"Tripoli", "Asia/Beirut"}
	case "KYF":
		return Location{"", "Australia/Perth"}
	case "KYI":
		return Location{"Yalata Mission", "Australia/Adelaide"}
	case "KYK":
		return Location{"Karluk", "America/Anchorage"}
	case "KYP":
		return Location{"Kyaukpyu", "Asia/Yangon"}
	case "KYS":
		return Location{"", "Africa/Bamako"}
	case "KYT":
		return Location{"Kyauktu", "Asia/Yangon"}
	case "KYU":
		return Location{"Koyukuk", "America/Anchorage"}
	case "KYZ":
		return Location{"Kyzyl", "Asia/Krasnoyarsk"}
	case "KZC":
		return Location{"", "Asia/Phnom_Penh"}
	case "KZF":
		return Location{"Kaintiba", "Pacific/Port_Moresby"}
	case "KZG":
		return Location{"Kitzingen", "Europe/Berlin"}
	case "KZI":
		return Location{"Kozani", "Europe/Athens"}
	case "KZN":
		return Location{"Kazan", "Europe/Moscow"}
	case "KZO":
		return Location{"Kzyl-Orda", "Asia/Qyzylorda"}
	case "KZR":
		return Location{"Altintas", "Europe/Istanbul"}
	case "KZS":
		return Location{"Kastelorizo Island", "Europe/Athens"}
	case "LAA":
		return Location{"Lamar", "America/Denver"}
	case "LAD":
		return Location{"Luanda", "Africa/Luanda"}
	case "LAE":
		return Location{"Nadzab", "Pacific/Port_Moresby"}
	case "LAF":
		return Location{"Lafayette", "America/Indiana/Indianapolis"}
	case "LAH":
		return Location{"Labuha-Halmahera Island", "Asia/Jayapura"}
	case "LAI":
		return Location{"Lannion", "Europe/Paris"}
	case "LAJ":
		return Location{"Lages", "America/Sao_Paulo"}
	case "LAK":
		return Location{"Aklavik", "America/Inuvik"}
	case "LAL":
		return Location{"Lakeland", "America/New_York"}
	case "LAM":
		return Location{"Los Alamos", "America/Denver"}
	case "LAN":
		return Location{"Lansing", "America/Detroit"}
	case "LAO":
		return Location{"Laoag City", "Asia/Manila"}
	case "LAP":
		return Location{"La Paz", "America/Mazatlan"}
	case "LAQ":
		return Location{"Al Bayda'", "Africa/Tripoli"}
	case "LAR":
		return Location{"Laramie", "America/Denver"}
	case "LAS":
		return Location{"Las Vegas", "America/Los_Angeles"}
	case "LAU":
		return Location{"Lamu", "Africa/Nairobi"}
	case "LAW":
		return Location{"Lawton", "America/Chicago"}
	case "LAX":
		return Location{"Los Angeles", "America/Los_Angeles"}
	case "LAY":
		return Location{"Ladysmith", "Africa/Johannesburg"}
	case "LAZ":
		return Location{"Bom Jesus Da Lapa", "America/Bahia"}
	case "LBA":
		return Location{"Leeds", "Europe/London"}
	case "LBB":
		return Location{"Lubbock", "America/Chicago"}
	case "LBC":
		return Location{"Lubeck", "Europe/Berlin"}
	case "LBD":
		return Location{"Khudzhand", "Asia/Dushanbe"}
	case "LBE":
		return Location{"Latrobe", "America/New_York"}
	case "LBF":
		return Location{"North Platte", "America/Chicago"}
	case "LBG":
		return Location{"Paris", "Europe/Paris"}
	case "LBI":
		return Location{"Albi/Le Sequestre", "Europe/Paris"}
	case "LBJ":
		return Location{"Labuan Bajo-Flores Island", "Asia/Makassar"}
	case "LBL":
		return Location{"Liberal", "America/Chicago"}
	case "LBO":
		return Location{"Lusambo", "Africa/Lubumbashi"}
	case "LBQ":
		return Location{"Lambarene", "Africa/Libreville"}
	case "LBR":
		return Location{"Labrea", "America/Manaus"}
	case "LBS":
		return Location{"", "Pacific/Fiji"}
	case "LBT":
		return Location{"Lumberton", "America/New_York"}
	case "LBU":
		return Location{"Labuan", "Asia/Kuching"}
	case "LBV":
		return Location{"Libreville", "Africa/Libreville"}
	case "LBW":
		return Location{"Long Bawan-Borneo Island", "Asia/Makassar"}
	case "LBX":
		return Location{"", "Asia/Manila"}
	case "LBY":
		return Location{"La Baule-Escoublac", "Europe/Paris"}
	case "LBZ":
		return Location{"Lucapa", "Africa/Luanda"}
	case "LCA":
		return Location{"Larnarca", "Asia/Nicosia"}
	case "LCC":
		return Location{"", "Europe/Rome"}
	case "LCD":
		return Location{"Louis Trichardt", "Africa/Johannesburg"}
	case "LCE":
		return Location{"La Ceiba", "America/Tegucigalpa"}
	case "LCF":
		return Location{"Rio Dulce", "America/Guatemala"}
	case "LCG":
		return Location{"Culleredo", "Europe/Madrid"}
	case "LCH":
		return Location{"Lake Charles", "America/Chicago"}
	case "LCI":
		return Location{"Laconia", "America/New_York"}
	case "LCJ":
		return Location{"Lodz", "Europe/Warsaw"}
	case "LCK":
		return Location{"Columbus", "America/New_York"}
	case "LCL":
		return Location{"Pinar del Rio", "America/Havana"}
	case "LCM":
		return Location{"La Cumbre", "America/Argentina/Cordoba"}
	case "LCN":
		return Location{"", "Australia/Adelaide"}
	case "LCO":
		return Location{"Lague", "Africa/Brazzaville"}
	case "LCQ":
		return Location{"Lake City", "America/New_York"}
	case "LCV":
		return Location{"Lucca", "Europe/Rome"}
	case "LCX":
		return Location{"Longyan", "Asia/Shanghai"}
	case "LCY":
		return Location{"London", "Europe/London"}
	case "LDA":
		return Location{"Malda", "Asia/Kolkata"}
	case "LDB":
		return Location{"Londrina", "America/Sao_Paulo"}
	case "LDC":
		return Location{"Lindeman Island", "Australia/Lindeman"}
	case "LDE":
		return Location{"Tarbes/Lourdes/Pyrenees", "Europe/Paris"}
	case "LDG":
		return Location{"Leshukonskoye", "Europe/Moscow"}
	case "LDH":
		return Location{"Lord Howe Island", "Australia/Lord_Howe"}
	case "LDI":
		return Location{"Lindi", "Africa/Dar_es_Salaam"}
	case "LDJ":
		return Location{"Linden", "America/New_York"}
	case "LDK":
		return Location{"Lidkoping", "Europe/Stockholm"}
	case "LDM":
		return Location{"Ludington", "America/Detroit"}
	case "LDN":
		return Location{"Lamidanda", "Asia/Kathmandu"}
	case "LDO":
		return Location{"Aurora", "America/Paramaribo"}
	case "LDS":
		return Location{"Yichun", "Asia/Shanghai"}
	case "LDU":
		return Location{"Lahad Datu", "Asia/Kuching"}
	case "LDV":
		return Location{"Landivisiau", "Europe/Paris"}
	case "LDX":
		return Location{"Saint-Laurent-du-Maroni", "America/Cayenne"}
	case "LDY":
		return Location{"Derry", "Europe/London"}
	case "LDZ":
		return Location{"Londolozi", "Africa/Johannesburg"}
	case "LEA":
		return Location{"Exmouth", "Australia/Perth"}
	case "LEB":
		return Location{"Lebanon", "America/New_York"}
	case "LEC":
		return Location{"Lencois", "America/Bahia"}
	case "LED":
		return Location{"St. Petersburg", "Europe/Moscow"}
	case "LEE":
		return Location{"Leesburg", "America/New_York"}
	case "LEF":
		return Location{"Lebakeng", "Africa/Maseru"}
	case "LEH":
		return Location{"Le Havre/Octeville", "Europe/Paris"}
	case "LEI":
		return Location{"Almeria", "Europe/Madrid"}
	case "LEJ":
		return Location{"Leipzig", "Europe/Berlin"}
	case "LEK":
		return Location{"", "Africa/Conakry"}
	case "LEL":
		return Location{"Lake Evella", "Australia/Darwin"}
	case "LEM":
		return Location{"Lemmon", "America/Denver"}
	case "LEN":
		return Location{"Leon", "Europe/Madrid"}
	case "LEP":
		return Location{"Leopoldina", "America/Sao_Paulo"}
	case "LEQ":
		return Location{"Land's End", "Europe/London"}
	case "LER":
		return Location{"", "Australia/Perth"}
	case "LES":
		return Location{"Lesobeng", "Africa/Maseru"}
	case "LET":
		return Location{"Leticia", "America/Bogota"}
	case "LEU":
		return Location{"Montferrer / Castellbo", "Europe/Madrid"}
	case "LEV":
		return Location{"Bureta", "Pacific/Fiji"}
	case "LEW":
		return Location{"Auburn/Lewiston", "America/New_York"}
	case "LEX":
		return Location{"Lexington", "America/New_York"}
	case "LEY":
		return Location{"Lelystad", "Europe/Amsterdam"}
	case "LEZ":
		return Location{"La Esperanza", "America/Tegucigalpa"}
	case "LFB":
		return Location{"Lumbo", "Africa/Maputo"}
	case "LFI":
		return Location{"Hampton", "America/New_York"}
	case "LFK":
		return Location{"Lufkin", "America/Chicago"}
	case "LFM":
		return Location{"Lamerd", "Asia/Tehran"}
	case "LFN":
		return Location{"Louisburg", "America/New_York"}
	case "LFO":
		return Location{"Kelafo", "Africa/Addis_Ababa"}
	case "LFP":
		return Location{"", "Australia/Brisbane"}
	case "LFR":
		return Location{"", "America/Caracas"}
	case "LFT":
		return Location{"Lafayette", "America/Chicago"}
	case "LFW":
		return Location{"Lome", "Africa/Lome"}
	case "LGA":
		return Location{"New York", "America/New_York"}
	case "LGB":
		return Location{"Long Beach", "America/Los_Angeles"}
	case "LGC":
		return Location{"Lagrange", "America/New_York"}
	case "LGD":
		return Location{"La Grande", "America/Los_Angeles"}
	case "LGF":
		return Location{"Yuma Proving Ground(Yuma)", "America/Phoenix"}
	case "LGG":
		return Location{"Liege", "Europe/Brussels"}
	case "LGH":
		return Location{"", "Australia/Adelaide"}
	case "LGI":
		return Location{"Deadman's Cay", "America/Nassau"}
	case "LGK":
		return Location{"Langkawi", "Asia/Kuala_Lumpur"}
	case "LGL":
		return Location{"Long Datih", "Asia/Kuching"}
	case "LGO":
		return Location{"Langeoog", "Europe/Berlin"}
	case "LGP":
		return Location{"Legazpi City", "Asia/Manila"}
	case "LGQ":
		return Location{"Lago Agrio", "America/Guayaquil"}
	case "LGR":
		return Location{"Cochrane", "America/Santiago"}
	case "LGS":
		return Location{"Malargue", "America/Argentina/Mendoza"}
	case "LGT":
		return Location{"", "America/Bogota"}
	case "LGU":
		return Location{"Logan", "America/Denver"}
	case "LGW":
		return Location{"London", "Europe/London"}
	case "LHA":
		return Location{"", "Europe/Berlin"}
	case "LHB":
		return Location{"Bruntingthorpe", "Europe/London"}
	case "LHE":
		return Location{"Lahore", "Asia/Karachi"}
	case "LHG":
		return Location{"", "Australia/Sydney"}
	case "LHI":
		return Location{"Lereh-Papua Island", "Asia/Jayapura"}
	case "LHK":
		return Location{"Guanghua", "Asia/Shanghai"}
	case "LHR":
		return Location{"London", "Europe/London"}
	case "LHS":
		return Location{"Las Heras", "America/Argentina/Rio_Gallegos"}
	case "LHU":
		return Location{"Muneambuanas", "Africa/Gaborone"}
	case "LHV":
		return Location{"Lock Haven", "America/New_York"}
	case "LHW":
		return Location{"Lanzhou", "Asia/Shanghai"}
	case "LIA":
		return Location{"Liangping", "Asia/Shanghai"}
	case "LIC":
		return Location{"Limon", "America/Denver"}
	case "LIE":
		return Location{"Libenge", "Africa/Kinshasa"}
	case "LIF":
		return Location{"Lifou", "Pacific/Noumea"}
	case "LIG":
		return Location{"Limoges/Bellegarde", "Europe/Paris"}
	case "LIH":
		return Location{"Lihue", "Pacific/Honolulu"}
	case "LII":
		return Location{"Mulia-Papua Island", "Asia/Jayapura"}
	case "LIL":
		return Location{"Lille/Lesquin", "Europe/Paris"}
	case "LIM":
		return Location{"Lima", "America/Lima"}
	case "LIN":
		return Location{"Milan", "Europe/Rome"}
	case "LIO":
		return Location{"Puerto Limon", "America/Costa_Rica"}
	case "LIP":
		return Location{"Lins", "America/Sao_Paulo"}
	case "LIQ":
		return Location{"Lisala", "Africa/Lubumbashi"}
	case "LIR":
		return Location{"Liberia", "America/Costa_Rica"}
	case "LIS":
		return Location{"Lisbon", "Europe/Lisbon"}
	case "LIT":
		return Location{"Little Rock", "America/Chicago"}
	case "LIW":
		return Location{"Loikaw", "Asia/Yangon"}
	case "LIX":
		return Location{"Likoma Island", "Africa/Blantyre"}
	case "LIY":
		return Location{"Fort Stewart(Hinesville)", "America/New_York"}
	case "LJA":
		return Location{"Lodja", "Africa/Lubumbashi"}
	case "LJG":
		return Location{"Lijiang", "Asia/Shanghai"}
	case "LJN":
		return Location{"Angleton/Lake Jackson", "America/Chicago"}
	case "LJU":
		return Location{"Ljubljana", "Europe/Ljubljana"}
	case "LKA":
		return Location{"Larantuka-Flores Island", "Asia/Makassar"}
	case "LKB":
		return Location{"Lakeba Island", "Pacific/Fiji"}
	case "LKD":
		return Location{"", "Australia/Brisbane"}
	case "LKG":
		return Location{"Lokichoggio", "Africa/Nairobi"}
	case "LKH":
		return Location{"Long Akah", "Asia/Kuching"}
	case "LKK":
		return Location{"Kulik Lake", "America/Anchorage"}
	case "LKL":
		return Location{"Lakselv", "Europe/Oslo"}
	case "LKN":
		return Location{"Leknes", "Europe/Oslo"}
	case "LKO":
		return Location{"Lucknow", "Asia/Kolkata"}
	case "LKP":
		return Location{"Lake Placid", "America/New_York"}
	case "LKV":
		return Location{"Lakeview", "America/Los_Angeles"}
	case "LKW":
		return Location{"", "Asia/Muscat"}
	case "LKY":
		return Location{"Lake Manyara National Park", "Africa/Dar_es_Salaam"}
	case "LKZ":
		return Location{"Lakenheath", "Europe/London"}
	case "LLA":
		return Location{"Lulea", "Europe/Stockholm"}
	case "LLB":
		return Location{"", "Asia/Shanghai"}
	case "LLE":
		return Location{"Malelane", "Africa/Johannesburg"}
	case "LLF":
		return Location{"Yongzhou", "Asia/Shanghai"}
	case "LLG":
		return Location{"", "Australia/Brisbane"}
	case "LLI":
		return Location{"Lalibela", "Africa/Addis_Ababa"}
	case "LLJ":
		return Location{"Lalmonirhat", "Asia/Dhaka"}
	case "LLK":
		return Location{"Lankaran", "Asia/Baku"}
	case "LLS":
		return Location{"Las Lomitas", "America/Argentina/Cordoba"}
	case "LLT":
		return Location{"Lobito", "Africa/Luanda"}
	case "LLW":
		return Location{"Lilongwe", "Africa/Blantyre"}
	case "LLX":
		return Location{"Lyndonville", "America/New_York"}
	case "LLY":
		return Location{"Mount Holly", "America/New_York"}
	case "LMA":
		return Location{"Minchumina", "America/Anchorage"}
	case "LMB":
		return Location{"Salima", "Africa/Blantyre"}
	case "LME":
		return Location{"Le Mans/Arnage", "Europe/Paris"}
	case "LMI":
		return Location{"Lumi", "Pacific/Port_Moresby"}
	case "LMM":
		return Location{"Los Mochis", "America/Mazatlan"}
	case "LMN":
		return Location{"Limbang", "Asia/Brunei"}
	case "LMO":
		return Location{"Lossiemouth", "Europe/London"}
	case "LMP":
		return Location{"Lampedusa", "Europe/Rome"}
	case "LMQ":
		return Location{"", "Africa/Tripoli"}
	case "LMR":
		return Location{"Lime Acres", "Africa/Johannesburg"}
	case "LMS":
		return Location{"Louisville", "America/Chicago"}
	case "LMT":
		return Location{"Klamath Falls", "America/Los_Angeles"}
	case "LMV":
		return Location{"Naifaru", "Indian/Maldives"}
	case "LMY":
		return Location{"Lake Murray", "Pacific/Port_Moresby"}
	case "LNA":
		return Location{"West Palm Beach", "America/New_York"}
	case "LNB":
		return Location{"Lamen Bay", "Pacific/Efate"}
	case "LND":
		return Location{"Lander", "America/Denver"}
	case "LNE":
		return Location{"Lonorore", "Pacific/Efate"}
	case "LNH":
		return Location{"", "Australia/Darwin"}
	case "LNI":
		return Location{"Lonely", "America/Anchorage"}
	case "LNJ":
		return Location{"Lincang", "Asia/Shanghai"}
	case "LNK":
		return Location{"Lincoln", "America/Chicago"}
	case "LNL":
		return Location{"Longnan", "Asia/Shanghai"}
	case "LNN":
		return Location{"Willoughby", "America/New_York"}
	case "LNO":
		return Location{"Leonora", "Australia/Perth"}
	case "LNP":
		return Location{"Wise", "America/New_York"}
	case "LNR":
		return Location{"Lone Rock", "America/Chicago"}
	case "LNS":
		return Location{"Lancaster", "America/New_York"}
	case "LNV":
		return Location{"Londolovit", "Pacific/Port_Moresby"}
	case "LNX":
		return Location{"Smolensk", "Europe/Moscow"}
	case "LNY":
		return Location{"Lanai City", "Pacific/Honolulu"}
	case "LNZ":
		return Location{"Linz", "Europe/Vienna"}
	case "LOA":
		return Location{"", "Australia/Brisbane"}
	case "LOB":
		return Location{"Los Andes", "America/Santiago"}
	case "LOC":
		return Location{"Lock", "Australia/Adelaide"}
	case "LOD":
		return Location{"Longana", "Pacific/Efate"}
	case "LOE":
		return Location{"", "Asia/Bangkok"}
	case "LOH":
		return Location{"La Toma (Catamayo)", "America/Guayaquil"}
	case "LOI":
		return Location{"Lontras", "America/Sao_Paulo"}
	case "LOK":
		return Location{"Lodwar", "Africa/Nairobi"}
	case "LOL":
		return Location{"Lovelock", "America/Los_Angeles"}
	case "LOO":
		return Location{"Laghouat", "Africa/Algiers"}
	case "LOP":
		return Location{"Mataram", "Asia/Makassar"}
	case "LOS":
		return Location{"Lagos", "Africa/Lagos"}
	case "LOT":
		return Location{"Chicago/Romeoville", "America/Chicago"}
	case "LOU":
		return Location{"Louisville", "America/Kentucky/Louisville"}
	case "LOV":
		return Location{"", "America/Monterrey"}
	case "LOW":
		return Location{"Louisa", "America/New_York"}
	case "LOY":
		return Location{"Loyengalani", "Africa/Nairobi"}
	case "LOZ":
		return Location{"London", "America/New_York"}
	case "LPA":
		return Location{"Gran Canaria Island", "Atlantic/Canary"}
	case "LPB":
		return Location{"La Paz / El Alto", "America/La_Paz"}
	case "LPC":
		return Location{"Lompoc", "America/Los_Angeles"}
	case "LPD":
		return Location{"La Pedrera", "America/Bogota"}
	case "LPF":
		return Location{"Liupanshui", "Asia/Shanghai"}
	case "LPG":
		return Location{"La Plata", "America/Argentina/Buenos_Aires"}
	case "LPI":
		return Location{"Linkoping", "Europe/Stockholm"}
	case "LPJ":
		return Location{"Guayabal", "America/Caracas"}
	case "LPK":
		return Location{"Lipetsk", "Europe/Moscow"}
	case "LPL":
		return Location{"Liverpool", "Europe/London"}
	case "LPM":
		return Location{"Lamap", "Pacific/Efate"}
	case "LPO":
		return Location{"La Porte", "America/Chicago"}
	case "LPP":
		return Location{"Lappeenranta", "Europe/Helsinki"}
	case "LPQ":
		return Location{"Luang Phabang", "Asia/Vientiane"}
	case "LPT":
		return Location{"", "Asia/Bangkok"}
	case "LPU":
		return Location{"Long Apung-Borneo Island", "Asia/Makassar"}
	case "LPX":
		return Location{"Liepaja", "Europe/Riga"}
	case "LPY":
		return Location{"Le Puy/Loudes", "Europe/Paris"}
	case "LQK":
		return Location{"Pickens", "America/New_York"}
	case "LQM":
		return Location{"Puerto Leguizamo", "America/Bogota"}
	case "LQN":
		return Location{"Qala-I-Naw", "Asia/Kabul"}
	case "LRA":
		return Location{"Larisa", "Europe/Athens"}
	case "LRB":
		return Location{"Leribe", "Africa/Johannesburg"}
	case "LRD":
		return Location{"Laredo", "America/Chicago"}
	case "LRE":
		return Location{"Longreach", "Australia/Brisbane"}
	case "LRF":
		return Location{"Jacksonville", "America/Chicago"}
	case "LRG":
		return Location{"Loralai", "Asia/Karachi"}
	case "LRH":
		return Location{"La Rochelle/Ile de Re", "Europe/Paris"}
	case "LRJ":
		return Location{"Le Mars", "America/Chicago"}
	case "LRL":
		return Location{"Niamtougou", "Africa/Lome"}
	case "LRM":
		return Location{"La Romana", "America/Santo_Domingo"}
	case "LRR":
		return Location{"Lar", "Asia/Tehran"}
	case "LRS":
		return Location{"Leros Island", "Europe/Athens"}
	case "LRT":
		return Location{"Lorient/Lann/Bihoue", "Europe/Paris"}
	case "LRU":
		return Location{"Las Cruces", "America/Denver"}
	case "LRV":
		return Location{"Los Roques", "America/Caracas"}
	case "LSA":
		return Location{"Losuia", "Pacific/Port_Moresby"}
	case "LSB":
		return Location{"Lordsburg", "America/Denver"}
	case "LSC":
		return Location{"La Serena-Coquimbo", "America/Santiago"}
	case "LSE":
		return Location{"La Crosse", "America/Chicago"}
	case "LSF":
		return Location{"Fort Benning(Columbus)", "America/Chicago"}
	case "LSH":
		return Location{"Lashio", "Asia/Yangon"}
	case "LSI":
		return Location{"Lerwick", "Europe/London"}
	case "LSK":
		return Location{"Lusk", "America/Denver"}
	case "LSL":
		return Location{"Los Chiles", "America/Costa_Rica"}
	case "LSM":
		return Location{"Long Semado", "Asia/Kuching"}
	case "LSN":
		return Location{"Los Banos", "America/Los_Angeles"}
	case "LSO":
		return Location{"Les Sables-d'Olonne", "Europe/Paris"}
	case "LSP":
		return Location{"Paraguana", "America/Caracas"}
	case "LSQ":
		return Location{"Los Angeles", "America/Santiago"}
	case "LSS":
		return Location{"Les Saintes", "America/Guadeloupe"}
	case "LST":
		return Location{"Launceston", "Australia/Hobart"}
	case "LSU":
		return Location{"Long Sukang", "Asia/Kuching"}
	case "LSV":
		return Location{"Las Vegas", "America/Los_Angeles"}
	case "LSW":
		return Location{"Lhok Seumawe-Sumatra Island", "Asia/Jakarta"}
	case "LSX":
		return Location{"Lhok Sukon-Sumatra Island", "Asia/Jakarta"}
	case "LSY":
		return Location{"Lismore", "Australia/Sydney"}
	case "LSZ":
		return Location{"Losinj", "Europe/Zagreb"}
	case "LTA":
		return Location{"Tzaneen", "Africa/Johannesburg"}
	case "LTC":
		return Location{"Lai", "Africa/Ndjamena"}
	case "LTD":
		return Location{"Ghadames", "Africa/Tripoli"}
	case "LTG":
		return Location{"Langtang", "Asia/Kathmandu"}
	case "LTI":
		return Location{"Altai", "Asia/Hovd"}
	case "LTK":
		return Location{"Latakia", "Asia/Damascus"}
	case "LTL":
		return Location{"Lastourville", "Africa/Libreville"}
	case "LTM":
		return Location{"Lethem", "America/Boa_Vista"}
	case "LTN":
		return Location{"London", "Europe/London"}
	case "LTO":
		return Location{"Loreto", "America/Mazatlan"}
	case "LTP":
		return Location{"Lyndhurst", "Australia/Brisbane"}
	case "LTQ":
		return Location{"Le Touquet-Paris-Plage", "Europe/Paris"}
	case "LTR":
		return Location{"Letterkenny", "Europe/Dublin"}
	case "LTS":
		return Location{"Altus", "America/Chicago"}
	case "LTT":
		return Location{"La Mole", "Europe/Paris"}
	case "LTU":
		return Location{"Latur", "Asia/Kolkata"}
	case "LTV":
		return Location{"Lotus Vale", "Australia/Brisbane"}
	case "LTX":
		return Location{"Latacunga", "America/Guayaquil"}
	case "LUA":
		return Location{"Lukla", "Asia/Kathmandu"}
	case "LUB":
		return Location{"Lumid Pau", "America/Guyana"}
	case "LUC":
		return Location{"Laucala Island", "Pacific/Fiji"}
	case "LUD":
		return Location{"Luderitz", "Africa/Windhoek"}
	case "LUE":
		return Location{"Lucenec", "Europe/Bratislava"}
	case "LUF":
		return Location{"Glendale", "America/Phoenix"}
	case "LUG":
		return Location{"Lugano", "Europe/Zurich"}
	case "LUH":
		return Location{"", "Asia/Kolkata"}
	case "LUJ":
		return Location{"Lusikisiki", "Africa/Johannesburg"}
	case "LUK":
		return Location{"Cincinnati", "America/New_York"}
	case "LUL":
		return Location{"Laurel", "America/Chicago"}
	case "LUM":
		return Location{"Luxi", "Asia/Shanghai"}
	case "LUN":
		return Location{"Lusaka", "Africa/Lusaka"}
	case "LUO":
		return Location{"Luena", "Africa/Luanda"}
	case "LUP":
		return Location{"Kalaupapa", "Pacific/Honolulu"}
	case "LUQ":
		return Location{"San Luis", "America/Argentina/San_Luis"}
	case "LUR":
		return Location{"Cape Lisburne", "America/Nome"}
	case "LUS":
		return Location{"Lusanga", "Africa/Kinshasa"}
	case "LUT":
		return Location{"", "Australia/Brisbane"}
	case "LUU":
		return Location{"", "Australia/Brisbane"}
	case "LUV":
		return Location{"Langgur-Seram Island", "Asia/Jayapura"}
	case "LUW":
		return Location{"Luwok-Celebes Island", "Asia/Makassar"}
	case "LUX":
		return Location{"Luxembourg", "Europe/Luxembourg"}
	case "LUZ":
		return Location{"Lublin", "Europe/Warsaw"}
	case "LVA":
		return Location{"Laval/Entrammes", "Europe/Paris"}
	case "LVI":
		return Location{"Livingstone", "Africa/Lusaka"}
	case "LVK":
		return Location{"Livermore", "America/Los_Angeles"}
	case "LVL":
		return Location{"Lawrenceville", "America/New_York"}
	case "LVM":
		return Location{"Livingston", "America/Denver"}
	case "LVO":
		return Location{"", "Australia/Perth"}
	case "LVP":
		return Location{"", "Asia/Tehran"}
	case "LVR":
		return Location{"Lucas do Rio Verde", "America/Cuiaba"}
	case "LVS":
		return Location{"Las Vegas", "America/Denver"}
	case "LWB":
		return Location{"Lewisburg", "America/New_York"}
	case "LWC":
		return Location{"Lawrence", "America/Chicago"}
	case "LWH":
		return Location{"", "Australia/Brisbane"}
	case "LWK":
		return Location{"Lerwick", "Europe/London"}
	case "LWL":
		return Location{"Wells", "America/Los_Angeles"}
	case "LWM":
		return Location{"Lawrence", "America/New_York"}
	case "LWN":
		return Location{"Gyumri", "Asia/Yerevan"}
	case "LWO":
		return Location{"Lviv", "Europe/Kyiv"}
	case "LWR":
		return Location{"Leeuwarden", "Europe/Amsterdam"}
	case "LWS":
		return Location{"Lewiston", "America/Los_Angeles"}
	case "LWT":
		return Location{"Lewistown", "America/Denver"}
	case "LWV":
		return Location{"Lawrenceville", "America/Chicago"}
	case "LWY":
		return Location{"Lawas", "Asia/Kuching"}
	case "LXA":
		return Location{"Lhasa", "Asia/Shanghai"}
	case "LXG":
		return Location{"Luang Namtha", "Asia/Vientiane"}
	case "LXN":
		return Location{"Lexington", "America/Chicago"}
	case "LXR":
		return Location{"Luxor", "Africa/Cairo"}
	case "LXS":
		return Location{"Limnos Island", "Europe/Athens"}
	case "LXU":
		return Location{"Lukulu", "Africa/Lusaka"}
	case "LXV":
		return Location{"Leadville", "America/Denver"}
	case "LYA":
		return Location{"Luoyang", "Asia/Shanghai"}
	case "LYB":
		return Location{"Little Cayman", "America/Cayman"}
	case "LYC":
		return Location{"", "Europe/Stockholm"}
	case "LYE":
		return Location{"Lyneham", "Europe/London"}
	case "LYG":
		return Location{"Lianyungang", "Asia/Shanghai"}
	case "LYH":
		return Location{"Lynchburg", "America/New_York"}
	case "LYI":
		return Location{"Linyi", "Asia/Shanghai"}
	case "LYN":
		return Location{"Lyon/Bron", "Europe/Paris"}
	case "LYO":
		return Location{"Lyons", "America/Chicago"}
	case "LYP":
		return Location{"Faisalabad", "Asia/Karachi"}
	case "LYR":
		return Location{"Longyearbyen", "Arctic/Longyearbyen"}
	case "LYS":
		return Location{"Lyon", "Europe/Paris"}
	case "LYU":
		return Location{"Ely", "America/Chicago"}
	case "LYX":
		return Location{"Lydd", "Europe/London"}
	case "LZA":
		return Location{"Luiza", "Africa/Lubumbashi"}
	case "LZC":
		return Location{"Lazaro Cardenas", "America/Mexico_City"}
	case "LZG":
		return Location{"Langzhong", "Asia/Shanghai"}
	case "LZH":
		return Location{"Liuzhou", "Asia/Shanghai"}
	case "LZI":
		return Location{"Luozi", "Africa/Kinshasa"}
	case "LZM":
		return Location{"Luzamba", "Africa/Luanda"}
	case "LZN":
		return Location{"Nangang Island", "Asia/Taipei"}
	case "LZO":
		return Location{"Luzhou", "Asia/Shanghai"}
	case "LZR":
		return Location{"", "Australia/Brisbane"}
	case "LZU":
		return Location{"Lawrenceville", "America/New_York"}
	case "LZY":
		return Location{"Nyingchi", "Asia/Shanghai"}
	case "MAA":
		return Location{"Chennai", "Asia/Kolkata"}
	case "MAB":
		return Location{"Maraba", "America/Belem"}
	case "MAC":
		return Location{"Macon", "America/New_York"}
	case "MAD":
		return Location{"Madrid", "Europe/Madrid"}
	case "MAE":
		return Location{"Madera", "America/Los_Angeles"}
	case "MAF":
		return Location{"Midland", "America/Chicago"}
	case "MAG":
		return Location{"Madang", "Pacific/Port_Moresby"}
	case "MAH":
		return Location{"Menorca Island", "Europe/Madrid"}
	case "MAI":
		return Location{"Mangochi", "Africa/Blantyre"}
	case "MAJ":
		return Location{"Majuro Atoll", "Pacific/Majuro"}
	case "MAK":
		return Location{"Malakal", "Africa/Juba"}
	case "MAL":
		return Location{"Mangole Island", "Asia/Jayapura"}
	case "MAM":
		return Location{"Matamoros", "America/Matamoros"}
	case "MAN":
		return Location{"Manchester", "Europe/London"}
	case "MAO":
		return Location{"Manaus", "America/Manaus"}
	case "MAQ":
		return Location{"", "Asia/Bangkok"}
	case "MAR":
		return Location{"Maracaibo", "America/Caracas"}
	case "MAS":
		return Location{"", "Pacific/Port_Moresby"}
	case "MAT":
		return Location{"Matadi", "Africa/Kinshasa"}
	case "MAU":
		return Location{"", "Pacific/Tahiti"}
	case "MAW":
		return Location{"Malden", "America/Chicago"}
	case "MAX":
		return Location{"Matam", "Africa/Dakar"}
	case "MAY":
		return Location{"Mangrove Cay", "America/Nassau"}
	case "MAZ":
		return Location{"Mayaguez", "America/Puerto_Rico"}
	case "MBA":
		return Location{"Mombasa", "Africa/Nairobi"}
	case "MBB":
		return Location{"", "Australia/Perth"}
	case "MBC":
		return Location{"M'Bigou", "Africa/Libreville"}
	case "MBD":
		return Location{"Mafeking", "Africa/Johannesburg"}
	case "MBE":
		return Location{"Monbetsu", "Asia/Tokyo"}
	case "MBF":
		return Location{"", "Australia/Melbourne"}
	case "MBG":
		return Location{"Mobridge", "America/Chicago"}
	case "MBH":
		return Location{"", "Australia/Brisbane"}
	case "MBI":
		return Location{"Mbeya", "Africa/Dar_es_Salaam"}
	case "MBJ":
		return Location{"Montego Bay", "America/Jamaica"}
	case "MBK":
		return Location{"Matupa", "America/Cuiaba"}
	case "MBL":
		return Location{"Manistee", "America/Detroit"}
	case "MBO":
		return Location{"", "Asia/Manila"}
	case "MBP":
		return Location{"Moyobamba", "America/Lima"}
	case "MBQ":
		return Location{"Mbarara", "Africa/Kampala"}
	case "MBS":
		return Location{"Saginaw", "America/Detroit"}
	case "MBT":
		return Location{"Masbate", "Asia/Manila"}
	case "MBU":
		return Location{"Mbambanakira", "Pacific/Guadalcanal"}
	case "MBW":
		return Location{"Melbourne", "Australia/Melbourne"}
	case "MBX":
		return Location{"", "Europe/Ljubljana"}
	case "MBY":
		return Location{"Moberly", "America/Chicago"}
	case "MBZ":
		return Location{"Maues", "America/Manaus"}
	case "MCA":
		return Location{"Macenta", "Africa/Conakry"}
	case "MCB":
		return Location{"Mc Comb", "America/Chicago"}
	case "MCC":
		return Location{"Sacramento", "America/Los_Angeles"}
	case "MCD":
		return Location{"Mackinac Island", "America/Detroit"}
	case "MCE":
		return Location{"Merced", "America/Los_Angeles"}
	case "MCF":
		return Location{"Tampa", "America/New_York"}
	case "MCG":
		return Location{"McGrath", "America/Anchorage"}
	case "MCH":
		return Location{"Machala", "America/Guayaquil"}
	case "MCI":
		return Location{"Kansas City", "America/Chicago"}
	case "MCJ":
		return Location{"La Mina-Maicao", "America/Bogota"}
	case "MCK":
		return Location{"Mc Cook", "America/Chicago"}
	case "MCL":
		return Location{"Mckinley Park", "America/Anchorage"}
	case "MCN":
		return Location{"Macon", "America/New_York"}
	case "MCO":
		return Location{"Orlando", "America/New_York"}
	case "MCP":
		return Location{"Macapa", "America/Belem"}
	case "MCQ":
		return Location{"Miskolc", "Europe/Budapest"}
	case "MCR":
		return Location{"Melchor de Mencos", "America/Belize"}
	case "MCS":
		return Location{"Monte Caseros", "America/Argentina/Cordoba"}
	case "MCT":
		return Location{"Muscat", "Asia/Muscat"}
	case "MCU":
		return Location{"Montlucon/Gueret", "Europe/Paris"}
	case "MCV":
		return Location{"McArthur River Mine", "Australia/Darwin"}
	case "MCW":
		return Location{"Mason City", "America/Chicago"}
	case "MCX":
		return Location{"Makhachkala", "Europe/Moscow"}
	case "MCY":
		return Location{"Maroochydore", "Australia/Brisbane"}
	case "MCZ":
		return Location{"Maceio", "America/Maceio"}
	case "MDC":
		return Location{"Manado-Celebes Island", "Asia/Makassar"}
	case "MDD":
		return Location{"Midland", "America/Chicago"}
	case "MDE":
		return Location{"Rionegro", "America/Bogota"}
	case "MDF":
		return Location{"Medford", "America/Chicago"}
	case "MDG":
		return Location{"Mudanjiang", "Asia/Shanghai"}
	case "MDH":
		return Location{"Carbondale/Murphysboro", "America/Chicago"}
	case "MDI":
		return Location{"Makurdi", "Africa/Lagos"}
	case "MDK":
		return Location{"Mbandaka", "Africa/Kinshasa"}
	case "MDL":
		return Location{"Mandalay", "Asia/Yangon"}
	case "MDN":
		return Location{"Madison", "America/Indiana/Indianapolis"}
	case "MDO":
		return Location{"Middleton Island", "America/Anchorage"}
	case "MDP":
		return Location{"Mindiptana-Papua Island", "Asia/Jayapura"}
	case "MDQ":
		return Location{"Mar del Plata", "America/Argentina/Buenos_Aires"}
	case "MDS":
		return Location{"Middle Caicos", "America/Grand_Turk"}
	case "MDT":
		return Location{"Harrisburg", "America/New_York"}
	case "MDU":
		return Location{"", "Pacific/Port_Moresby"}
	case "MDW":
		return Location{"Chicago", "America/Chicago"}
	case "MDX":
		return Location{"Mercedes", "America/Argentina/Cordoba"}
	case "MDY":
		return Location{"Sand Island", "Pacific/Midway"}
	case "MDZ":
		return Location{"Mendoza", "America/Argentina/Mendoza"}
	case "MEA":
		return Location{"Macae", "America/Sao_Paulo"}
	case "MEB":
		return Location{"", "Australia/Melbourne"}
	case "MEC":
		return Location{"Manta", "America/Guayaquil"}
	case "MED":
		return Location{"Medina", "Asia/Riyadh"}
	case "MEE":
		return Location{"Mare", "Pacific/Noumea"}
	case "MEG":
		return Location{"Malanje", "Africa/Luanda"}
	case "MEH":
		return Location{"Mehamn", "Europe/Oslo"}
	case "MEI":
		return Location{"Meridian", "America/Chicago"}
	case "MEJ":
		return Location{"Meadville", "America/New_York"}
	case "MEK":
		return Location{"Meknes", "Africa/Casablanca"}
	case "MEL":
		return Location{"Melbourne", "Australia/Melbourne"}
	case "MEM":
		return Location{"Memphis", "America/Chicago"}
	case "MEN":
		return Location{"Mende/Brenoux", "Europe/Paris"}
	case "MEO":
		return Location{"Manteo", "America/New_York"}
	case "MEP":
		return Location{"Mersing", "Asia/Kuala_Lumpur"}
	case "MEQ":
		return Location{"Peureumeue-Sumatra Island", "Asia/Jakarta"}
	case "MER":
		return Location{"Merced", "America/Los_Angeles"}
	case "MES":
		return Location{"Medan", "Asia/Jakarta"}
	case "MET":
		return Location{"Moreton", "Australia/Brisbane"}
	case "MEU":
		return Location{"Almeirim", "America/Santarem"}
	case "MEV":
		return Location{"Minden", "America/Los_Angeles"}
	case "MEW":
		return Location{"Mweka", "Africa/Lubumbashi"}
	case "MEX":
		return Location{"Mexico City", "America/Mexico_City"}
	case "MEY":
		return Location{"Meghauli", "Asia/Kathmandu"}
	case "MEZ":
		return Location{"Musina", "Africa/Johannesburg"}
	case "MFA":
		return Location{"Mafia Island", "Africa/Dar_es_Salaam"}
	case "MFC":
		return Location{"Mafeteng", "Africa/Maseru"}
	case "MFD":
		return Location{"Mansfield", "America/New_York"}
	case "MFE":
		return Location{"Mc Allen", "America/Chicago"}
	case "MFF":
		return Location{"Moanda", "Africa/Libreville"}
	case "MFG":
		return Location{"Muzaffarabad", "Asia/Karachi"}
	case "MFI":
		return Location{"Marshfield", "America/Chicago"}
	case "MFJ":
		return Location{"Moala", "Pacific/Fiji"}
	case "MFK":
		return Location{"Beigan Island", "Asia/Taipei"}
	case "MFM":
		return Location{"Taipa", "Asia/Macau"}
	case "MFN":
		return Location{"", "Pacific/Auckland"}
	case "MFO":
		return Location{"Manguna", "Pacific/Port_Moresby"}
	case "MFP":
		return Location{"", "Australia/Darwin"}
	case "MFQ":
		return Location{"Maradi", "Africa/Niamey"}
	case "MFR":
		return Location{"Medford", "America/Los_Angeles"}
	case "MFS":
		return Location{"Miraflores", "America/Bogota"}
	case "MFU":
		return Location{"Mfuwe", "Africa/Lusaka"}
	case "MFV":
		return Location{"Melfa", "America/New_York"}
	case "MFX":
		return Location{"Ajaccio", "Europe/Paris"}
	case "MGA":
		return Location{"Managua", "America/Managua"}
	case "MGB":
		return Location{"", "Australia/Adelaide"}
	case "MGC":
		return Location{"Michigan City", "America/Chicago"}
	case "MGD":
		return Location{"Magdalena", "America/La_Paz"}
	case "MGE":
		return Location{"Marietta", "America/New_York"}
	case "MGF":
		return Location{"Maringa", "America/Sao_Paulo"}
	case "MGH":
		return Location{"Margate", "Africa/Johannesburg"}
	case "MGJ":
		return Location{"Montgomery", "America/New_York"}
	case "MGK":
		return Location{"Mong Tong", "Asia/Yangon"}
	case "MGL":
		return Location{"Monchengladbach", "Europe/Berlin"}
	case "MGM":
		return Location{"Montgomery", "America/Chicago"}
	case "MGN":
		return Location{"Magangue", "America/Bogota"}
	case "MGQ":
		return Location{"Mogadishu", "Africa/Mogadishu"}
	case "MGR":
		return Location{"Moultrie", "America/New_York"}
	case "MGS":
		return Location{"Mangaia Island", "Pacific/Rarotonga"}
	case "MGT":
		return Location{"Milingimbi Island", "Australia/Darwin"}
	case "MGU":
		return Location{"Manaung", "Asia/Yangon"}
	case "MGV":
		return Location{"", "Australia/Perth"}
	case "MGW":
		return Location{"Morgantown", "America/New_York"}
	case "MGX":
		return Location{"Moabi", "Africa/Libreville"}
	case "MGY":
		return Location{"Dayton", "America/New_York"}
	case "MGZ":
		return Location{"Mkeik", "Asia/Yangon"}
	case "MHA":
		return Location{"Mahdia", "America/Guyana"}
	case "MHC":
		return Location{"Dalcahue", "America/Santiago"}
	case "MHD":
		return Location{"Mashhad", "Asia/Tehran"}
	case "MHE":
		return Location{"Mitchell", "America/Chicago"}
	case "MHG":
		return Location{"Mannheim", "Europe/Berlin"}
	case "MHH":
		return Location{"Marsh Harbour", "America/Nassau"}
	case "MHI":
		return Location{"Moucha Island", "Africa/Djibouti"}
	case "MHK":
		return Location{"Manhattan", "America/Chicago"}
	case "MHL":
		return Location{"Marshall", "America/Chicago"}
	case "MHO":
		return Location{"", "Australia/Perth"}
	case "MHP":
		return Location{"Minsk", "Europe/Minsk"}
	case "MHQ":
		return Location{"", "Europe/Mariehamn"}
	case "MHR":
		return Location{"Sacramento", "America/Los_Angeles"}
	case "MHS":
		return Location{"Dunsmuir", "America/Los_Angeles"}
	case "MHT":
		return Location{"Manchester", "America/New_York"}
	case "MHU":
		return Location{"Mount Hotham", "Australia/Melbourne"}
	case "MHV":
		return Location{"Mojave", "America/Los_Angeles"}
	case "MHW":
		return Location{"El Banado", "America/La_Paz"}
	case "MHX":
		return Location{"Manihiki Island", "Pacific/Rarotonga"}
	case "MHZ":
		return Location{"Mildenhall", "Europe/London"}
	case "MIA":
		return Location{"Miami", "America/New_York"}
	case "MIB":
		return Location{"Minot", "America/Chicago"}
	case "MIC":
		return Location{"Minneapolis", "America/Chicago"}
	case "MID":
		return Location{"Merida", "America/Merida"}
	case "MIE":
		return Location{"Muncie", "America/Indiana/Indianapolis"}
	case "MIG":
		return Location{"Mianyang", "Asia/Shanghai"}
	case "MIH":
		return Location{"Mitchell Plateau", "Australia/Perth"}
	case "MII":
		return Location{"Marilia", "America/Sao_Paulo"}
	case "MIJ":
		return Location{"Mili Island", "Pacific/Majuro"}
	case "MIK":
		return Location{"Mikkeli", "Europe/Helsinki"}
	case "MIM":
		return Location{"Merimbula", "Australia/Sydney"}
	case "MIN":
		return Location{"", "Australia/Adelaide"}
	case "MIO":
		return Location{"Miami", "America/Chicago"}
	case "MIP":
		return Location{"Mitzpe Ramon", "Asia/Jerusalem"}
	case "MIQ":
		return Location{"Omaha", "America/Chicago"}
	case "MIR":
		return Location{"Monastir", "Africa/Tunis"}
	case "MIS":
		return Location{"Misima Island", "Pacific/Port_Moresby"}
	case "MIT":
		return Location{"Shafter", "America/Los_Angeles"}
	case "MIU":
		return Location{"Maiduguri", "Africa/Lagos"}
	case "MIV":
		return Location{"Millville", "America/New_York"}
	case "MIW":
		return Location{"Marshalltown", "America/Chicago"}
	case "MJA":
		return Location{"Manja", "Indian/Antananarivo"}
	case "MJC":
		return Location{"", "Africa/Abidjan"}
	case "MJD":
		return Location{"Moenjodaro", "Asia/Karachi"}
	case "MJF":
		return Location{"", "Europe/Oslo"}
	case "MJG":
		return Location{"Mayajigua", "America/Havana"}
	case "MJI":
		return Location{"Tripoli", "Africa/Tripoli"}
	case "MJK":
		return Location{"Monkey Mia", "Australia/Perth"}
	case "MJL":
		return Location{"Mouila", "Africa/Libreville"}
	case "MJM":
		return Location{"Mbuji Mayi", "Africa/Lubumbashi"}
	case "MJN":
		return Location{"", "Indian/Antananarivo"}
	case "MJO":
		return Location{"", "Africa/Windhoek"}
	case "MJP":
		return Location{"", "Australia/Perth"}
	case "MJQ":
		return Location{"Jackson", "America/Chicago"}
	case "MJR":
		return Location{"Miramar", "America/Argentina/Buenos_Aires"}
	case "MJT":
		return Location{"Mytilene", "Europe/Athens"}
	case "MJU":
		return Location{"Mamuju-Celebes Island", "Asia/Makassar"}
	case "MJV":
		return Location{"San Javier", "Europe/Madrid"}
	case "MJX":
		return Location{"Toms River", "America/New_York"}
	case "MJZ":
		return Location{"Mirny", "Asia/Yakutsk"}
	case "MKA":
		return Location{"Marianske Lazne", "Europe/Prague"}
	case "MKB":
		return Location{"Mekambo", "Africa/Libreville"}
	case "MKC":
		return Location{"Kansas City", "America/Chicago"}
	case "MKE":
		return Location{"Milwaukee", "America/Chicago"}
	case "MKG":
		return Location{"Muskegon", "America/Detroit"}
	case "MKH":
		return Location{"Mokhotlong", "Africa/Maseru"}
	case "MKI":
		return Location{"Mboki", "Africa/Bangui"}
	case "MKJ":
		return Location{"Makoua", "Africa/Brazzaville"}
	case "MKK":
		return Location{"Kaunakakai", "Pacific/Honolulu"}
	case "MKL":
		return Location{"Jackson", "America/Chicago"}
	case "MKM":
		return Location{"Mukah", "Asia/Kuching"}
	case "MKO":
		return Location{"Muskogee", "America/Chicago"}
	case "MKP":
		return Location{"", "Pacific/Tahiti"}
	case "MKQ":
		return Location{"Merauke-Papua Island", "Asia/Jayapura"}
	case "MKR":
		return Location{"", "Australia/Perth"}
	case "MKT":
		return Location{"Mankato", "America/Chicago"}
	case "MKU":
		return Location{"Makokou", "Africa/Libreville"}
	case "MKV":
		return Location{"", "Australia/Darwin"}
	case "MKW":
		return Location{"Manokwari-Papua Island", "Asia/Jayapura"}
	case "MKY":
		return Location{"Mackay", "Australia/Brisbane"}
	case "MKZ":
		return Location{"Malacca", "Asia/Kuala_Lumpur"}
	case "MLA":
		return Location{"Luqa", "Europe/Malta"}
	case "MLB":
		return Location{"Melbourne", "America/New_York"}
	case "MLC":
		return Location{"Mc Alester", "America/Chicago"}
	case "MLD":
		return Location{"Malad City", "America/Boise"}
	case "MLE":
		return Location{"Male", "Indian/Maldives"}
	case "MLF":
		return Location{"Milford", "America/Denver"}
	case "MLG":
		return Location{"Malang-Java Island", "Asia/Jakarta"}
	case "MLH":
		return Location{"Bale/Mulhouse", "Europe/Paris"}
	case "MLI":
		return Location{"Moline", "America/Chicago"}
	case "MLJ":
		return Location{"Milledgeville", "America/New_York"}
	case "MLL":
		return Location{"Marshall", "America/Anchorage"}
	case "MLM":
		return Location{"Morelia", "America/Mexico_City"}
	case "MLN":
		return Location{"Melilla Island", "Africa/Casablanca"}
	case "MLO":
		return Location{"Milos Island", "Europe/Athens"}
	case "MLP":
		return Location{"Malabang", "Asia/Manila"}
	case "MLR":
		return Location{"", "Australia/Adelaide"}
	case "MLS":
		return Location{"Miles City", "America/Denver"}
	case "MLT":
		return Location{"Millinocket", "America/New_York"}
	case "MLU":
		return Location{"Monroe", "America/Chicago"}
	case "MLV":
		return Location{"", "Australia/Brisbane"}
	case "MLW":
		return Location{"Monrovia", "Africa/Monrovia"}
	case "MLX":
		return Location{"Malatya", "Europe/Istanbul"}
	case "MLY":
		return Location{"Manley Hot Springs", "America/Anchorage"}
	case "MLZ":
		return Location{"Melo", "America/Montevideo"}
	case "MMB":
		return Location{"Ozora", "Asia/Tokyo"}
	case "MMC":
		return Location{"Ciudad Mante", "America/Monterrey"}
	case "MMD":
		return Location{"", "Asia/Tokyo"}
	case "MME":
		return Location{"Durham", "Europe/London"}
	case "MMF":
		return Location{"Mamfe", "Africa/Douala"}
	case "MMG":
		return Location{"", "Australia/Perth"}
	case "MMH":
		return Location{"Mammoth Lakes", "America/Los_Angeles"}
	case "MMI":
		return Location{"Athens", "America/New_York"}
	case "MMJ":
		return Location{"Matsumoto", "Asia/Tokyo"}
	case "MMK":
		return Location{"Murmansk", "Europe/Moscow"}
	case "MML":
		return Location{"Marshall", "America/Chicago"}
	case "MMM":
		return Location{"", "Australia/Brisbane"}
	case "MMN":
		return Location{"Stow", "America/New_York"}
	case "MMO":
		return Location{"Vila do Maio", "Atlantic/Cape_Verde"}
	case "MMP":
		return Location{"Mompos", "America/Bogota"}
	case "MMQ":
		return Location{"Mbala", "Africa/Lusaka"}
	case "MMS":
		return Location{"Marks", "America/Chicago"}
	case "MMT":
		return Location{"Eastover", "America/New_York"}
	case "MMU":
		return Location{"Morristown", "America/New_York"}
	case "MMX":
		return Location{"Malmo", "Europe/Stockholm"}
	case "MMY":
		return Location{"Miyako City", "Asia/Tokyo"}
	case "MMZ":
		return Location{"", "Asia/Kabul"}
	case "MNA":
		return Location{"Karakelong Island", "Asia/Makassar"}
	case "MNB":
		return Location{"", "Africa/Kinshasa"}
	case "MNC":
		return Location{"Nacala", "Africa/Maputo"}
	case "MNE":
		return Location{"Mungeranie", "Australia/Adelaide"}
	case "MNF":
		return Location{"Mana Island", "Pacific/Fiji"}
	case "MNG":
		return Location{"Maningrida", "Australia/Darwin"}
	case "MNH":
		return Location{"Al Muladdah", "Asia/Muscat"}
	case "MNI":
		return Location{"Gerald's Park", "America/Montserrat"}
	case "MNJ":
		return Location{"", "Indian/Antananarivo"}
	case "MNK":
		return Location{"Maiana", "Pacific/Tarawa"}
	case "MNL":
		return Location{"Manila", "Asia/Manila"}
	case "MNM":
		return Location{"Menominee", "America/Chicago"}
	case "MNN":
		return Location{"Marion", "America/New_York"}
	case "MNO":
		return Location{"Manono", "Africa/Lubumbashi"}
	case "MNQ":
		return Location{"", "Australia/Brisbane"}
	case "MNR":
		return Location{"Mongu", "Africa/Lusaka"}
	case "MNS":
		return Location{"Mansa", "Africa/Lusaka"}
	case "MNU":
		return Location{"Mawlamyine", "Asia/Yangon"}
	case "MNW":
		return Location{"", "Australia/Darwin"}
	case "MNX":
		return Location{"Manicore", "America/Manaus"}
	case "MNY":
		return Location{"Stirling Island", "Pacific/Guadalcanal"}
	case "MNZ":
		return Location{"Manassas", "America/New_York"}
	case "MOA":
		return Location{"Moa", "America/Havana"}
	case "MOB":
		return Location{"Mobile", "America/Chicago"}
	case "MOC":
		return Location{"Montes Claros", "America/Sao_Paulo"}
	case "MOD":
		return Location{"Modesto", "America/Los_Angeles"}
	case "MOE":
		return Location{"", "Asia/Yangon"}
	case "MOF":
		return Location{"Maumere-Flores Island", "Asia/Makassar"}
	case "MOG":
		return Location{"Mong Hsat", "Asia/Yangon"}
	case "MOI":
		return Location{"Mitiaro Island", "Pacific/Rarotonga"}
	case "MOJ":
		return Location{"Moengo", "America/Paramaribo"}
	case "MOK":
		return Location{"Muynak", "Asia/Samarkand"}
	case "MOL":
		return Location{"Aro", "Europe/Oslo"}
	case "MOM":
		return Location{"Moudjeria", "Africa/Nouakchott"}
	case "MON":
		return Location{"", "Pacific/Auckland"}
	case "MOO":
		return Location{"", "Australia/Adelaide"}
	case "MOP":
		return Location{"Mount Pleasant", "America/Detroit"}
	case "MOQ":
		return Location{"", "Indian/Antananarivo"}
	case "MOR":
		return Location{"Morristown", "America/New_York"}
	case "MOT":
		return Location{"Minot", "America/Chicago"}
	case "MOU":
		return Location{"Mountain Village", "America/Nome"}
	case "MOV":
		return Location{"Moranbah", "Australia/Brisbane"}
	case "MOX":
		return Location{"Morris", "America/Chicago"}
	case "MOZ":
		return Location{"", "Pacific/Tahiti"}
	case "MPA":
		return Location{"Mpacha", "Africa/Windhoek"}
	case "MPC":
		return Location{"Muko Muko-Sumatra Island", "Asia/Jakarta"}
	case "MPD":
		return Location{"Sindhri", "Asia/Karachi"}
	case "MPH":
		return Location{"Malay", "Asia/Manila"}
	case "MPJ":
		return Location{"Morrilton", "America/Chicago"}
	case "MPL":
		return Location{"Montpellier/Mediterranee", "Europe/Paris"}
	case "MPM":
		return Location{"Maputo", "Africa/Maputo"}
	case "MPN":
		return Location{"Mount Pleasant", "Atlantic/Stanley"}
	case "MPO":
		return Location{"Mount Pocono", "America/New_York"}
	case "MPR":
		return Location{"Mc Pherson", "America/Chicago"}
	case "MPS":
		return Location{"Mount Pleasant", "America/Chicago"}
	case "MPT":
		return Location{"Maliana", "Asia/Dili"}
	case "MPV":
		return Location{"Barre/Montpelier", "America/New_York"}
	case "MPW":
		return Location{"Mariupol", "Europe/Kyiv"}
	case "MPY":
		return Location{"Maripasoula", "America/Cayenne"}
	case "MPZ":
		return Location{"Mount Pleasant", "America/Chicago"}
	case "MQA":
		return Location{"Mandora", "Australia/Perth"}
	case "MQB":
		return Location{"Macomb", "America/Chicago"}
	case "MQC":
		return Location{"Miquelon", "America/Miquelon"}
	case "MQD":
		return Location{"Maquinchao", "America/Argentina/Salta"}
	case "MQE":
		return Location{"Marqua", "Australia/Darwin"}
	case "MQF":
		return Location{"Magnitogorsk", "Asia/Yekaterinburg"}
	case "MQH":
		return Location{"Minacu", "America/Sao_Paulo"}
	case "MQJ":
		return Location{"Honuu", "Asia/Srednekolymsk"}
	case "MQK":
		return Location{"San Matias", "America/Cuiaba"}
	case "MQL":
		return Location{"Mildura", "Australia/Melbourne"}
	case "MQM":
		return Location{"Mardin", "Europe/Istanbul"}
	case "MQN":
		return Location{"Mo i Rana", "Europe/Oslo"}
	case "MQP":
		return Location{"Mpumalanga", "Africa/Johannesburg"}
	case "MQQ":
		return Location{"", "Africa/Ndjamena"}
	case "MQS":
		return Location{"Mustique Island", "America/St_Vincent"}
	case "MQT":
		return Location{"Marquette", "America/Detroit"}
	case "MQU":
		return Location{"Mariquita", "America/Bogota"}
	case "MQW":
		return Location{"Mc Rae", "America/New_York"}
	case "MQX":
		return Location{"", "Africa/Addis_Ababa"}
	case "MQY":
		return Location{"Smyrna", "America/Chicago"}
	case "MQZ":
		return Location{"", "Australia/Perth"}
	case "MRB":
		return Location{"Martinsburg", "America/New_York"}
	case "MRC":
		return Location{"Columbia/Mount Pleasant", "America/Chicago"}
	case "MRD":
		return Location{"Merida", "America/Caracas"}
	case "MRE":
		return Location{"Masai Mara", "Africa/Nairobi"}
	case "MRF":
		return Location{"Marfa", "America/Chicago"}
	case "MRG":
		return Location{"", "Australia/Brisbane"}
	case "MRI":
		return Location{"Anchorage", "America/Anchorage"}
	case "MRK":
		return Location{"Marco Island", "America/New_York"}
	case "MRN":
		return Location{"Morganton", "America/New_York"}
	case "MRO":
		return Location{"Masterton", "Pacific/Auckland"}
	case "MRP":
		return Location{"", "Australia/Adelaide"}
	case "MRQ":
		return Location{"Gasan", "Asia/Manila"}
	case "MRR":
		return Location{"Macara", "America/Guayaquil"}
	case "MRS":
		return Location{"Marseille", "Europe/Paris"}
	case "MRU":
		return Location{"Port Louis", "Indian/Mauritius"}
	case "MRV":
		return Location{"Mineralnyye Vody", "Europe/Moscow"}
	case "MRW":
		return Location{"Lolland Falster / Maribo", "Europe/Copenhagen"}
	case "MRX":
		return Location{"", "Asia/Tehran"}
	case "MRY":
		return Location{"Monterey", "America/Los_Angeles"}
	case "MRZ":
		return Location{"Moree", "Australia/Sydney"}
	case "MSA":
		return Location{"Muskrat Dam", "America/Winnipeg"}
	case "MSC":
		return Location{"Mesa", "America/Phoenix"}
	case "MSE":
		return Location{"Manston", "Europe/London"}
	case "MSF":
		return Location{"", "Australia/Darwin"}
	case "MSG":
		return Location{"Matsaile", "Africa/Maseru"}
	case "MSH":
		return Location{"Masirah", "Asia/Muscat"}
	case "MSI":
		return Location{"Morro Agudo", "America/Sao_Paulo"}
	case "MSJ":
		return Location{"Misawa", "Asia/Tokyo"}
	case "MSL":
		return Location{"Muscle Shoals", "America/Chicago"}
	case "MSM":
		return Location{"Masi Manimba", "Africa/Kinshasa"}
	case "MSN":
		return Location{"Madison", "America/Chicago"}
	case "MSO":
		return Location{"Missoula", "America/Denver"}
	case "MSP":
		return Location{"Minneapolis", "America/Chicago"}
	case "MSQ":
		return Location{"Minsk", "Europe/Minsk"}
	case "MSR":
		return Location{"Mus", "Europe/Istanbul"}
	case "MSS":
		return Location{"Massena", "America/New_York"}
	case "MST":
		return Location{"Maastricht", "Europe/Amsterdam"}
	case "MSU":
		return Location{"Maseru", "Africa/Maseru"}
	case "MSV":
		return Location{"Monticello", "America/New_York"}
	case "MSW":
		return Location{"Massawa", "Africa/Asmara"}
	case "MSX":
		return Location{"Mossendjo", "Africa/Brazzaville"}
	case "MSY":
		return Location{"New Orleans", "America/Chicago"}
	case "MSZ":
		return Location{"Namibe", "Africa/Luanda"}
	case "MTA":
		return Location{"", "Pacific/Auckland"}
	case "MTB":
		return Location{"Montelibano", "America/Bogota"}
	case "MTC":
		return Location{"Mount Clemens", "America/Detroit"}
	case "MTD":
		return Location{"", "Australia/Darwin"}
	case "MTE":
		return Location{"Monte Alegre", "America/Santarem"}
	case "MTF":
		return Location{"Mizan Teferi", "Africa/Addis_Ababa"}
	case "MTG":
		return Location{"Vila Bela Da Santissima Trindade", "America/Cuiaba"}
	case "MTH":
		return Location{"Marathon", "America/New_York"}
	case "MTI":
		return Location{"Vila do Mosteiros", "Atlantic/Cape_Verde"}
	case "MTJ":
		return Location{"Montrose", "America/Denver"}
	case "MTK":
		return Location{"Makin Island", "Pacific/Tarawa"}
	case "MTL":
		return Location{"", "Australia/Sydney"}
	case "MTN":
		return Location{"Baltimore", "America/New_York"}
	case "MTO":
		return Location{"Mattoon/Charleston", "America/Chicago"}
	case "MTP":
		return Location{"Montauk", "America/New_York"}
	case "MTQ":
		return Location{"", "Australia/Brisbane"}
	case "MTR":
		return Location{"Monteria", "America/Bogota"}
	case "MTS":
		return Location{"Manzini", "Africa/Mbabane"}
	case "MTT":
		return Location{"Minatitlan", "America/Mexico_City"}
	case "MTV":
		return Location{"Ablow", "Pacific/Efate"}
	case "MTW":
		return Location{"Manitowoc", "America/Chicago"}
	case "MTY":
		return Location{"Monterrey", "America/Monterrey"}
	case "MTZ":
		return Location{"Masada", "Asia/Amman"}
	case "MUA":
		return Location{"", "Pacific/Guadalcanal"}
	case "MUB":
		return Location{"Maun", "Africa/Gaborone"}
	case "MUC":
		return Location{"Munich", "Europe/Berlin"}
	case "MUD":
		return Location{"Mueda", "Africa/Maputo"}
	case "MUE":
		return Location{"Kamuela", "Pacific/Honolulu"}
	case "MUG":
		return Location{"Mulege", "America/Mazatlan"}
	case "MUH":
		return Location{"Mersa Matruh", "Africa/Cairo"}
	case "MUI":
		return Location{"Fort Indiantown Gap(Annville)", "America/New_York"}
	case "MUK":
		return Location{"Mauke Island", "Pacific/Rarotonga"}
	case "MUL":
		return Location{"Moultrie", "America/New_York"}
	case "MUN":
		return Location{"", "America/Caracas"}
	case "MUO":
		return Location{"Mountain Home", "America/Boise"}
	case "MUP":
		return Location{"", "Australia/Darwin"}
	case "MUQ":
		return Location{"Muccan Station", "Australia/Perth"}
	case "MUR":
		return Location{"Marudi", "Asia/Kuching"}
	case "MUS":
		return Location{"", "Asia/Tokyo"}
	case "MUT":
		return Location{"Muscatine", "America/Chicago"}
	case "MUW":
		return Location{"", "Africa/Algiers"}
	case "MUX":
		return Location{"Multan", "Asia/Karachi"}
	case "MUY":
		return Location{"Mouyondzi", "Africa/Brazzaville"}
	case "MUZ":
		return Location{"Musoma", "Africa/Dar_es_Salaam"}
	case "MVA":
		return Location{"Myvatn", "Atlantic/Reykjavik"}
	case "MVB":
		return Location{"Franceville", "Africa/Libreville"}
	case "MVC":
		return Location{"Monroeville", "America/Chicago"}
	case "MVD":
		return Location{"Montevideo", "America/Montevideo"}
	case "MVE":
		return Location{"Montevideo", "America/Chicago"}
	case "MVF":
		return Location{"Mossoro", "America/Fortaleza"}
	case "MVK":
		return Location{"Mulka", "Australia/Adelaide"}
	case "MVL":
		return Location{"Morrisville", "America/New_York"}
	case "MVN":
		return Location{"Mount Vernon", "America/Chicago"}
	case "MVO":
		return Location{"Mongo", "Africa/Ndjamena"}
	case "MVP":
		return Location{"Mitu", "America/Bogota"}
	case "MVQ":
		return Location{"Mogilev", "Europe/Minsk"}
	case "MVR":
		return Location{"Maroua", "Africa/Douala"}
	case "MVS":
		return Location{"Mucuri", "America/Bahia"}
	case "MVT":
		return Location{"", "Pacific/Tahiti"}
	case "MVU":
		return Location{"", "Australia/Brisbane"}
	case "MVV":
		return Location{"Verdun", "Europe/Paris"}
	case "MVW":
		return Location{"Burlington/Mount Vernon", "America/Los_Angeles"}
	case "MVX":
		return Location{"Minvoul", "Africa/Libreville"}
	case "MVY":
		return Location{"Martha's Vineyard", "America/New_York"}
	case "MVZ":
		return Location{"Masvingo", "Africa/Harare"}
	case "MWA":
		return Location{"Marion", "America/Chicago"}
	case "MWB":
		return Location{"", "Australia/Perth"}
	case "MWC":
		return Location{"Milwaukee", "America/Chicago"}
	case "MWD":
		return Location{"Mianwali", "Asia/Karachi"}
	case "MWE":
		return Location{"Merowe", "Africa/Khartoum"}
	case "MWF":
		return Location{"Maewo Island", "Pacific/Efate"}
	case "MWH":
		return Location{"Moses Lake", "America/Los_Angeles"}
	case "MWJ":
		return Location{"Matthews Ridge", "America/Guyana"}
	case "MWK":
		return Location{"Matak Island", "Asia/Jakarta"}
	case "MWL":
		return Location{"Mineral Wells", "America/Chicago"}
	case "MWM":
		return Location{"Windom", "America/Chicago"}
	case "MWN":
		return Location{"Mwadui", "Africa/Dar_es_Salaam"}
	case "MWO":
		return Location{"Middletown", "America/New_York"}
	case "MWQ":
		return Location{"Magway", "Asia/Yangon"}
	case "MWT":
		return Location{"", "Australia/Adelaide"}
	case "MWV":
		return Location{"Sen Monorom", "Asia/Phnom_Penh"}
	case "MWX":
		return Location{"", "Asia/Seoul"}
	case "MWY":
		return Location{"", "Australia/Sydney"}
	case "MWZ":
		return Location{"Mwanza", "Africa/Dar_es_Salaam"}
	case "MXA":
		return Location{"Manila", "America/Chicago"}
	case "MXB":
		return Location{"Masamba-Celebes Island", "Asia/Makassar"}
	case "MXD":
		return Location{"", "Australia/Brisbane"}
	case "MXE":
		return Location{"Maxton", "America/New_York"}
	case "MXF":
		return Location{"Montgomery", "America/Chicago"}
	case "MXH":
		return Location{"Moro", "Pacific/Port_Moresby"}
	case "MXI":
		return Location{"", "Asia/Manila"}
	case "MXJ":
		return Location{"Minna", "Africa/Lagos"}
	case "MXK":
		return Location{"Mindik", "Pacific/Port_Moresby"}
	case "MXL":
		return Location{"Mexicali", "America/Tijuana"}
	case "MXM":
		return Location{"", "Indian/Antananarivo"}
	case "MXN":
		return Location{"Morlaix/Ploujean", "Europe/Paris"}
	case "MXO":
		return Location{"Monticello", "America/Chicago"}
	case "MXP":
		return Location{"Milan", "Europe/Rome"}
	case "MXR":
		return Location{"Myrhorod", "Europe/Kiev"}
	case "MXS":
		return Location{"Maota", "Pacific/Apia"}
	case "MXT":
		return Location{"Maintirano", "Indian/Antananarivo"}
	case "MXU":
		return Location{"", "Australia/Perth"}
	case "MXV":
		return Location{"Moron", "Asia/Ulaanbaatar"}
	case "MXX":
		return Location{"", "Europe/Stockholm"}
	case "MXY":
		return Location{"Mccarthy", "America/Anchorage"}
	case "MXZ":
		return Location{"Meixian", "Asia/Shanghai"}
	case "MYA":
		return Location{"Moruya", "Australia/Sydney"}
	case "MYB":
		return Location{"Mayumba", "Africa/Libreville"}
	case "MYC":
		return Location{"Maracay", "America/Caracas"}
	case "MYD":
		return Location{"Malindi", "Africa/Nairobi"}
	case "MYE":
		return Location{"Miyakejima", "Asia/Tokyo"}
	case "MYF":
		return Location{"San Diego", "America/Los_Angeles"}
	case "MYG":
		return Location{"Mayaguana", "America/Nassau"}
	case "MYI":
		return Location{"Murray Island", "Australia/Brisbane"}
	case "MYJ":
		return Location{"Matsuyama", "Asia/Tokyo"}
	case "MYL":
		return Location{"Mc Call", "America/Boise"}
	case "MYM":
		return Location{"Monkey Mountain", "America/Guyana"}
	case "MYN":
		return Location{"Mareb", "Asia/Aden"}
	case "MYO":
		return Location{"", "Australia/Perth"}
	case "MYP":
		return Location{"Mary", "Asia/Ashgabat"}
	case "MYQ":
		return Location{"Mysore", "Asia/Kolkata"}
	case "MYR":
		return Location{"Myrtle Beach", "America/New_York"}
	case "MYT":
		return Location{"Myitkyina", "Asia/Yangon"}
	case "MYU":
		return Location{"Mekoryuk", "America/Nome"}
	case "MYV":
		return Location{"Marysville", "America/Los_Angeles"}
	case "MYW":
		return Location{"Mtwara", "Africa/Dar_es_Salaam"}
	case "MYX":
		return Location{"Menyamya", "Pacific/Port_Moresby"}
	case "MYY":
		return Location{"Miri", "Asia/Kuching"}
	case "MYZ":
		return Location{"Monkey Bay", "Africa/Blantyre"}
	case "MZA":
		return Location{"Mazamari", "America/Lima"}
	case "MZB":
		return Location{"Mocimboa da Praia", "Africa/Maputo"}
	case "MZC":
		return Location{"Mitzic", "Africa/Libreville"}
	case "MZG":
		return Location{"Makung City", "Asia/Taipei"}
	case "MZH":
		return Location{"Amasya", "Europe/Istanbul"}
	case "MZI":
		return Location{"", "Africa/Bamako"}
	case "MZJ":
		return Location{"Marana", "America/Phoenix"}
	case "MZK":
		return Location{"Marakei", "Pacific/Tarawa"}
	case "MZL":
		return Location{"Manizales", "America/Bogota"}
	case "MZM":
		return Location{"Metz/Frescaty", "Europe/Paris"}
	case "MZO":
		return Location{"Manzanillo", "America/Havana"}
	case "MZP":
		return Location{"", "Pacific/Auckland"}
	case "MZQ":
		return Location{"Mkuze", "Africa/Johannesburg"}
	case "MZR":
		return Location{"", "Asia/Kabul"}
	case "MZT":
		return Location{"Mazatlan", "America/Mazatlan"}
	case "MZU":
		return Location{"", "Asia/Kolkata"}
	case "MZV":
		return Location{"Mulu", "Asia/Kuching"}
	case "MZW":
		return Location{"Mecheria", "Africa/Algiers"}
	case "MZX":
		return Location{"Masslo", "Africa/Addis_Ababa"}
	case "MZY":
		return Location{"Mossel Bay", "Africa/Johannesburg"}
	case "MZZ":
		return Location{"Marion", "America/Indiana/Indianapolis"}
	case "NAA":
		return Location{"Narrabri", "Australia/Sydney"}
	case "NAC":
		return Location{"", "Australia/Adelaide"}
	case "NAE":
		return Location{"Natitingou", "Africa/Porto-Novo"}
	case "NAG":
		return Location{"Naqpur", "Asia/Kolkata"}
	case "NAH":
		return Location{"Tahuna-Sangihe Island", "Asia/Makassar"}
	case "NAI":
		return Location{"Annai", "America/Guyana"}
	case "NAJ":
		return Location{"Nakhchivan", "Asia/Baku"}
	case "NAK":
		return Location{"", "Asia/Bangkok"}
	case "NAL":
		return Location{"Nalchik", "Europe/Moscow"}
	case "NAM":
		return Location{"Namlea-Buru Island", "Asia/Jayapura"}
	case "NAN":
		return Location{"Nadi", "Pacific/Fiji"}
	case "NAO":
		return Location{"Nanchong", "Asia/Shanghai"}
	case "NAP":
		return Location{"Napoli", "Europe/Rome"}
	case "NAQ":
		return Location{"Qaanaaq", "America/Thule"}
	case "NAR":
		return Location{"Armenia", "America/Bogota"}
	case "NAS":
		return Location{"Nassau", "America/Nassau"}
	case "NAT":
		return Location{"Natal", "America/Fortaleza"}
	case "NAU":
		return Location{"Napuka Island", "Pacific/Tahiti"}
	case "NAV":
		return Location{"Nevsehir", "Europe/Istanbul"}
	case "NAW":
		return Location{"", "Asia/Bangkok"}
	case "NAY":
		return Location{"Beijing", "Asia/Shanghai"}
	case "NBC":
		return Location{"Nizhnekamsk", "Europe/Moscow"}
	case "NBE":
		return Location{"Enfidha", "Africa/Tunis"}
	case "NBG":
		return Location{"New Orleans", "America/Chicago"}
	case "NBH":
		return Location{"Nambucca Heads", "Australia/Sydney"}
	case "NBJ":
		return Location{"Luanda", "Africa/Luanda"}
	case "NBL":
		return Location{"Wannukandi", "America/Panama"}
	case "NBO":
		return Location{"Nairobi", "Africa/Nairobi"}
	case "NBS":
		return Location{"Baishan", "Asia/Shanghai"}
	case "NBW":
		return Location{"Guantanamo Bay Naval Station", "America/Havana"}
	case "NBX":
		return Location{"Nabire-Papua Island", "Asia/Jayapura"}
	case "NCA":
		return Location{"", "America/Grand_Turk"}
	case "NCE":
		return Location{"Nice", "Europe/Paris"}
	case "NCG":
		return Location{"", "America/Chihuahua"}
	case "NCH":
		return Location{"Nachingwea", "Africa/Dar_es_Salaam"}
	case "NCI":
		return Location{"Necocli", "America/Bogota"}
	case "NCJ":
		return Location{"Sunchales", "America/Argentina/Cordoba"}
	case "NCL":
		return Location{"Newcastle", "Europe/London"}
	case "NCN":
		return Location{"Chenega", "America/Anchorage"}
	case "NCO":
		return Location{"North Kingstown", "America/New_York"}
	case "NCR":
		return Location{"San Carlos", "America/Managua"}
	case "NCS":
		return Location{"Newcastle", "Africa/Johannesburg"}
	case "NCT":
		return Location{"Nicoya/Guanacate", "America/Costa_Rica"}
	case "NCU":
		return Location{"Nukus", "Asia/Samarkand"}
	case "NCY":
		return Location{"Annecy/Meythet", "Europe/Paris"}
	case "NDA":
		return Location{"Banda Island", "Asia/Jakarta"}
	case "NDB":
		return Location{"Nouadhibou", "Africa/El_Aaiun"}
	case "NDC":
		return Location{"Nanded", "Asia/Kolkata"}
	case "NDD":
		return Location{"Sumbe", "Africa/Luanda"}
	case "NDE":
		return Location{"Mandera", "Africa/Addis_Ababa"}
	case "NDG":
		return Location{"Qiqihar", "Asia/Shanghai"}
	case "NDJ":
		return Location{"N'Djamena", "Africa/Ndjamena"}
	case "NDL":
		return Location{"N'Dele", "Africa/Bangui"}
	case "NDM":
		return Location{"Mendi", "Africa/Addis_Ababa"}
	case "NDR":
		return Location{"Nador", "Africa/Casablanca"}
	case "NDS":
		return Location{"Sandstone", "Australia/Perth"}
	case "NDU":
		return Location{"Rundu", "Africa/Windhoek"}
	case "NDY":
		return Location{"Sanday", "Europe/London"}
	case "NDZ":
		return Location{"", "Europe/Berlin"}
	case "NEC":
		return Location{"Necochea", "America/Argentina/Buenos_Aires"}
	case "NEF":
		return Location{"Neftekamsk", "Asia/Yekaterinburg"}
	case "NEG":
		return Location{"Negril", "America/Jamaica"}
	case "NEJ":
		return Location{"Nejjo", "Africa/Addis_Ababa"}
	case "NEK":
		return Location{"Nekemte", "Africa/Addis_Ababa"}
	case "NEL":
		return Location{"Lakehurst", "America/New_York"}
	case "NEN":
		return Location{"Jacksonville", "America/New_York"}
	case "NER":
		return Location{"Chulman", "Asia/Yakutsk"}
	case "NEU":
		return Location{"Nong Khang", "Asia/Vientiane"}
	case "NEV":
		return Location{"Charlestown", "America/St_Kitts"}
	case "NEW":
		return Location{"New Orleans", "America/Chicago"}
	case "NFG":
		return Location{"Nefteyugansk", "Asia/Yekaterinburg"}
	case "NFL":
		return Location{"Fallon", "America/Los_Angeles"}
	case "NFR":
		return Location{"Nafurah 1", "Africa/Tripoli"}
	case "NGA":
		return Location{"", "Australia/Sydney"}
	case "NGB":
		return Location{"Ningbo", "Asia/Shanghai"}
	case "NGD":
		return Location{"Anegada", "America/Tortola"}
	case "NGE":
		return Location{"N'Gaoundere", "Africa/Douala"}
	case "NGF":
		return Location{"Kaneohe", "Pacific/Honolulu"}
	case "NGI":
		return Location{"Ngau", "Pacific/Fiji"}
	case "NGL":
		return Location{"Ngala", "Africa/Johannesburg"}
	case "NGO":
		return Location{"Tokoname", "Asia/Tokyo"}
	case "NGP":
		return Location{"Corpus Christi", "America/Chicago"}
	case "NGQ":
		return Location{"Shiquanhe", "Asia/Shanghai"}
	case "NGS":
		return Location{"Nagasaki", "Asia/Tokyo"}
	case "NGU":
		return Location{"Norfolk", "America/New_York"}
	case "NGW":
		return Location{"Corpus Christi", "America/Chicago"}
	case "NGX":
		return Location{"Ngawal", "Asia/Kathmandu"}
	case "NHA":
		return Location{"Nha Trang", "Asia/Ho_Chi_Minh"}
	case "NHD":
		return Location{"Dubai", "Asia/Dubai"}
	case "NHF":
		return Location{"New Halfa", "Africa/Khartoum"}
	case "NHK":
		return Location{"Patuxent River", "America/New_York"}
	case "NHS":
		return Location{"Nushki", "Asia/Karachi"}
	case "NHT":
		return Location{"London", "Europe/London"}
	case "NHV":
		return Location{"", "Pacific/Marquesas"}
	case "NHX":
		return Location{"Foley", "America/Chicago"}
	case "NHZ":
		return Location{"Brunswick", "America/New_York"}
	case "NIA":
		return Location{"Nimba", "Africa/Monrovia"}
	case "NIB":
		return Location{"Nikolai", "America/Anchorage"}
	case "NIF":
		return Location{"", "Australia/Perth"}
	case "NIG":
		return Location{"Nikunau", "Pacific/Tarawa"}
	case "NIM":
		return Location{"Niamey", "Africa/Niamey"}
	case "NIO":
		return Location{"Nioki", "Africa/Kinshasa"}
	case "NIP":
		return Location{"Jacksonville", "America/New_York"}
	case "NIS":
		return Location{"Simberi Island", "Pacific/Port_Moresby"}
	case "NIT":
		return Location{"Niort/Souche", "Europe/Paris"}
	case "NIX":
		return Location{"Nioro du Sahel", "Africa/Bamako"}
	case "NJA":
		return Location{"", "Asia/Tokyo"}
	case "NJC":
		return Location{"Nizhnevartovsk", "Asia/Yekaterinburg"}
	case "NJF":
		return Location{"Najaf", "Asia/Baghdad"}
	case "NJK":
		return Location{"El Centro", "America/Los_Angeles"}
	case "NKC":
		return Location{"Nouakchott", "Africa/Nouakchott"}
	case "NKG":
		return Location{"Nanjing", "Asia/Shanghai"}
	case "NKL":
		return Location{"Nkolo Fuma", "Africa/Kinshasa"}
	case "NKM":
		return Location{"Nagoya", "Asia/Tokyo"}
	case "NKS":
		return Location{"Nkongsamba", "Africa/Douala"}
	case "NKU":
		return Location{"Nkaus", "Africa/Maseru"}
	case "NKW":
		return Location{"Diego Garcia", "Indian/Chagos"}
	case "NKX":
		return Location{"San Diego", "America/Los_Angeles"}
	case "NKY":
		return Location{"Nkayi", "Africa/Brazzaville"}
	case "NLA":
		return Location{"Ndola", "Africa/Lusaka"}
	case "NLC":
		return Location{"Lemoore", "America/Los_Angeles"}
	case "NLD":
		return Location{"Nuevo Laredo", "America/Matamoros"}
	case "NLF":
		return Location{"Darnley Island", "Australia/Brisbane"}
	case "NLG":
		return Location{"Nelson Lagoon", "America/Anchorage"}
	case "NLI":
		return Location{"Nikolayevsk-na-Amure Airport", "Asia/Vladivostok"}
	case "NLK":
		return Location{"Burnt Pine", "Pacific/Norfolk"}
	case "NLL":
		return Location{"", "Australia/Perth"}
	case "NLO":
		return Location{"", "Africa/Kinshasa"}
	case "NLP":
		return Location{"Nelspruit", "Africa/Johannesburg"}
	case "NLS":
		return Location{"", "Australia/Perth"}
	case "NLU":
		return Location{"Santa Lucia", "America/Mexico_City"}
	case "NLV":
		return Location{"Nikolayev", "Europe/Kyiv"}
	case "NMA":
		return Location{"Namangan", "Asia/Tashkent"}
	case "NMB":
		return Location{"", "Asia/Kolkata"}
	case "NMC":
		return Location{"", "America/Nassau"}
	case "NME":
		return Location{"Nightmute", "America/Nome"}
	case "NMF":
		return Location{"Maafaru", "Indian/Maldives"}
	case "NML":
		return Location{"Fort McMurray", "America/Edmonton"}
	case "NMR":
		return Location{"", "Australia/Brisbane"}
	case "NMS":
		return Location{"Namsang", "Asia/Yangon"}
	case "NMT":
		return Location{"Namtu", "Asia/Yangon"}
	case "NNA":
		return Location{"", "Africa/Casablanca"}
	case "NNB":
		return Location{"Santa Ana Island", "Pacific/Guadalcanal"}
	case "NNG":
		return Location{"Nanning", "Asia/Shanghai"}
	case "NNI":
		return Location{"Namutoni", "Africa/Windhoek"}
	case "NNL":
		return Location{"Nondalton", "America/Anchorage"}
	case "NNM":
		return Location{"Naryan Mar", "Europe/Moscow"}
	case "NNR":
		return Location{"Inverin", "Europe/Dublin"}
	case "NNT":
		return Location{"", "Asia/Bangkok"}
	case "NNU":
		return Location{"Nanuque", "America/Sao_Paulo"}
	case "NNX":
		return Location{"Nunukan-Nunukan Island", "Asia/Kuching"}
	case "NNY":
		return Location{"Nanyang", "Asia/Shanghai"}
	case "NOA":
		return Location{"", "Australia/Sydney"}
	case "NOB":
		return Location{"Nicoya", "America/Costa_Rica"}
	case "NOC":
		return Location{"Charleston", "Europe/Dublin"}
	case "NOD":
		return Location{"Norddeich", "Europe/Berlin"}
	case "NOG":
		return Location{"", "America/Hermosillo"}
	case "NOJ":
		return Location{"Noyabrsk", "Asia/Yekaterinburg"}
	case "NOK":
		return Location{"Nova Xavantina", "America/Cuiaba"}
	case "NON":
		return Location{"Nonouti", "Pacific/Tarawa"}
	case "NOP":
		return Location{"Sinop", "Europe/Istanbul"}
	case "NOR":
		return Location{"Nordfjordur", "Atlantic/Reykjavik"}
	case "NOS":
		return Location{"Nosy Be", "Indian/Antananarivo"}
	case "NOT":
		return Location{"Novato", "America/Los_Angeles"}
	case "NOU":
		return Location{"Noumea", "Pacific/Noumea"}
	case "NOV":
		return Location{"Huambo", "Africa/Luanda"}
	case "NOZ":
		return Location{"Novokuznetsk", "Asia/Novokuznetsk"}
	case "NPA":
		return Location{"Pensacola", "America/Chicago"}
	case "NPE":
		return Location{"", "Pacific/Auckland"}
	case "NPL":
		return Location{"New Plymouth", "Pacific/Auckland"}
	case "NPO":
		return Location{"Nanga Pinoh-Borneo Island", "Asia/Pontianak"}
	case "NPR":
		return Location{"Novo Progresso", "America/Santarem"}
	case "NPT":
		return Location{"Newport", "America/New_York"}
	case "NQA":
		return Location{"Millington", "America/Chicago"}
	case "NQI":
		return Location{"Kingsville", "America/Chicago"}
	case "NQL":
		return Location{"Niquelandia", "America/Sao_Paulo"}
	case "NQN":
		return Location{"Neuquen", "America/Argentina/Salta"}
	case "NQT":
		return Location{"Nottingham", "Europe/London"}
	case "NQU":
		return Location{"Nuqui", "America/Bogota"}
	case "NQX":
		return Location{"Key West", "America/New_York"}
	case "NQY":
		return Location{"Newquay", "Europe/London"}
	case "NQZ":
		return Location{"Astana", "Asia/Almaty"}
	case "NRA":
		return Location{"Narrandera", "Australia/Sydney"}
	case "NRB":
		return Location{"Mayport", "America/New_York"}
	case "NRD":
		return Location{"Norderney", "Europe/Berlin"}
	case "NRE":
		return Location{"Namrole-Buru Island", "Asia/Jayapura"}
	case "NRG":
		return Location{"", "Australia/Perth"}
	case "NRK":
		return Location{"Norrkoping", "Europe/Stockholm"}
	case "NRL":
		return Location{"North Ronaldsay", "Europe/London"}
	case "NRM":
		return Location{"Nara", "Africa/Bamako"}
	case "NRN":
		return Location{"Weeze", "Europe/Amsterdam"}
	case "NRR":
		return Location{"Ceiba", "America/Puerto_Rico"}
	case "NRS":
		return Location{"Imperial Beach", "America/Los_Angeles"}
	case "NRT":
		return Location{"Tokyo", "Asia/Tokyo"}
	case "NSE":
		return Location{"Milton", "America/Chicago"}
	case "NSH":
		return Location{"", "Asia/Tehran"}
	case "NSI":
		return Location{"Yaounde", "Africa/Douala"}
	case "NSK":
		return Location{"Norilsk", "Asia/Krasnoyarsk"}
	case "NSL":
		return Location{"Slayton", "America/Chicago"}
	case "NSM":
		return Location{"", "Australia/Perth"}
	case "NSN":
		return Location{"Nelson", "Pacific/Auckland"}
	case "NSO":
		return Location{"", "Australia/Sydney"}
	case "NSP":
		return Location{"Cavite City", "Asia/Manila"}
	case "NSR":
		return Location{"Sao Raimundo Nonato", "America/Fortaleza"}
	case "NST":
		return Location{"Nakhon Si Thammarat", "Asia/Bangkok"}
	case "NSV":
		return Location{"", "Australia/Brisbane"}
	case "NSY":
		return Location{"", "Europe/Rome"}
	case "NTB":
		return Location{"", "Europe/Oslo"}
	case "NTD":
		return Location{"Point Mugu", "America/Los_Angeles"}
	case "NTE":
		return Location{"Nantes", "Europe/Paris"}
	case "NTG":
		return Location{"Nantong", "Asia/Shanghai"}
	case "NTI":
		return Location{"Bintuni-Papua Island", "Asia/Jayapura"}
	case "NTL":
		return Location{"Williamtown", "Australia/Sydney"}
	case "NTN":
		return Location{"", "Australia/Brisbane"}
	case "NTO":
		return Location{"Ponta do Sol", "Atlantic/Cape_Verde"}
	case "NTQ":
		return Location{"Wajima", "Asia/Tokyo"}
	case "NTR":
		return Location{"", "America/Monterrey"}
	case "NTT":
		return Location{"Niuatoputapu", "Pacific/Tongatapu"}
	case "NTU":
		return Location{"Virginia Beach", "America/New_York"}
	case "NTX":
		return Location{"Ranai-Natuna Besar Island", "Asia/Jakarta"}
	case "NTY":
		return Location{"Pilanesberg", "Africa/Johannesburg"}
	case "NUB":
		return Location{"", "Australia/Darwin"}
	case "NUD":
		return Location{"En Nahud", "Africa/Khartoum"}
	case "NUE":
		return Location{"Nuremberg", "Europe/Berlin"}
	case "NUI":
		return Location{"Nuiqsut", "America/Anchorage"}
	case "NUJ":
		return Location{"Hamadan", "Asia/Tehran"}
	case "NUK":
		return Location{"Nukutavake", "Pacific/Tahiti"}
	case "NUL":
		return Location{"Nulato", "America/Anchorage"}
	case "NUM":
		return Location{"Neom", "Asia/Riyadh"}
	case "NUN":
		return Location{"Pensacola", "America/Chicago"}
	case "NUQ":
		return Location{"Mountain View", "America/Los_Angeles"}
	case "NUR":
		return Location{"", "Australia/Adelaide"}
	case "NUS":
		return Location{"Norsup", "Pacific/Efate"}
	case "NUU":
		return Location{"Nakuru", "Africa/Nairobi"}
	case "NUW":
		return Location{"Oak Harbor", "America/Los_Angeles"}
	case "NUX":
		return Location{"Novy Urengoy", "Asia/Yekaterinburg"}
	case "NVA":
		return Location{"Neiva", "America/Bogota"}
	case "NVD":
		return Location{"Nevada", "America/Chicago"}
	case "NVG":
		return Location{"Nueva Guinea", "America/Managua"}
	case "NVI":
		return Location{"Navoi", "Asia/Samarkand"}
	case "NVK":
		return Location{"Narvik", "Europe/Oslo"}
	case "NVN":
		return Location{"Beckwourth", "America/Los_Angeles"}
	case "NVP":
		return Location{"Novo Aripuana", "America/Manaus"}
	case "NVS":
		return Location{"Nevers/Fourchambault", "Europe/Paris"}
	case "NVT":
		return Location{"Navegantes", "America/Sao_Paulo"}
	case "NWA":
		return Location{"", "Indian/Comoro"}
	case "NWI":
		return Location{"Norwich", "Europe/London"}
	case "NYA":
		return Location{"Nyagan", "Asia/Yekaterinburg"}
	case "NYE":
		return Location{"Nyeri", "Africa/Nairobi"}
	case "NYG":
		return Location{"Quantico", "America/New_York"}
	case "NYI":
		return Location{"Sunyani", "Africa/Accra"}
	case "NYK":
		return Location{"Nanyuki", "Africa/Nairobi"}
	case "NYM":
		return Location{"Nadym", "Asia/Yekaterinburg"}
	case "NYN":
		return Location{"", "Australia/Sydney"}
	case "NYO":
		return Location{"Stockholm / Nykoping", "Europe/Stockholm"}
	case "NYR":
		return Location{"Nyurba", "Asia/Yakutsk"}
	case "NYS":
		return Location{"New York", "America/New_York"}
	case "NYT":
		return Location{"Pyinmana", "Asia/Yangon"}
	case "NYU":
		return Location{"Nyaung U", "Asia/Yangon"}
	case "NYW":
		return Location{"Monywar", "Asia/Yangon"}
	case "NZA":
		return Location{"Nzagi", "Africa/Luanda"}
	case "NZC":
		return Location{"", "America/Lima"}
	case "NZE":
		return Location{"Nzerekore", "Africa/Conakry"}
	case "NZH":
		return Location{"Manzhouli", "Asia/Shanghai"}
	case "NZL":
		return Location{"Zalantun", "Asia/Shanghai"}
	case "NZY":
		return Location{"San Diego", "America/Los_Angeles"}
	case "Niu":
		return Location{"Angaha", "Pacific/Tongatapu"}
	case "OAG":
		return Location{"Orange", "Australia/Sydney"}
	case "OAH":
		return Location{"", "Asia/Kabul"}
	case "OAI":
		return Location{"Bagram", "Asia/Kabul"}
	case "OAJ":
		return Location{"Jacksonville", "America/New_York"}
	case "OAK":
		return Location{"Oakland", "America/Los_Angeles"}
	case "OAL":
		return Location{"Cacoal", "America/Porto_Velho"}
	case "OAM":
		return Location{"", "Pacific/Auckland"}
	case "OAN":
		return Location{"Olanchito", "America/Tegucigalpa"}
	case "OAR":
		return Location{"Marina", "America/Los_Angeles"}
	case "OAS":
		return Location{"Sharana", "Asia/Kabul"}
	case "OAX":
		return Location{"Oaxaca", "America/Mexico_City"}
	case "OAZ":
		return Location{"", "Asia/Kabul"}
	case "OBC":
		return Location{"Obock", "Africa/Djibouti"}
	case "OBE":
		return Location{"Okeechobee", "America/New_York"}
	case "OBF":
		return Location{"", "Europe/Berlin"}
	case "OBI":
		return Location{"Obidos", "America/Santarem"}
	case "OBL":
		return Location{"Zoersel", "Europe/Brussels"}
	case "OBN":
		return Location{"North Connel", "Europe/London"}
	case "OBO":
		return Location{"Obihiro", "Asia/Tokyo"}
	case "OBS":
		return Location{"Aubenas/Ardeche Meridional", "Europe/Paris"}
	case "OBU":
		return Location{"Kobuk", "America/Anchorage"}
	case "OCC":
		return Location{"Coca", "America/Guayaquil"}
	case "OCE":
		return Location{"Ocean City", "America/New_York"}
	case "OCF":
		return Location{"Ocala", "America/New_York"}
	case "OCH":
		return Location{"Nacogdoches", "America/Chicago"}
	case "OCJ":
		return Location{"Ocho Rios", "America/Jamaica"}
	case "OCM":
		return Location{"", "Australia/Perth"}
	case "OCN":
		return Location{"Oceanside", "America/Los_Angeles"}
	case "OCV":
		return Location{"Ocana", "America/Bogota"}
	case "OCW":
		return Location{"Washington", "America/New_York"}
	case "ODA":
		return Location{"Ouadda", "Africa/Bangui"}
	case "ODB":
		return Location{"Cordoba", "Europe/Madrid"}
	case "ODC":
		return Location{"Oakdale", "America/Los_Angeles"}
	case "ODD":
		return Location{"", "Australia/Adelaide"}
	case "ODE":
		return Location{"Odense", "Europe/Copenhagen"}
	case "ODH":
		return Location{"Odiham", "Europe/London"}
	case "ODJ":
		return Location{"Ouanda Djalle", "Africa/Bangui"}
	case "ODL":
		return Location{"Cordillo Downs", "Australia/Adelaide"}
	case "ODN":
		return Location{"Long Seridan", "Asia/Kuching"}
	case "ODO":
		return Location{"Bodaybo", "Asia/Irkutsk"}
	case "ODR":
		return Location{"Ord River", "Australia/Perth"}
	case "ODS":
		return Location{"Odessa", "Europe/Kyiv"}
	case "ODW":
		return Location{"Oak Harbor", "America/Los_Angeles"}
	case "ODY":
		return Location{"Oudomsay", "Asia/Vientiane"}
	case "OEC":
		return Location{"Oecussi-Ambeno", "Asia/Dili"}
	case "OEL":
		return Location{"Orel", "Europe/Moscow"}
	case "OEM":
		return Location{"Paloemeu", "America/Paramaribo"}
	case "OEO":
		return Location{"Osceola", "America/Chicago"}
	case "OER":
		return Location{"Ornskoldsvik", "Europe/Stockholm"}
	case "OES":
		return Location{"San Antonio Oeste", "America/Argentina/Salta"}
	case "OFF":
		return Location{"Omaha", "America/Chicago"}
	case "OFI":
		return Location{"Ouango Fitini", "Africa/Abidjan"}
	case "OFJ":
		return Location{"Olafsfjordur", "Atlantic/Reykjavik"}
	case "OFK":
		return Location{"Norfolk", "America/Chicago"}
	case "OFU":
		return Location{"Ofu Village", "Pacific/Pago_Pago"}
	case "OGA":
		return Location{"Ogallala", "America/Denver"}
	case "OGB":
		return Location{"Orangeburg", "America/New_York"}
	case "OGD":
		return Location{"Ogden", "America/Denver"}
	case "OGE":
		return Location{"", "Pacific/Port_Moresby"}
	case "OGG":
		return Location{"Kahului", "Pacific/Honolulu"}
	case "OGL":
		return Location{"Ogle", "America/Guyana"}
	case "OGN":
		return Location{"", "Asia/Tokyo"}
	case "OGO":
		return Location{"Abengourou", "Africa/Abidjan"}
	case "OGR":
		return Location{"Bongor", "Africa/Ndjamena"}
	case "OGS":
		return Location{"Ogdensburg", "America/New_York"}
	case "OGU":
		return Location{"Ordu", "Europe/Istanbul"}
	case "OGX":
		return Location{"Ouargla", "Africa/Algiers"}
	case "OGZ":
		return Location{"Beslan", "Europe/Moscow"}
	case "OHA":
		return Location{"", "Pacific/Auckland"}
	case "OHB":
		return Location{"Moramanga", "Indian/Antananarivo"}
	case "OHD":
		return Location{"Ohrid", "Europe/Skopje"}
	case "OHE":
		return Location{"Mohe", "Asia/Shanghai"}
	case "OHH":
		return Location{"Okha", "Asia/Sakhalin"}
	case "OHO":
		return Location{"Okhotsk", "Asia/Vladivostok"}
	case "OHR":
		return Location{"Wyk auf Fohr", "Europe/Berlin"}
	case "OHS":
		return Location{"Sohar", "Asia/Muscat"}
	case "OHT":
		return Location{"Kohat", "Asia/Karachi"}
	case "OIA":
		return Location{"Ourilandia Do Norte", "America/Belem"}
	case "OIC":
		return Location{"Norwich", "America/New_York"}
	case "OIM":
		return Location{"Izu Oshima", "Asia/Tokyo"}
	case "OIR":
		return Location{"", "Asia/Tokyo"}
	case "OIT":
		return Location{"Oita", "Asia/Tokyo"}
	case "OIZ":
		return Location{"Sidi Ahmed", "Africa/Tunis"}
	case "OJC":
		return Location{"Olathe", "America/Chicago"}
	case "OKA":
		return Location{"Naha", "Asia/Tokyo"}
	case "OKC":
		return Location{"Oklahoma City", "America/Chicago"}
	case "OKD":
		return Location{"Sapporo", "Asia/Tokyo"}
	case "OKE":
		return Location{"", "Asia/Tokyo"}
	case "OKF":
		return Location{"Okaukuejo", "Africa/Windhoek"}
	case "OKH":
		return Location{"Cottesmore", "Europe/London"}
	case "OKI":
		return Location{"Okinoshima", "Asia/Tokyo"}
	case "OKJ":
		return Location{"Okayama City", "Asia/Tokyo"}
	case "OKK":
		return Location{"Kokomo", "America/Indiana/Indianapolis"}
	case "OKL":
		return Location{"Oksibil-Papua Island", "Asia/Jayapura"}
	case "OKM":
		return Location{"Okmulgee", "America/Chicago"}
	case "OKN":
		return Location{"Okondja", "Africa/Libreville"}
	case "OKO":
		return Location{"Fussa", "Asia/Tokyo"}
	case "OKQ":
		return Location{"Okaba-Papua Island", "Asia/Jayapura"}
	case "OKR":
		return Location{"Yorke Island", "Australia/Brisbane"}
	case "OKS":
		return Location{"Oshkosh", "America/Denver"}
	case "OKT":
		return Location{"Kzyl-Yar", "Asia/Yekaterinburg"}
	case "OKU":
		return Location{"Mokuti Lodge", "Africa/Windhoek"}
	case "OKY":
		return Location{"", "Australia/Brisbane"}
	case "OLA":
		return Location{"Orland", "Europe/Oslo"}
	case "OLB":
		return Location{"Olbia", "Europe/Rome"}
	case "OLC":
		return Location{"Sao Paulo De Olivenca", "America/Manaus"}
	case "OLD":
		return Location{"Old Town", "America/New_York"}
	case "OLE":
		return Location{"Olean", "America/New_York"}
	case "OLF":
		return Location{"Wolf Point", "America/Denver"}
	case "OLI":
		return Location{"Rif", "Atlantic/Reykjavik"}
	case "OLJ":
		return Location{"Olpoi", "Pacific/Efate"}
	case "OLK":
		return Location{"Fuerte Olimpo", "America/Asuncion"}
	case "OLM":
		return Location{"Olympia", "America/Los_Angeles"}
	case "OLN":
		return Location{"Sarmiento", "America/Argentina/Catamarca"}
	case "OLO":
		return Location{"Olomouc", "Europe/Prague"}
	case "OLP":
		return Location{"Olympic Dam", "Australia/Adelaide"}
	case "OLS":
		return Location{"Nogales", "America/Phoenix"}
	case "OLU":
		return Location{"Columbus", "America/Chicago"}
	case "OLV":
		return Location{"Olive Branch", "America/Chicago"}
	case "OLY":
		return Location{"Olney-Noble", "America/Chicago"}
	case "OLZ":
		return Location{"Olyokminsk", "Asia/Yakutsk"}
	case "OMA":
		return Location{"Omaha", "America/Chicago"}
	case "OMB":
		return Location{"Omboue", "Africa/Libreville"}
	case "OMC":
		return Location{"Ormoc City", "Asia/Manila"}
	case "OMD":
		return Location{"Oranjemund", "Africa/Johannesburg"}
	case "OME":
		return Location{"Nome", "America/Nome"}
	case "OMF":
		return Location{"Mafraq", "Asia/Amman"}
	case "OMG":
		return Location{"Omega", "Africa/Windhoek"}
	case "OMH":
		return Location{"Urmia", "Asia/Tehran"}
	case "OMI":
		return Location{"Omidiyeh", "Asia/Tehran"}
	case "OMK":
		return Location{"Omak", "America/Los_Angeles"}
	case "OMM":
		return Location{"Marmul", "Asia/Muscat"}
	case "OMN":
		return Location{"Lyaylyakul", "Asia/Tashkent"}
	case "OMO":
		return Location{"Mostar", "Europe/Sarajevo"}
	case "OMR":
		return Location{"Oradea", "Europe/Bucharest"}
	case "OMS":
		return Location{"Omsk", "Asia/Omsk"}
	case "ONA":
		return Location{"Winona", "America/Chicago"}
	case "OND":
		return Location{"Ondangwa", "Africa/Windhoek"}
	case "ONG":
		return Location{"", "Australia/Brisbane"}
	case "ONI":
		return Location{"Moanamani-Papua Island", "Asia/Jayapura"}
	case "ONJ":
		return Location{"Odate", "Asia/Tokyo"}
	case "ONK":
		return Location{"Olenyok", "Asia/Yakutsk"}
	case "ONL":
		return Location{"O'Neill", "America/Chicago"}
	case "ONM":
		return Location{"Socorro", "America/Denver"}
	case "ONO":
		return Location{"Ontario", "America/Boise"}
	case "ONP":
		return Location{"Newport", "America/Los_Angeles"}
	case "ONQ":
		return Location{"Zonguldak", "Europe/Istanbul"}
	case "ONR":
		return Location{"", "Australia/Brisbane"}
	case "ONS":
		return Location{"", "Australia/Perth"}
	case "ONT":
		return Location{"Ontario", "America/Los_Angeles"}
	case "ONU":
		return Location{"Ono-I-Lau", "Pacific/Fiji"}
	case "ONX":
		return Location{"Colon", "America/Panama"}
	case "ONY":
		return Location{"Olney", "America/Chicago"}
	case "OOA":
		return Location{"Oskaloosa", "America/Chicago"}
	case "OOD":
		return Location{"Newman", "Australia/Perth"}
	case "OOK":
		return Location{"Toksook Bay", "America/Nome"}
	case "OOL":
		return Location{"Gold Coast", "Australia/Brisbane"}
	case "OOM":
		return Location{"Cooma", "Australia/Sydney"}
	case "OOR":
		return Location{"", "Australia/Brisbane"}
	case "OOT":
		return Location{"Onotoa", "Pacific/Tarawa"}
	case "OPA":
		return Location{"Kopasker", "Atlantic/Reykjavik"}
	case "OPF":
		return Location{"Miami", "America/New_York"}
	case "OPI":
		return Location{"", "Australia/Darwin"}
	case "OPL":
		return Location{"Opelousas", "America/Chicago"}
	case "OPO":
		return Location{"Porto", "Europe/Lisbon"}
	case "OPP":
		return Location{"Salinopolis", "America/Belem"}
	case "OPS":
		return Location{"Sinop", "America/Cuiaba"}
	case "OPU":
		return Location{"Balimo", "Pacific/Port_Moresby"}
	case "OQN":
		return Location{"Kokand", "Asia/Tashkent"}
	case "ORA":
		return Location{"Oran", "America/Argentina/Salta"}
	case "ORB":
		return Location{"Orebro", "Europe/Stockholm"}
	case "ORC":
		return Location{"Orocue", "America/Bogota"}
	case "ORD":
		return Location{"Chicago", "America/Chicago"}
	case "ORE":
		return Location{"Orleans/Saint-Denis-de-l'Hotel", "Europe/Paris"}
	case "ORF":
		return Location{"Norfolk", "America/New_York"}
	case "ORG":
		return Location{"Paramaribo", "America/Paramaribo"}
	case "ORH":
		return Location{"Worcester", "America/New_York"}
	case "ORJ":
		return Location{"Orinduik", "America/Boa_Vista"}
	case "ORK":
		return Location{"Cork", "Europe/Dublin"}
	case "ORL":
		return Location{"Orlando", "America/New_York"}
	case "ORM":
		return Location{"Northampton", "Europe/London"}
	case "ORN":
		return Location{"Oran", "Africa/Algiers"}
	case "ORP":
		return Location{"", "Africa/Gaborone"}
	case "ORR":
		return Location{"", "Australia/Adelaide"}
	case "ORT":
		return Location{"Northway", "America/Anchorage"}
	case "ORU":
		return Location{"Oruro", "America/La_Paz"}
	case "ORV":
		return Location{"Noorvik", "America/Anchorage"}
	case "ORW":
		return Location{"Ormara Raik", "Asia/Karachi"}
	case "ORX":
		return Location{"Oriximina", "America/Santarem"}
	case "ORY":
		return Location{"Paris", "Europe/Paris"}
	case "OSB":
		return Location{"Mosul", "Asia/Baghdad"}
	case "OSC":
		return Location{"Oscoda", "America/Detroit"}
	case "OSD":
		return Location{"Ostersund", "Europe/Stockholm"}
	case "OSE":
		return Location{"Omora", "Pacific/Port_Moresby"}
	case "OSF":
		return Location{"Moscow", "Europe/Moscow"}
	case "OSH":
		return Location{"Oshkosh", "America/Chicago"}
	case "OSI":
		return Location{"Osijek", "Europe/Zagreb"}
	case "OSK":
		return Location{"", "Europe/Stockholm"}
	case "OSL":
		return Location{"Oslo", "Europe/Oslo"}
	case "OSN":
		return Location{"", "Asia/Seoul"}
	case "OSO":
		return Location{"", "Australia/Brisbane"}
	case "OSR":
		return Location{"Ostrava", "Europe/Prague"}
	case "OSS":
		return Location{"Osh", "Asia/Bishkek"}
	case "OST":
		return Location{"Ostend", "Europe/Brussels"}
	case "OSU":
		return Location{"Columbus", "America/New_York"}
	case "OSW":
		return Location{"Orsk", "Asia/Yekaterinburg"}
	case "OSX":
		return Location{"Kosciusko", "America/Chicago"}
	case "OSY":
		return Location{"Namsos", "Europe/Oslo"}
	case "OSZ":
		return Location{"", "Europe/Warsaw"}
	case "OTC":
		return Location{"Bol", "Africa/Ndjamena"}
	case "OTG":
		return Location{"Worthington", "America/Chicago"}
	case "OTH":
		return Location{"North Bend", "America/Los_Angeles"}
	case "OTI":
		return Location{"Gotalalamo-Morotai Island", "Asia/Jayapura"}
	case "OTJ":
		return Location{"Otjiwarongo", "Africa/Windhoek"}
	case "OTK":
		return Location{"Tillamook", "America/Los_Angeles"}
	case "OTL":
		return Location{"Boutilimit", "Africa/Nouakchott"}
	case "OTM":
		return Location{"Ottumwa", "America/Chicago"}
	case "OTP":
		return Location{"Bucharest", "Europe/Bucharest"}
	case "OTR":
		return Location{"Corredores", "America/Costa_Rica"}
	case "OTU":
		return Location{"El Rhin", "America/Bogota"}
	case "OTZ":
		return Location{"Kotzebue", "America/Nome"}
	case "OUA":
		return Location{"Ouagadougou", "Africa/Ouagadougou"}
	case "OUD":
		return Location{"Oujda", "Africa/Casablanca"}
	case "OUE":
		return Location{"", "Africa/Brazzaville"}
	case "OUG":
		return Location{"Ouahigouya", "Africa/Ouagadougou"}
	case "OUH":
		return Location{"Oudtshoorn", "Africa/Johannesburg"}
	case "OUI":
		return Location{"", "Asia/Bangkok"}
	case "OUL":
		return Location{"Oulu / Oulunsalo", "Europe/Helsinki"}
	case "OUN":
		return Location{"Norman", "America/Chicago"}
	case "OUR":
		return Location{"Batouri", "Africa/Douala"}
	case "OUS":
		return Location{"Ourinhos", "America/Sao_Paulo"}
	case "OUT":
		return Location{"Bousso", "Africa/Ndjamena"}
	case "OUZ":
		return Location{"Zouerate", "Africa/Nouakchott"}
	case "OVA":
		return Location{"Bekily", "Indian/Antananarivo"}
	case "OVB":
		return Location{"Novosibirsk", "Asia/Novosibirsk"}
	case "OVD":
		return Location{"Ranon", "Europe/Madrid"}
	case "OVE":
		return Location{"Oroville", "America/Los_Angeles"}
	case "OVG":
		return Location{"Overberg", "Africa/Johannesburg"}
	case "OVL":
		return Location{"Ovalle", "America/Santiago"}
	case "OVR":
		return Location{"Olavarria", "America/Argentina/Buenos_Aires"}
	case "OVS":
		return Location{"Sovetskiy", "Asia/Yekaterinburg"}
	case "OWA":
		return Location{"Owatonna", "America/Chicago"}
	case "OWB":
		return Location{"Owensboro", "America/Chicago"}
	case "OWD":
		return Location{"Norwood", "America/New_York"}
	case "OWK":
		return Location{"Norridgewock", "America/New_York"}
	case "OXB":
		return Location{"Bissau", "Africa/Bissau"}
	case "OXC":
		return Location{"Oxford", "America/New_York"}
	case "OXD":
		return Location{"Oxford", "America/New_York"}
	case "OXF":
		return Location{"Kidlington", "Europe/London"}
	case "OXP":
		return Location{"Saint-Georges-de-l'Oyapock Airport", "America/Belem"}
	case "OXR":
		return Location{"Oxnard", "America/Los_Angeles"}
	case "OXY":
		return Location{"", "Australia/Brisbane"}
	case "OYA":
		return Location{"Goya", "America/Argentina/Cordoba"}
	case "OYE":
		return Location{"Oyem", "Africa/Libreville"}
	case "OYK":
		return Location{"Oiapoque", "America/Belem"}
	case "OYL":
		return Location{"Moyale (Lower)", "Africa/Nairobi"}
	case "OYN":
		return Location{"", "Australia/Melbourne"}
	case "OYO":
		return Location{"Tres Arroyos", "America/Argentina/Buenos_Aires"}
	case "OZA":
		return Location{"Ozona", "America/Chicago"}
	case "OZC":
		return Location{"Ozamiz City", "Asia/Manila"}
	case "OZG":
		return Location{"Zagora", "Africa/Casablanca"}
	case "OZH":
		return Location{"Zaporizhia", "Europe/Kyiv"}
	case "OZP":
		return Location{"Moron", "Europe/Madrid"}
	case "OZR":
		return Location{"Fort Rucker/Ozark", "America/Chicago"}
	case "OZZ":
		return Location{"Ouarzazate", "Africa/Casablanca"}
	case "PAA":
		return Location{"Hpa-N", "Asia/Yangon"}
	case "PAB":
		return Location{"", "Asia/Kolkata"}
	case "PAC":
		return Location{"Albrook", "America/Panama"}
	case "PAD":
		return Location{"Paderborn", "Europe/Berlin"}
	case "PAE":
		return Location{"Everett", "America/Los_Angeles"}
	case "PAF":
		return Location{"", "Africa/Kampala"}
	case "PAG":
		return Location{"Pagadian City", "Asia/Manila"}
	case "PAH":
		return Location{"Paducah", "America/Chicago"}
	case "PAJ":
		return Location{"Parachinar", "Asia/Karachi"}
	case "PAK":
		return Location{"Hanapepe", "Pacific/Honolulu"}
	case "PAM":
		return Location{"Panama City", "America/Chicago"}
	case "PAN":
		return Location{"", "Asia/Bangkok"}
	case "PAO":
		return Location{"Palo Alto", "America/Los_Angeles"}
	case "PAP":
		return Location{"Port-au-Prince", "America/Port-au-Prince"}
	case "PAQ":
		return Location{"Palmer", "America/Anchorage"}
	case "PAS":
		return Location{"Paros Island", "Europe/Athens"}
	case "PAT":
		return Location{"Patna", "Asia/Kolkata"}
	case "PAU":
		return Location{"Pauk", "Asia/Yangon"}
	case "PAV":
		return Location{"Paulo Afonso", "America/Bahia"}
	case "PAX":
		return Location{"Port-de-Paix", "America/Port-au-Prince"}
	case "PAY":
		return Location{"Pamol", "Asia/Kuching"}
	case "PAZ":
		return Location{"Poza Rica", "America/Mexico_City"}
	case "PBA":
		return Location{"Cairu", "America/Bahia"}
	case "PBB":
		return Location{"Paranaiba", "America/Campo_Grande"}
	case "PBC":
		return Location{"Puebla", "America/Mexico_City"}
	case "PBD":
		return Location{"Porbandar", "Asia/Kolkata"}
	case "PBE":
		return Location{"Puerto Berrio", "America/Bogota"}
	case "PBF":
		return Location{"Pine Bluff", "America/Chicago"}
	case "PBG":
		return Location{"Plattsburgh", "America/New_York"}
	case "PBH":
		return Location{"Paro", "Asia/Thimphu"}
	case "PBI":
		return Location{"West Palm Beach", "America/New_York"}
	case "PBJ":
		return Location{"Paama Island", "Pacific/Efate"}
	case "PBL":
		return Location{"", "America/Caracas"}
	case "PBM":
		return Location{"Zandery", "America/Paramaribo"}
	case "PBN":
		return Location{"Port Amboim", "Africa/Luanda"}
	case "PBO":
		return Location{"Paraburdoo", "Australia/Perth"}
	case "PBP":
		return Location{"Nandayure", "America/Costa_Rica"}
	case "PBQ":
		return Location{"Pimenta Bueno", "America/Porto_Velho"}
	case "PBR":
		return Location{"Puerto Barrios", "America/Guatemala"}
	case "PBU":
		return Location{"Putao", "Asia/Yangon"}
	case "PBV":
		return Location{"Porto Dos Gauchos", "America/Cuiaba"}
	case "PBX":
		return Location{"Porto Alegre Do Norte", "America/Cuiaba"}
	case "PBZ":
		return Location{"Plettenberg Bay", "Africa/Johannesburg"}
	case "PCA":
		return Location{"Portage Creek", "America/Anchorage"}
	case "PCB":
		return Location{"Jakarta", "Asia/Jakarta"}
	case "PCD":
		return Location{"Prairie Du Chien", "America/Chicago"}
	case "PCF":
		return Location{"Potchefstroom", "Africa/Johannesburg"}
	case "PCG":
		return Location{"Paso Caballos", "America/Guatemala"}
	case "PCH":
		return Location{"Palacios", "America/Tegucigalpa"}
	case "PCL":
		return Location{"Pucallpa", "America/Lima"}
	case "PCN":
		return Location{"Picton", "Pacific/Auckland"}
	case "PCO":
		return Location{"La Ribera", "America/Mazatlan"}
	case "PCP":
		return Location{"", "Africa/Sao_Tome"}
	case "PCQ":
		return Location{"Boun Neau", "Asia/Vientiane"}
	case "PCR":
		return Location{"Puerto Carreno", "America/Bogota"}
	case "PCS":
		return Location{"Picos", "America/Fortaleza"}
	case "PDA":
		return Location{"Puerto Inirida", "America/Bogota"}
	case "PDC":
		return Location{"Mueo", "Pacific/Noumea"}
	case "PDD":
		return Location{"Ponta do Ouro", "Africa/Maputo"}
	case "PDE":
		return Location{"", "Australia/Adelaide"}
	case "PDF":
		return Location{"Prado", "America/Bahia"}
	case "PDG":
		return Location{"Ketaping/Padang-Sumatra Island", "Asia/Jakarta"}
	case "PDI":
		return Location{"Pindiu", "Pacific/Port_Moresby"}
	case "PDK":
		return Location{"Atlanta", "America/New_York"}
	case "PDL":
		return Location{"Ponta Delgada", "Atlantic/Azores"}
	case "PDN":
		return Location{"Parndana", "Australia/Adelaide"}
	case "PDO":
		return Location{"Talang Gudang-Sumatra Island", "Asia/Jakarta"}
	case "PDP":
		return Location{"Punta del Este", "America/Montevideo"}
	case "PDS":
		return Location{"", "America/Matamoros"}
	case "PDT":
		return Location{"Pendleton", "America/Los_Angeles"}
	case "PDU":
		return Location{"Paysandu", "America/Montevideo"}
	case "PDV":
		return Location{"Plovdiv", "Europe/Sofia"}
	case "PDX":
		return Location{"Portland", "America/Los_Angeles"}
	case "PDZ":
		return Location{"", "America/Caracas"}
	case "PEA":
		return Location{"Ironstone", "Australia/Adelaide"}
	case "PED":
		return Location{"Pardubice", "Europe/Prague"}
	case "PEE":
		return Location{"Perm", "Asia/Yekaterinburg"}
	case "PEF":
		return Location{"Peenemunde", "Europe/Berlin"}
	case "PEG":
		return Location{"Perugia", "Europe/Rome"}
	case "PEH":
		return Location{"Pehuajo", "America/Argentina/Buenos_Aires"}
	case "PEI":
		return Location{"Pereira", "America/Bogota"}
	case "PEK":
		return Location{"Beijing", "Asia/Shanghai"}
	case "PEL":
		return Location{"Pelaneng", "Africa/Maseru"}
	case "PEM":
		return Location{"Puerto Maldonado", "America/Lima"}
	case "PEN":
		return Location{"Penang", "Asia/Kuala_Lumpur"}
	case "PEQ":
		return Location{"Pecos", "America/Chicago"}
	case "PER":
		return Location{"Perth", "Australia/Perth"}
	case "PES":
		return Location{"Petrozavodsk", "Europe/Moscow"}
	case "PET":
		return Location{"Pelotas", "America/Sao_Paulo"}
	case "PEU":
		return Location{"Puerto Lempira", "America/Tegucigalpa"}
	case "PEV":
		return Location{"Pecs-Pogany", "Europe/Budapest"}
	case "PEW":
		return Location{"Peshawar", "Asia/Karachi"}
	case "PEX":
		return Location{"Pechora", "Europe/Moscow"}
	case "PEZ":
		return Location{"Penza", "Europe/Moscow"}
	case "PFB":
		return Location{"Passo Fundo", "America/Sao_Paulo"}
	case "PFC":
		return Location{"Pacific City", "America/Los_Angeles"}
	case "PFJ":
		return Location{"Patreksfjordur", "Atlantic/Reykjavik"}
	case "PFN":
		return Location{"Panama City", "America/Chicago"}
	case "PFO":
		return Location{"Paphos", "Asia/Nicosia"}
	case "PFQ":
		return Location{"Parsabad", "Asia/Tehran"}
	case "PFR":
		return Location{"Ilebo", "Africa/Lubumbashi"}
	case "PGA":
		return Location{"Page", "America/Phoenix"}
	case "PGC":
		return Location{"Petersburg", "America/New_York"}
	case "PGD":
		return Location{"Punta Gorda", "America/New_York"}
	case "PGF":
		return Location{"Perpignan/Rivesaltes", "Europe/Paris"}
	case "PGH":
		return Location{"Pantnagar", "Asia/Kolkata"}
	case "PGI":
		return Location{"Chitato", "Africa/Luanda"}
	case "PGK":
		return Location{"Pangkal Pinang-Palaubangka Island", "Asia/Jakarta"}
	case "PGL":
		return Location{"Pascagoula", "America/Chicago"}
	case "PGO":
		return Location{"Pagosa Springs", "America/Denver"}
	case "PGR":
		return Location{"Paragould", "America/Chicago"}
	case "PGU":
		return Location{"Asalouyeh", "Asia/Tehran"}
	case "PGV":
		return Location{"Greenville", "America/New_York"}
	case "PGX":
		return Location{"Perigueux/Bassillac", "Europe/Paris"}
	case "PGZ":
		return Location{"Ponta Grossa", "America/Sao_Paulo"}
	case "PHA":
		return Location{"Phan Rang", "Asia/Ho_Chi_Minh"}
	case "PHB":
		return Location{"Parnaiba", "America/Fortaleza"}
	case "PHC":
		return Location{"Port Harcourt", "Africa/Lagos"}
	case "PHD":
		return Location{"New Philadelphia", "America/New_York"}
	case "PHE":
		return Location{"Port Hedland", "Australia/Perth"}
	case "PHF":
		return Location{"Newport News", "America/New_York"}
	case "PHH":
		return Location{"Pokhara", "Asia/Kathmandu"}
	case "PHI":
		return Location{"Pinheiro", "America/Fortaleza"}
	case "PHK":
		return Location{"Pahokee", "America/New_York"}
	case "PHL":
		return Location{"Philadelphia", "America/New_York"}
	case "PHN":
		return Location{"Port Huron", "America/Detroit"}
	case "PHO":
		return Location{"Point Hope", "America/Nome"}
	case "PHP":
		return Location{"Philip", "America/Denver"}
	case "PHQ":
		return Location{"", "Australia/Brisbane"}
	case "PHS":
		return Location{"", "Asia/Bangkok"}
	case "PHT":
		return Location{"Paris", "America/Chicago"}
	case "PHW":
		return Location{"Phalaborwa", "Africa/Johannesburg"}
	case "PHX":
		return Location{"Phoenix", "America/Phoenix"}
	case "PHY":
		return Location{"", "Asia/Bangkok"}
	case "PIA":
		return Location{"Peoria", "America/Chicago"}
	case "PIB":
		return Location{"Hattiesburg/Laurel", "America/Chicago"}
	case "PIC":
		return Location{"Pine Cay", "America/Grand_Turk"}
	case "PID":
		return Location{"Nassau", "America/Nassau"}
	case "PIE":
		return Location{"St Petersburg-Clearwater", "America/New_York"}
	case "PIF":
		return Location{"Pingtung", "Asia/Taipei"}
	case "PIH":
		return Location{"Pocatello", "America/Boise"}
	case "PIK":
		return Location{"Glasgow", "Europe/London"}
	case "PIL":
		return Location{"Pilar", "America/Argentina/Cordoba"}
	case "PIM":
		return Location{"Pine Mountain", "America/New_York"}
	case "PIN":
		return Location{"Parintins", "America/Manaus"}
	case "PIO":
		return Location{"Pisco", "America/Lima"}
	case "PIP":
		return Location{"Pilot Point", "America/Anchorage"}
	case "PIR":
		return Location{"Pierre", "America/Chicago"}
	case "PIS":
		return Location{"Poitiers/Biard", "Europe/Paris"}
	case "PIT":
		return Location{"Pittsburgh", "America/New_York"}
	case "PIU":
		return Location{"Piura", "America/Lima"}
	case "PIV":
		return Location{"Pirapora", "America/Sao_Paulo"}
	case "PIW":
		return Location{"Pikwitonei", "America/Winnipeg"}
	case "PIX":
		return Location{"Pico Island", "Atlantic/Azores"}
	case "PIZ":
		return Location{"Point Lay", "America/Nome"}
	case "PJA":
		return Location{"", "Europe/Stockholm"}
	case "PJB":
		return Location{"Payson", "America/Phoenix"}
	case "PJC":
		return Location{"Pedro Juan Caballero", "America/Asuncion"}
	case "PJG":
		return Location{"Panjgur", "Asia/Karachi"}
	case "PJM":
		return Location{"Puerto Jimenez", "America/Costa_Rica"}
	case "PKA":
		return Location{"Napaskiak", "America/Anchorage"}
	case "PKB":
		return Location{"Parkersburg", "America/New_York"}
	case "PKC":
		return Location{"Petropavlovsk-Kamchatsky", "Asia/Kamchatka"}
	case "PKD":
		return Location{"Park Rapids", "America/Chicago"}
	case "PKE":
		return Location{"Parkes", "Australia/Sydney"}
	case "PKF":
		return Location{"Park Falls", "America/Chicago"}
	case "PKG":
		return Location{"Pangkor Island", "Asia/Kuala_Lumpur"}
	case "PKH":
		return Location{"Porto Cheli", "Europe/Athens"}
	case "PKJ":
		return Location{"Playa Grande", "America/Guatemala"}
	case "PKK":
		return Location{"Pakhokku", "Asia/Yangon"}
	case "PKN":
		return Location{"Pangkalanbun-Borneo Island", "Asia/Pontianak"}
	case "PKO":
		return Location{"Parakou", "Africa/Porto-Novo"}
	case "PKP":
		return Location{"", "Pacific/Tahiti"}
	case "PKR":
		return Location{"Pokhara", "Asia/Kathmandu"}
	case "PKT":
		return Location{"", "Australia/Darwin"}
	case "PKU":
		return Location{"Pekanbaru-Sumatra Island", "Asia/Jakarta"}
	case "PKV":
		return Location{"Pskov", "Europe/Moscow"}
	case "PKW":
		return Location{"", "Africa/Gaborone"}
	case "PKX":
		return Location{"Beijing", "Asia/Shanghai"}
	case "PKY":
		return Location{"Palangkaraya-Kalimantan Tengah", "Asia/Pontianak"}
	case "PKZ":
		return Location{"Pakse", "Asia/Vientiane"}
	case "PLD":
		return Location{"Playa Samara", "America/Costa_Rica"}
	case "PLF":
		return Location{"Pala", "Africa/Ndjamena"}
	case "PLH":
		return Location{"Plymouth", "Europe/London"}
	case "PLK":
		return Location{"Point Lookout", "America/Chicago"}
	case "PLL":
		return Location{"Manaus", "America/Manaus"}
	case "PLM":
		return Location{"Palembang-Sumatra Island", "Asia/Jakarta"}
	case "PLN":
		return Location{"Pellston", "America/Detroit"}
	case "PLO":
		return Location{"Port Lincoln", "Australia/Adelaide"}
	case "PLP":
		return Location{"La Palma", "America/Panama"}
	case "PLQ":
		return Location{"Palanga", "Europe/Vilnius"}
	case "PLR":
		return Location{"Pell City", "America/Chicago"}
	case "PLS":
		return Location{"Providenciales Island", "America/Grand_Turk"}
	case "PLT":
		return Location{"Plato", "America/Bogota"}
	case "PLU":
		return Location{"Belo Horizonte", "America/Sao_Paulo"}
	case "PLV":
		return Location{"Poltava", "Europe/Kyiv"}
	case "PLW":
		return Location{"Palu-Celebes Island", "Asia/Makassar"}
	case "PLX":
		return Location{"Semey", "Asia/Almaty"}
	case "PLZ":
		return Location{"Port Elizabeth", "Africa/Johannesburg"}
	case "PMA":
		return Location{"Chake", "Africa/Dar_es_Salaam"}
	case "PMB":
		return Location{"Pembina", "America/Chicago"}
	case "PMC":
		return Location{"Puerto Montt", "America/Santiago"}
	case "PMD":
		return Location{"Palmdale", "America/Los_Angeles"}
	case "PMF":
		return Location{"Parma", "Europe/Rome"}
	case "PMG":
		return Location{"Ponta Pora", "America/Asuncion"}
	case "PMH":
		return Location{"Portsmouth", "America/New_York"}
	case "PMI":
		return Location{"Palma De Mallorca", "Europe/Madrid"}
	case "PMK":
		return Location{"", "Australia/Brisbane"}
	case "PML":
		return Location{"Cold Bay", "America/Anchorage"}
	case "PMO":
		return Location{"Palermo", "Europe/Rome"}
	case "PMQ":
		return Location{"Perito Moreno", "America/Argentina/Rio_Gallegos"}
	case "PMR":
		return Location{"", "Pacific/Auckland"}
	case "PMS":
		return Location{"", "Asia/Damascus"}
	case "PMV":
		return Location{"Isla Margarita", "America/Caracas"}
	case "PMW":
		return Location{"Palmas", "America/Araguaina"}
	case "PMY":
		return Location{"Puerto Madryn", "America/Argentina/Catamarca"}
	case "PMZ":
		return Location{"Palmar Sur", "America/Costa_Rica"}
	case "PNA":
		return Location{"Pamplona", "Europe/Madrid"}
	case "PNB":
		return Location{"Porto Nacional", "America/Araguaina"}
	case "PNC":
		return Location{"Ponca City", "America/Chicago"}
	case "PNE":
		return Location{"Philadelphia", "America/New_York"}
	case "PNG":
		return Location{"Paranagua", "America/Sao_Paulo"}
	case "PNH":
		return Location{"Phnom Penh", "Asia/Phnom_Penh"}
	case "PNI":
		return Location{"Pohnpei Island", "Pacific/Pohnpei"}
	case "PNK":
		return Location{"Pontianak-Borneo Island", "Asia/Pontianak"}
	case "PNL":
		return Location{"Pantelleria", "Europe/Rome"}
	case "PNN":
		return Location{"Princeton", "America/New_York"}
	case "PNP":
		return Location{"Popondetta", "Pacific/Port_Moresby"}
	case "PNQ":
		return Location{"Pune", "Asia/Kolkata"}
	case "PNR":
		return Location{"Pointe Noire", "Africa/Brazzaville"}
	case "PNS":
		return Location{"Pensacola", "America/Chicago"}
	case "PNT":
		return Location{"Puerto Natales", "America/Punta_Arenas"}
	case "PNV":
		return Location{"Panevezys", "Europe/Vilnius"}
	case "PNX":
		return Location{"Sherman/Denison", "America/Chicago"}
	case "PNY":
		return Location{"", "Asia/Kolkata"}
	case "PNZ":
		return Location{"Petrolina", "America/Recife"}
	case "POA":
		return Location{"Porto Alegre", "America/Sao_Paulo"}
	case "POB":
		return Location{"Fayetteville", "America/New_York"}
	case "POC":
		return Location{"La Verne", "America/Los_Angeles"}
	case "POD":
		return Location{"Podor", "Africa/Dakar"}
	case "POE":
		return Location{"Fort Polk", "America/Chicago"}
	case "POF":
		return Location{"Poplar Bluff", "America/Chicago"}
	case "POG":
		return Location{"Port Gentil", "Africa/Libreville"}
	case "POH":
		return Location{"Pocahontas", "America/Chicago"}
	case "POI":
		return Location{"Potosi", "America/La_Paz"}
	case "POJ":
		return Location{"Patos De Minas", "America/Sao_Paulo"}
	case "POL":
		return Location{"Pemba / Porto Amelia", "Africa/Maputo"}
	case "POM":
		return Location{"Port Moresby", "Pacific/Port_Moresby"}
	case "PON":
		return Location{"Poptun", "America/Guatemala"}
	case "POO":
		return Location{"Pocos De Caldas", "America/Sao_Paulo"}
	case "POP":
		return Location{"Puerto Plata", "America/Santo_Domingo"}
	case "POR":
		return Location{"Pori", "Europe/Helsinki"}
	case "POS":
		return Location{"Port of Spain", "America/Port_of_Spain"}
	case "POT":
		return Location{"Ken Jones", "America/Jamaica"}
	case "POU":
		return Location{"Poughkeepsie", "America/New_York"}
	case "POV":
		return Location{"Presov", "Europe/Bratislava"}
	case "POW":
		return Location{"Portoroz", "Europe/Ljubljana"}
	case "POX":
		return Location{"Pontoise", "Europe/Paris"}
	case "POY":
		return Location{"Powell", "America/Denver"}
	case "POZ":
		return Location{"Poznan", "Europe/Warsaw"}
	case "PPA":
		return Location{"Pampa", "America/Chicago"}
	case "PPB":
		return Location{"Presidente Prudente", "America/Sao_Paulo"}
	case "PPC":
		return Location{"Prospect Creek", "America/Anchorage"}
	case "PPE":
		return Location{"Puerto Penasco", "America/Hermosillo"}
	case "PPF":
		return Location{"Parsons", "America/Chicago"}
	case "PPG":
		return Location{"Pago Pago", "Pacific/Pago_Pago"}
	case "PPH":
		return Location{"", "America/Caracas"}
	case "PPI":
		return Location{"", "Australia/Adelaide"}
	case "PPJ":
		return Location{"Tjipara-Java Island", "Asia/Jakarta"}
	case "PPK":
		return Location{"Petropavlosk", "Asia/Almaty"}
	case "PPL":
		return Location{"Phaplu", "Asia/Kathmandu"}
	case "PPM":
		return Location{"Pompano Beach", "America/New_York"}
	case "PPN":
		return Location{"Popayan", "America/Bogota"}
	case "PPP":
		return Location{"Proserpine", "Australia/Brisbane"}
	case "PPQ":
		return Location{"", "Pacific/Auckland"}
	case "PPR":
		return Location{"Pasir Pengarayan-Sumatra Island", "Asia/Jakarta"}
	case "PPS":
		return Location{"Puerto Princesa City", "Asia/Manila"}
	case "PPT":
		return Location{"Papeete", "Pacific/Tahiti"}
	case "PPU":
		return Location{"Pa Pun", "Asia/Yangon"}
	case "PPW":
		return Location{"Papa Westray", "Europe/London"}
	case "PPY":
		return Location{"Pouso Alegre", "America/Sao_Paulo"}
	case "PQC":
		return Location{"Duong Dong", "Asia/Ho_Chi_Minh"}
	case "PQE":
		return Location{"La Dorada", "America/Bogota"}
	case "PQI":
		return Location{"Presque Isle", "America/New_York"}
	case "PQM":
		return Location{"", "America/Mexico_City"}
	case "PQQ":
		return Location{"Port Macquarie", "Australia/Sydney"}
	case "PRA":
		return Location{"Parana", "America/Argentina/Cordoba"}
	case "PRB":
		return Location{"Paso Robles", "America/Los_Angeles"}
	case "PRC":
		return Location{"Prescott", "America/Phoenix"}
	case "PRD":
		return Location{"Pardoo", "Australia/Perth"}
	case "PRG":
		return Location{"Prague", "Europe/Prague"}
	case "PRH":
		return Location{"", "Asia/Bangkok"}
	case "PRI":
		return Location{"Praslin Island", "Indian/Mahe"}
	case "PRK":
		return Location{"Prieska", "Africa/Johannesburg"}
	case "PRM":
		return Location{"", "Europe/Lisbon"}
	case "PRN":
		return Location{"Prishtina", "Europe/Belgrade"}
	case "PRO":
		return Location{"Perry", "America/Chicago"}
	case "PRP":
		return Location{"Propriano", "Europe/Paris"}
	case "PRQ":
		return Location{"Presidencia Roque Saenz Pena", "America/Argentina/Cordoba"}
	case "PRR":
		return Location{"Paruma", "America/Guyana"}
	case "PRS":
		return Location{"Parasi", "Pacific/Guadalcanal"}
	case "PRU":
		return Location{"Pye", "Asia/Yangon"}
	case "PRV":
		return Location{"Prerov", "Europe/Prague"}
	case "PRX":
		return Location{"Paris", "America/Chicago"}
	case "PRY":
		return Location{"Pretoria", "Africa/Johannesburg"}
	case "PSA":
		return Location{"Pisa", "Europe/Rome"}
	case "PSB":
		return Location{"Philipsburg", "America/New_York"}
	case "PSC":
		return Location{"Pasco", "America/Los_Angeles"}
	case "PSD":
		return Location{"Port Said", "Africa/Cairo"}
	case "PSE":
		return Location{"Ponce", "America/Puerto_Rico"}
	case "PSF":
		return Location{"Pittsfield", "America/New_York"}
	case "PSG":
		return Location{"Petersburg", "America/Sitka"}
	case "PSH":
		return Location{"Sankt Peter-Ording", "Europe/Berlin"}
	case "PSI":
		return Location{"Pasni", "Asia/Karachi"}
	case "PSJ":
		return Location{"Poso-Celebes Island", "Asia/Makassar"}
	case "PSK":
		return Location{"Dublin", "America/New_York"}
	case "PSL":
		return Location{"Perth", "Europe/London"}
	case "PSM":
		return Location{"Portsmouth", "America/New_York"}
	case "PSN":
		return Location{"Palestine", "America/Chicago"}
	case "PSO":
		return Location{"Pasto", "America/Bogota"}
	case "PSP":
		return Location{"Palm Springs", "America/Los_Angeles"}
	case "PSR":
		return Location{"Pescara", "Europe/Rome"}
	case "PSS":
		return Location{"Posadas", "America/Argentina/Cordoba"}
	case "PSU":
		return Location{"Putussibau-Borneo Island", "Asia/Pontianak"}
	case "PSW":
		return Location{"Passos", "America/Sao_Paulo"}
	case "PSX":
		return Location{"Palacios", "America/Chicago"}
	case "PSY":
		return Location{"Stanley", "Atlantic/Stanley"}
	case "PSZ":
		return Location{"Puerto Suarez", "America/La_Paz"}
	case "PTA":
		return Location{"Port Alsworth", "America/Anchorage"}
	case "PTB":
		return Location{"Petersburg", "America/New_York"}
	case "PTF":
		return Location{"Malolo Lailai Island", "Pacific/Fiji"}
	case "PTG":
		return Location{"Potgietersrus", "Africa/Johannesburg"}
	case "PTH":
		return Location{"Port Heiden", "America/Anchorage"}
	case "PTJ":
		return Location{"", "Australia/Melbourne"}
	case "PTK":
		return Location{"Pontiac", "America/Detroit"}
	case "PTM":
		return Location{"Palmarito", "America/Caracas"}
	case "PTN":
		return Location{"Patterson", "America/Chicago"}
	case "PTO":
		return Location{"Pato Branco", "America/Sao_Paulo"}
	case "PTP":
		return Location{"Pointe-a-Pitre Le Raizet", "America/Guadeloupe"}
	case "PTQ":
		return Location{"Porto De Moz", "America/Belem"}
	case "PTS":
		return Location{"Pittsburg", "America/Chicago"}
	case "PTT":
		return Location{"Pratt", "America/Chicago"}
	case "PTU":
		return Location{"Platinum", "America/Anchorage"}
	case "PTV":
		return Location{"Porterville", "America/Los_Angeles"}
	case "PTW":
		return Location{"Pottstown", "America/New_York"}
	case "PTX":
		return Location{"Pitalito", "America/Bogota"}
	case "PTY":
		return Location{"Tocumen", "America/Panama"}
	case "PTZ":
		return Location{"Shell Mera", "America/Guayaquil"}
	case "PUB":
		return Location{"Pueblo", "America/Denver"}
	case "PUC":
		return Location{"Price", "America/Denver"}
	case "PUD":
		return Location{"Puerto Deseado", "America/Argentina/Rio_Gallegos"}
	case "PUE":
		return Location{"Puerto Obaldia", "America/Panama"}
	case "PUF":
		return Location{"Pau/Pyrenees (Uzein)", "Europe/Paris"}
	case "PUG":
		return Location{"", "Australia/Adelaide"}
	case "PUJ":
		return Location{"Punta Cana", "America/Santo_Domingo"}
	case "PUK":
		return Location{"Pukarua", "Pacific/Tahiti"}
	case "PUN":
		return Location{"Punia", "Africa/Lubumbashi"}
	case "PUP":
		return Location{"Po", "Africa/Ouagadougou"}
	case "PUQ":
		return Location{"Punta Arenas", "America/Punta_Arenas"}
	case "PUR":
		return Location{"Puerto Rico/Manuripi", "America/La_Paz"}
	case "PUS":
		return Location{"Busan", "Asia/Seoul"}
	case "PUU":
		return Location{"Puerto Asis", "America/Bogota"}
	case "PUV":
		return Location{"Poum", "Pacific/Noumea"}
	case "PUW":
		return Location{"Pullman/Moscow", "America/Los_Angeles"}
	case "PUX":
		return Location{"Puerto Varas", "America/Santiago"}
	case "PUY":
		return Location{"Pula", "Europe/Zagreb"}
	case "PUZ":
		return Location{"Puerto Cabezas", "America/Managua"}
	case "PVA":
		return Location{"Providencia", "America/Bogota"}
	case "PVC":
		return Location{"Provincetown", "America/New_York"}
	case "PVD":
		return Location{"Providence", "America/New_York"}
	case "PVE":
		return Location{"El Porvenir", "America/Panama"}
	case "PVF":
		return Location{"Placerville", "America/Los_Angeles"}
	case "PVG":
		return Location{"Shanghai", "Asia/Shanghai"}
	case "PVH":
		return Location{"Porto Velho", "America/Porto_Velho"}
	case "PVI":
		return Location{"Paranavai", "America/Sao_Paulo"}
	case "PVK":
		return Location{"Preveza/Lefkada", "Europe/Athens"}
	case "PVL":
		return Location{"Pikeville", "America/New_York"}
	case "PVN":
		return Location{"Dolna Mitropoliya", "Europe/Sofia"}
	case "PVO":
		return Location{"Portoviejo", "America/Guayaquil"}
	case "PVR":
		return Location{"Puerto Vallarta", "America/Bahia_Banderas"}
	case "PVS":
		return Location{"Chukotka", "Asia/Anadyr"}
	case "PVU":
		return Location{"Provo", "America/Denver"}
	case "PVW":
		return Location{"Plainview", "America/Chicago"}
	case "PWA":
		return Location{"Oklahoma City", "America/Chicago"}
	case "PWD":
		return Location{"Plentywood", "America/Denver"}
	case "PWE":
		return Location{"Pevek", "Asia/Anadyr"}
	case "PWK":
		return Location{"Chicago/Prospect Heights/Wheeling", "America/Chicago"}
	case "PWM":
		return Location{"Portland", "America/New_York"}
	case "PWN":
		return Location{"Pitts Town", "America/Nassau"}
	case "PWO":
		return Location{"Pweto", "Africa/Lubumbashi"}
	case "PWQ":
		return Location{"Pavlodar", "Asia/Almaty"}
	case "PWT":
		return Location{"Bremerton", "America/Los_Angeles"}
	case "PWY":
		return Location{"Pinedale", "America/Denver"}
	case "PXM":
		return Location{"Puerto Escondido", "America/Mexico_City"}
	case "PXO":
		return Location{"Vila Baleira", "Atlantic/Madeira"}
	case "PXR":
		return Location{"", "Asia/Bangkok"}
	case "PXU":
		return Location{"Pleiku", "Asia/Ho_Chi_Minh"}
	case "PYA":
		return Location{"Velasquez", "America/Bogota"}
	case "PYB":
		return Location{"Jeypore", "Asia/Kolkata"}
	case "PYE":
		return Location{"Penrhyn Island", "Pacific/Rarotonga"}
	case "PYG":
		return Location{"Pakyong", "Asia/Kolkata"}
	case "PYH":
		return Location{"", "America/Bogota"}
	case "PYJ":
		return Location{"Yakutia", "Asia/Yakutsk"}
	case "PYK":
		return Location{"", "Asia/Tehran"}
	case "PYM":
		return Location{"Plymouth", "America/New_York"}
	case "PYO":
		return Location{"Puerto Putumayo", "America/Guayaquil"}
	case "PYR":
		return Location{"Andravida", "Europe/Athens"}
	case "PYY":
		return Location{"Mae Hong Son", "Asia/Bangkok"}
	case "PZA":
		return Location{"Paz De Ariporo", "America/Bogota"}
	case "PZB":
		return Location{"Pietermaritzburg", "Africa/Johannesburg"}
	case "PZH":
		return Location{"Fort Sandeman", "Asia/Karachi"}
	case "PZI":
		return Location{"Panzhihua", "Asia/Shanghai"}
	case "PZL":
		return Location{"Phinda", "Africa/Johannesburg"}
	case "PZO":
		return Location{"Puerto Ordaz-Ciudad Guayana", "America/Caracas"}
	case "PZU":
		return Location{"Port Sudan", "Africa/Khartoum"}
	case "PZY":
		return Location{"Piestany", "Europe/Bratislava"}
	case "QAC":
		return Location{"Castro", "America/Sao_Paulo"}
	case "QAK":
		return Location{"Barbacena", "America/Sao_Paulo"}
	case "QAM":
		return Location{"Amiens/Glisy", "Europe/Paris"}
	case "QAQ":
		return Location{"L'Aquila", "Europe/Rome"}
	case "QAR":
		return Location{"Arnhem", "Europe/Amsterdam"}
	case "QBC":
		return Location{"Bella Coola", "America/Vancouver"}
	case "QBQ":
		return Location{"St Nazaire", "Europe/Paris"}
	case "QBX":
		return Location{"Sobral", "America/Fortaleza"}
	case "QCB":
		return Location{"Bamberg", "Europe/Berlin"}
	case "QCH":
		return Location{"Colatina", "America/Sao_Paulo"}
	case "QCJ":
		return Location{"Botucatu", "America/Sao_Paulo"}
	case "QCN":
		return Location{"Hohn", "Europe/Berlin"}
	case "QCO":
		return Location{"Colon", "America/Havana"}
	case "QCP":
		return Location{"Currais Novos", "America/Fortaleza"}
	case "QCR":
		return Location{"Curitibanos", "America/Sao_Paulo"}
	case "QCY":
		return Location{"Coningsby", "Europe/London"}
	case "QDB":
		return Location{"Cachoeira Do Sul", "America/Sao_Paulo"}
	case "QDC":
		return Location{"Dracena", "America/Sao_Paulo"}
	case "QDF":
		return Location{"Conselheiro Lafaiete", "America/Sao_Paulo"}
	case "QDJ":
		return Location{"", "Africa/Algiers"}
	case "QEF":
		return Location{"Egelsbach", "Europe/Berlin"}
	case "QEO":
		return Location{"Czechowice-Dziedzice", "Europe/Warsaw"}
	case "QFD":
		return Location{"", "Africa/Algiers"}
	case "QFO":
		return Location{"Duxford", "Europe/London"}
	case "QFR":
		return Location{"Frosinone", "Europe/Rome"}
	case "QGA":
		return Location{"Guaira", "America/Campo_Grande"}
	case "QGB":
		return Location{"Limeira", "America/Sao_Paulo"}
	case "QGC":
		return Location{"Lencois Paulista", "America/Sao_Paulo"}
	case "QGF":
		return Location{"Montenegro", "America/Sao_Paulo"}
	case "QGP":
		return Location{"Garanhuns", "America/Recife"}
	case "QGS":
		return Location{"Alagoinhas", "America/Bahia"}
	case "QGU":
		return Location{"Gifu", "Asia/Tokyo"}
	case "QGY":
		return Location{"Gyor", "Europe/Budapest"}
	case "QHA":
		return Location{"Hasselt", "Europe/Brussels"}
	case "QHB":
		return Location{"Piracicaba", "America/Sao_Paulo"}
	case "QHN":
		return Location{"Taguatinga", "America/Araguaina"}
	case "QHP":
		return Location{"Taubate", "America/Sao_Paulo"}
	case "QHR":
		return Location{"Debre Zeyit", "Africa/Addis_Ababa"}
	case "QHU":
		return Location{"Husum", "Europe/Berlin"}
	case "QHV":
		return Location{"Novo Hamburgo", "America/Sao_Paulo"}
	case "QID":
		return Location{"Tres Coracoes", "America/Sao_Paulo"}
	case "QIE":
		return Location{"Istres/Le Tube", "Europe/Paris"}
	case "QIG":
		return Location{"Iguatu", "America/Fortaleza"}
	case "QIQ":
		return Location{"Rio Claro", "America/Sao_Paulo"}
	case "QIT":
		return Location{"Itapetinga", "America/Bahia"}
	case "QJB":
		return Location{"Jubail", "Asia/Riyadh"}
	case "QJD":
		return Location{"", "Australia/Sydney"}
	case "QKX":
		return Location{"", "Europe/Oslo"}
	case "QLA":
		return Location{"Lasham", "Europe/London"}
	case "QLC":
		return Location{"", "Europe/Warsaw"}
	case "QLD":
		return Location{"", "Africa/Algiers"}
	case "QLF":
		return Location{"", "Europe/Helsinki"}
	case "QLS":
		return Location{"Lausanne", "Europe/Zurich"}
	case "QLT":
		return Location{"Latina", "Europe/Rome"}
	case "QLU":
		return Location{"", "Europe/Warsaw"}
	case "QMF":
		return Location{"Mafra", "America/Sao_Paulo"}
	case "QMJ":
		return Location{"", "Asia/Tehran"}
	case "QMM":
		return Location{"Marina Di Massa", "Europe/Rome"}
	case "QMZ":
		return Location{"Mainz", "Europe/Berlin"}
	case "QNC":
		return Location{"", "Europe/Zurich"}
	case "QND":
		return Location{"Novi Sad", "Europe/Belgrade"}
	case "QNJ":
		return Location{"Annemasse", "Europe/Paris"}
	case "QNM":
		return Location{"Namur", "Europe/Brussels"}
	case "QNS":
		return Location{"Porto Alegre", "America/Sao_Paulo"}
	case "QNV":
		return Location{"Nova Iguacu", "America/Sao_Paulo"}
	case "QNX":
		return Location{"Macon/Charnay", "Europe/Paris"}
	case "QOA":
		return Location{"Mococa", "America/Sao_Paulo"}
	case "QOB":
		return Location{"Ansbach", "Europe/Berlin"}
	case "QOE":
		return Location{"", "Europe/Berlin"}
	case "QOJ":
		return Location{"Sao Borja", "America/Sao_Paulo"}
	case "QOW":
		return Location{"Owerri", "Africa/Lagos"}
	case "QPA":
		return Location{"Padova", "Europe/Rome"}
	case "QPD":
		return Location{"Pinar del Rio", "America/Havana"}
	case "QPG":
		return Location{"", "Asia/Singapore"}
	case "QPH":
		return Location{"Palapye", "Africa/Gaborone"}
	case "QPM":
		return Location{"Opole", "Europe/Warsaw"}
	case "QPS":
		return Location{"Pirassununga", "America/Sao_Paulo"}
	case "QPZ":
		return Location{"Piacenza", "Europe/Rome"}
	case "QRA":
		return Location{"Johannesburg", "Africa/Johannesburg"}
	case "QRC":
		return Location{"Rancagua", "America/Santiago"}
	case "QRM":
		return Location{"", "Australia/Sydney"}
	case "QRO":
		return Location{"Queretaro", "America/Mexico_City"}
	case "QRR":
		return Location{"", "Australia/Sydney"}
	case "QRT":
		return Location{"Rieti", "Europe/Rome"}
	case "QRZ":
		return Location{"Resende", "America/Sao_Paulo"}
	case "QSA":
		return Location{"Sabadell", "Europe/Madrid"}
	case "QSC":
		return Location{"Sao Carlos", "America/Sao_Paulo"}
	case "QSF":
		return Location{"Setif", "Africa/Algiers"}
	case "QSI":
		return Location{"Moshi", "Africa/Dar_es_Salaam"}
	case "QSN":
		return Location{"San Nicolas", "America/Havana"}
	case "QSR":
		return Location{"Salerno", "Europe/Rome"}
	case "QSX":
		return Location{"New Amsterdam", "America/Guyana"}
	case "QTJ":
		return Location{"Blois", "Europe/Paris"}
	case "QUG":
		return Location{"Chichester", "Europe/London"}
	case "QUN":
		return Location{"Chun Chon City", "Asia/Seoul"}
	case "QUS":
		return Location{"Gusau", "Africa/Lagos"}
	case "QUT":
		return Location{"", "Asia/Tokyo"}
	case "QUY":
		return Location{"St. Ives", "Europe/London"}
	case "QVA":
		return Location{"Varese", "Europe/Rome"}
	case "QVB":
		return Location{"Uniao Da Vitoria", "America/Sao_Paulo"}
	case "QVE":
		return Location{"Forssa", "Europe/Helsinki"}
	case "QVP":
		return Location{"Avare", "America/Sao_Paulo"}
	case "QWS":
		return Location{"Nowy Targ", "Europe/Warsaw"}
	case "QWV":
		return Location{"Valjevo", "Europe/Belgrade"}
	case "QXB":
		return Location{"Lyon", "Europe/Paris"}
	case "QXC":
		return Location{"Barra De Santo Antonio", "America/Maceio"}
	case "QYD":
		return Location{"Gdynia", "Europe/Warsaw"}
	case "QYR":
		return Location{"Troyes/Barberey", "Europe/Paris"}
	case "QYY":
		return Location{"Bialystok", "Europe/Warsaw"}
	case "QZD":
		return Location{"Szeged", "Europe/Budapest"}
	case "QZN":
		return Location{"", "Africa/Algiers"}
	case "QZO":
		return Location{"Arezzo", "Europe/Rome"}
	case "RAB":
		return Location{"Tokua", "Pacific/Port_Moresby"}
	case "RAC":
		return Location{"Racine", "America/Chicago"}
	case "RAE":
		return Location{"Arar", "Asia/Riyadh"}
	case "RAF":
		return Location{"Rafaela", "America/Argentina/Cordoba"}
	case "RAG":
		return Location{"", "Pacific/Auckland"}
	case "RAH":
		return Location{"Rafha", "Asia/Riyadh"}
	case "RAI":
		return Location{"Praia", "Atlantic/Cape_Verde"}
	case "RAJ":
		return Location{"Rajkot", "Asia/Kolkata"}
	case "RAK":
		return Location{"Marrakech", "Africa/Casablanca"}
	case "RAL":
		return Location{"Riverside", "America/Los_Angeles"}
	case "RAM":
		return Location{"", "Australia/Darwin"}
	case "RAN":
		return Location{"Ravenna", "Europe/Rome"}
	case "RAO":
		return Location{"Ribeirao Preto", "America/Sao_Paulo"}
	case "RAP":
		return Location{"Rapid City", "America/Denver"}
	case "RAR":
		return Location{"Avarua", "Pacific/Rarotonga"}
	case "RAS":
		return Location{"Rasht", "Asia/Tehran"}
	case "RAV":
		return Location{"Cravo Norte", "America/Bogota"}
	case "RAZ":
		return Location{"Rawalakot", "Asia/Karachi"}
	case "RBA":
		return Location{"Rabat", "Africa/Casablanca"}
	case "RBB":
		return Location{"Borba", "America/Manaus"}
	case "RBC":
		return Location{"", "Australia/Melbourne"}
	case "RBD":
		return Location{"Dallas", "America/Chicago"}
	case "RBE":
		return Location{"Ratanakiri", "Asia/Phnom_Penh"}
	case "RBG":
		return Location{"Roseburg", "America/Los_Angeles"}
	case "RBJ":
		return Location{"", "Asia/Tokyo"}
	case "RBL":
		return Location{"Red Bluff", "America/Los_Angeles"}
	case "RBM":
		return Location{"Straubing", "Europe/Berlin"}
	case "RBO":
		return Location{"Robore", "America/La_Paz"}
	case "RBQ":
		return Location{"Rurenabaque", "America/La_Paz"}
	case "RBR":
		return Location{"Rio Branco", "America/Rio_Branco"}
	case "RBS":
		return Location{"", "Australia/Melbourne"}
	case "RBT":
		return Location{"Marsabit", "Africa/Nairobi"}
	case "RBU":
		return Location{"Roebourne", "Australia/Perth"}
	case "RBV":
		return Location{"Ramata", "Pacific/Guadalcanal"}
	case "RBW":
		return Location{"Walterboro", "America/New_York"}
	case "RBX":
		return Location{"Rumbek", "Africa/Juba"}
	case "RBY":
		return Location{"Ruby", "America/Anchorage"}
	case "RCA":
		return Location{"Rapid City", "America/Denver"}
	case "RCB":
		return Location{"Richards Bay", "Africa/Johannesburg"}
	case "RCH":
		return Location{"Riohacha", "America/Bogota"}
	case "RCK":
		return Location{"Rockdale", "America/Chicago"}
	case "RCL":
		return Location{"Redcliffe", "Pacific/Efate"}
	case "RCM":
		return Location{"", "Australia/Brisbane"}
	case "RCO":
		return Location{"Rochefort/Saint-Agnant", "Europe/Paris"}
	case "RCQ":
		return Location{"Reconquista", "America/Argentina/Cordoba"}
	case "RCR":
		return Location{"Rochester", "America/Indiana/Indianapolis"}
	case "RCS":
		return Location{"Rochester", "Europe/London"}
	case "RCT":
		return Location{"Reed City", "America/Detroit"}
	case "RCU":
		return Location{"Rio Cuarto", "America/Argentina/Cordoba"}
	case "RCY":
		return Location{"", "America/Nassau"}
	case "RDB":
		return Location{"Red Dog", "America/Nome"}
	case "RDC":
		return Location{"Redencao", "America/Belem"}
	case "RDD":
		return Location{"Redding", "America/Los_Angeles"}
	case "RDE":
		return Location{"Merdei-Papua Island", "Asia/Jayapura"}
	case "RDG":
		return Location{"Reading", "America/New_York"}
	case "RDM":
		return Location{"Redmond", "America/Los_Angeles"}
	case "RDN":
		return Location{"Redang", "Asia/Kuala_Lumpur"}
	case "RDO":
		return Location{"Radom", "Europe/Warsaw"}
	case "RDR":
		return Location{"Grand Forks", "America/Chicago"}
	case "RDS":
		return Location{"Rincon de los Sauces", "America/Argentina/Mendoza"}
	case "RDT":
		return Location{"Richard Toll", "Africa/Dakar"}
	case "RDU":
		return Location{"Raleigh/Durham", "America/New_York"}
	case "RDZ":
		return Location{"Rodez/Marcillac", "Europe/Paris"}
	case "REA":
		return Location{"", "Pacific/Tahiti"}
	case "REB":
		return Location{"Larz", "Europe/Berlin"}
	case "REC":
		return Location{"Recife", "America/Recife"}
	case "RED":
		return Location{"Reedsville", "America/New_York"}
	case "REG":
		return Location{"Reggio Calabria", "Europe/Rome"}
	case "REI":
		return Location{"Regina", "America/Cayenne"}
	case "REL":
		return Location{"Rawson", "America/Argentina/Catamarca"}
	case "REN":
		return Location{"Orenburg", "Asia/Yekaterinburg"}
	case "REO":
		return Location{"Rome", "America/Boise"}
	case "REP":
		return Location{"Siem Reap", "Asia/Phnom_Penh"}
	case "REQ":
		return Location{"Requena", "America/Lima"}
	case "RER":
		return Location{"Retalhuleu", "America/Guatemala"}
	case "RES":
		return Location{"Resistencia", "America/Argentina/Cordoba"}
	case "RET":
		return Location{"", "Europe/Oslo"}
	case "REU":
		return Location{"Reus", "Europe/Madrid"}
	case "REX":
		return Location{"Reynosa", "America/Matamoros"}
	case "REY":
		return Location{"Reyes", "America/La_Paz"}
	case "RFA":
		return Location{"Rafai", "Africa/Bangui"}
	case "RFD":
		return Location{"Chicago/Rockford", "America/Chicago"}
	case "RFG":
		return Location{"Refugio", "America/Chicago"}
	case "RFN":
		return Location{"Raufarhofn", "Atlantic/Reykjavik"}
	case "RFP":
		return Location{"Uturoa", "Pacific/Tahiti"}
	case "RFR":
		return Location{"Rio Frio / Progreso", "America/Costa_Rica"}
	case "RFS":
		return Location{"La Rosita", "America/Managua"}
	case "RGA":
		return Location{"Rio Grande", "America/Argentina/Ushuaia"}
	case "RGH":
		return Location{"Balurghat", "Asia/Kolkata"}
	case "RGI":
		return Location{"", "Pacific/Tahiti"}
	case "RGK":
		return Location{"Gorno-Altaysk", "Asia/Barnaul"}
	case "RGL":
		return Location{"Rio Gallegos", "America/Argentina/Rio_Gallegos"}
	case "RGN":
		return Location{"Yangon", "Asia/Yangon"}
	case "RGO":
		return Location{"", "Asia/Pyongyang"}
	case "RGS":
		return Location{"Burgos", "Europe/Madrid"}
	case "RGT":
		return Location{"Rengat-Sumatra Island", "Asia/Jakarta"}
	case "RHA":
		return Location{"Reykholar", "Atlantic/Reykjavik"}
	case "RHD":
		return Location{"Rio Hondo", "America/Argentina/Cordoba"}
	case "RHE":
		return Location{"Reims/Champagne", "Europe/Paris"}
	case "RHG":
		return Location{"Ruhengeri", "Africa/Kigali"}
	case "RHI":
		return Location{"Rhinelander", "America/Chicago"}
	case "RHL":
		return Location{"", "Australia/Perth"}
	case "RHN":
		return Location{"Rosh Pinah", "Africa/Windhoek"}
	case "RHO":
		return Location{"Rodes Island", "Europe/Athens"}
	case "RHP":
		return Location{"Ramechhap", "Asia/Kathmandu"}
	case "RHT":
		return Location{"Badanjilin", "Asia/Shanghai"}
	case "RHV":
		return Location{"San Jose", "America/Los_Angeles"}
	case "RIA":
		return Location{"Santa Maria", "America/Sao_Paulo"}
	case "RIB":
		return Location{"Riberalta", "America/La_Paz"}
	case "RIC":
		return Location{"Richmond", "America/New_York"}
	case "RID":
		return Location{"Richmond", "America/Indiana/Indianapolis"}
	case "RIE":
		return Location{"Rice Lake", "America/Chicago"}
	case "RIF":
		return Location{"Richfield", "America/Denver"}
	case "RIG":
		return Location{"Rio Grande", "America/Sao_Paulo"}
	case "RIH":
		return Location{"Rio Hato", "America/Panama"}
	case "RIJ":
		return Location{"Rioja", "America/Lima"}
	case "RIK":
		return Location{"Nicoya", "America/Costa_Rica"}
	case "RIL":
		return Location{"Rifle", "America/Denver"}
	case "RIM":
		return Location{"Rodriguez de Mendoza", "America/Lima"}
	case "RIN":
		return Location{"Ringi Cove", "Pacific/Guadalcanal"}
	case "RIR":
		return Location{"Riverside/Rubidoux/", "America/Los_Angeles"}
	case "RIS":
		return Location{"Rishiri", "Asia/Tokyo"}
	case "RIV":
		return Location{"Riverside", "America/Los_Angeles"}
	case "RIW":
		return Location{"Riverton", "America/Denver"}
	case "RIX":
		return Location{"Riga", "Europe/Riga"}
	case "RIY":
		return Location{"Riyan", "Asia/Aden"}
	case "RJA":
		return Location{"Rajahmundry", "Asia/Kolkata"}
	case "RJB":
		return Location{"Rajbiraj", "Asia/Kathmandu"}
	case "RJH":
		return Location{"Rajshahi", "Asia/Dhaka"}
	case "RJK":
		return Location{"Rijeka", "Europe/Zagreb"}
	case "RJL":
		return Location{"Logrono", "Europe/Madrid"}
	case "RJN":
		return Location{"", "Asia/Tehran"}
	case "RKA":
		return Location{"", "Pacific/Tahiti"}
	case "RKC":
		return Location{"Montague", "America/Los_Angeles"}
	case "RKD":
		return Location{"Rockland", "America/New_York"}
	case "RKE":
		return Location{"Copenhagen", "Europe/Copenhagen"}
	case "RKH":
		return Location{"Rock Hill", "America/New_York"}
	case "RKO":
		return Location{"Sipora Island", "Asia/Jakarta"}
	case "RKP":
		return Location{"Rockport", "America/Chicago"}
	case "RKR":
		return Location{"Poteau", "America/Chicago"}
	case "RKS":
		return Location{"Rock Springs", "America/Denver"}
	case "RKT":
		return Location{"Ras Al Khaimah", "Asia/Dubai"}
	case "RKV":
		return Location{"Reykjavik", "Atlantic/Reykjavik"}
	case "RKW":
		return Location{"Rockwood", "America/Chicago"}
	case "RLD":
		return Location{"Richland", "America/Los_Angeles"}
	case "RLG":
		return Location{"Rostock", "Europe/Berlin"}
	case "RLK":
		return Location{"Bayannur", "Asia/Shanghai"}
	case "RLO":
		return Location{"Merlo", "America/Argentina/San_Luis"}
	case "RLT":
		return Location{"Arlit", "Africa/Niamey"}
	case "RMA":
		return Location{"Roma", "Australia/Brisbane"}
	case "RMB":
		return Location{"Buraimi", "Asia/Dubai"}
	case "RME":
		return Location{"Rome", "America/New_York"}
	case "RMF":
		return Location{"Marsa Alam", "Africa/Cairo"}
	case "RMG":
		return Location{"Rome", "America/New_York"}
	case "RMI":
		return Location{"Rimini", "Europe/Rome"}
	case "RMK":
		return Location{"", "Australia/Adelaide"}
	case "RML":
		return Location{"Colombo", "Asia/Colombo"}
	case "RMN":
		return Location{"", "Pacific/Port_Moresby"}
	case "RMO":
		return Location{"Chisinau", "Europe/Chisinau"}
	case "RMQ":
		return Location{"Taichung City", "Asia/Taipei"}
	case "RMS":
		return Location{"Ramstein", "Europe/Berlin"}
	case "RMU":
		return Location{"Corvera", "Europe/Madrid"}
	case "RMY":
		return Location{"Mariposa", "America/Los_Angeles"}
	case "RNA":
		return Location{"Arona", "Pacific/Guadalcanal"}
	case "RNB":
		return Location{"", "Europe/Stockholm"}
	case "RNC":
		return Location{"Mc Minnville", "America/Chicago"}
	case "RND":
		return Location{"Universal City", "America/Chicago"}
	case "RNE":
		return Location{"Roanne/Renaison", "Europe/Paris"}
	case "RNH":
		return Location{"New Richmond", "America/Chicago"}
	case "RNI":
		return Location{"Corn Island", "America/Managua"}
	case "RNJ":
		return Location{"", "Asia/Tokyo"}
	case "RNL":
		return Location{"Rennell Island", "Pacific/Guadalcanal"}
	case "RNM":
		return Location{"Ghaba", "Asia/Muscat"}
	case "RNN":
		return Location{"Ronne", "Europe/Copenhagen"}
	case "RNO":
		return Location{"Reno", "America/Los_Angeles"}
	case "RNS":
		return Location{"Rennes/Saint-Jacques", "Europe/Paris"}
	case "RNT":
		return Location{"Renton", "America/Los_Angeles"}
	case "RNU":
		return Location{"Ranau", "Asia/Kuching"}
	case "RNZ":
		return Location{"Rensselaer", "America/Chicago"}
	case "ROA":
		return Location{"Roanoke", "America/New_York"}
	case "ROB":
		return Location{"Monrovia", "Africa/Monrovia"}
	case "ROC":
		return Location{"Rochester", "America/New_York"}
	case "ROD":
		return Location{"Robertson", "Africa/Johannesburg"}
	case "ROG":
		return Location{"Rogers", "America/Chicago"}
	case "ROH":
		return Location{"", "Australia/Brisbane"}
	case "ROI":
		return Location{"", "Asia/Bangkok"}
	case "ROK":
		return Location{"Rockhampton", "Australia/Brisbane"}
	case "RON":
		return Location{"Paipa", "America/Bogota"}
	case "ROO":
		return Location{"Rondonopolis", "America/Cuiaba"}
	case "ROP":
		return Location{"Rota Island", "Pacific/Saipan"}
	case "ROR":
		return Location{"Babelthuap Island", "Pacific/Palau"}
	case "ROS":
		return Location{"Rosario", "America/Argentina/Cordoba"}
	case "ROT":
		return Location{"Rotorua", "Pacific/Auckland"}
	case "ROV":
		return Location{"Rostov-on-Don", "Europe/Moscow"}
	case "ROW":
		return Location{"Roswell", "America/Denver"}
	case "ROX":
		return Location{"Roseau", "America/Chicago"}
	case "ROY":
		return Location{"Rio Mayo", "America/Argentina/Catamarca"}
	case "ROZ":
		return Location{"Rota", "Europe/Madrid"}
	case "RPA":
		return Location{"Rolpa", "Asia/Kathmandu"}
	case "RPB":
		return Location{"", "Australia/Darwin"}
	case "RPM":
		return Location{"", "Australia/Darwin"}
	case "RPN":
		return Location{"Rosh Pina", "Asia/Jerusalem"}
	case "RPR":
		return Location{"Raipur", "Asia/Kolkata"}
	case "RPX":
		return Location{"Roundup", "America/Denver"}
	case "RQA":
		return Location{"Ruoqiang", "Asia/Shanghai"}
	case "RQO":
		return Location{"El Reno", "America/Chicago"}
	case "RQW":
		return Location{"Qayyarah", "Asia/Baghdad"}
	case "RQY":
		return Location{"Shimoga", "Asia/Kolkata"}
	case "RRE":
		return Location{"", "Australia/Adelaide"}
	case "RRG":
		return Location{"Port Mathurin", "Indian/Mauritius"}
	case "RRK":
		return Location{"", "Asia/Kolkata"}
	case "RRL":
		return Location{"Merrill", "America/Chicago"}
	case "RRR":
		return Location{"", "Pacific/Tahiti"}
	case "RRS":
		return Location{"Roros", "Europe/Oslo"}
	case "RRT":
		return Location{"Warroad", "America/Chicago"}
	case "RSA":
		return Location{"Santa Rosa", "America/Argentina/Salta"}
	case "RSB":
		return Location{"", "Australia/Brisbane"}
	case "RSD":
		return Location{"Rock Sound", "America/Nassau"}
	case "RSH":
		return Location{"Russian Mission", "America/Anchorage"}
	case "RSI":
		return Location{"Hanak", "Asia/Riyadh"}
	case "RSK":
		return Location{"Ransiki-Papua Island", "Asia/Jayapura"}
	case "RSL":
		return Location{"Russell", "America/Chicago"}
	case "RSN":
		return Location{"Ruston", "America/Chicago"}
	case "RSS":
		return Location{"Ad Damazin", "Africa/Khartoum"}
	case "RST":
		return Location{"Rochester", "America/Chicago"}
	case "RSU":
		return Location{"Yeosu", "Asia/Seoul"}
	case "RSW":
		return Location{"Fort Myers", "America/New_York"}
	case "RTA":
		return Location{"Rotuma", "Pacific/Fiji"}
	case "RTB":
		return Location{"Roatan Island", "America/Tegucigalpa"}
	case "RTC":
		return Location{"", "Asia/Kolkata"}
	case "RTG":
		return Location{"Satar Tacik-Flores Island", "Asia/Makassar"}
	case "RTM":
		return Location{"Rotterdam", "Europe/Amsterdam"}
	case "RTN":
		return Location{"Raton", "America/Denver"}
	case "RTP":
		return Location{"", "Australia/Brisbane"}
	case "RTS":
		return Location{"", "Australia/Perth"}
	case "RTU":
		return Location{"Maratua Island", "Asia/Makassar"}
	case "RTW":
		return Location{"Saratov", "Europe/Saratov"}
	case "RTY":
		return Location{"", "Australia/Adelaide"}
	case "RUA":
		return Location{"Arua", "Africa/Kampala"}
	case "RUD":
		return Location{"", "Asia/Tehran"}
	case "RUE":
		return Location{"Butembo", "Africa/Lubumbashi"}
	case "RUG":
		return Location{"Rugao", "Asia/Shanghai"}
	case "RUH":
		return Location{"Riyadh", "Asia/Riyadh"}
	case "RUI":
		return Location{"Ruidoso", "America/Denver"}
	case "RUK":
		return Location{"Rukumkot", "Asia/Kathmandu"}
	case "RUL":
		return Location{"Gadhdhoo", "Indian/Maldives"}
	case "RUM":
		return Location{"Rumjatar", "Asia/Kathmandu"}
	case "RUN":
		return Location{"St Denis", "Indian/Reunion"}
	case "RUP":
		return Location{"", "Asia/Kolkata"}
	case "RUR":
		return Location{"", "Pacific/Tahiti"}
	case "RUS":
		return Location{"Marau", "Pacific/Guadalcanal"}
	case "RUT":
		return Location{"Rutland", "America/New_York"}
	case "RUV":
		return Location{"Rubelsanto", "America/Guatemala"}
	case "RUY":
		return Location{"Ruinas de Copan", "America/Tegucigalpa"}
	case "RVA":
		return Location{"", "Indian/Antananarivo"}
	case "RVD":
		return Location{"Rio Verde", "America/Sao_Paulo"}
	case "RVE":
		return Location{"Saravena", "America/Bogota"}
	case "RVI":
		return Location{"Rostov-on-Don", "Europe/Moscow"}
	case "RVK":
		return Location{"Rorvik", "Europe/Oslo"}
	case "RVN":
		return Location{"Rovaniemi", "Europe/Helsinki"}
	case "RVO":
		return Location{"Reivilo", "Africa/Johannesburg"}
	case "RVR":
		return Location{"Green River", "America/Denver"}
	case "RVS":
		return Location{"Tulsa", "America/Chicago"}
	case "RVT":
		return Location{"", "Australia/Perth"}
	case "RVV":
		return Location{"", "Pacific/Tahiti"}
	case "RVY":
		return Location{"Rivera", "America/Montevideo"}
	case "RWF":
		return Location{"Redwood Falls", "America/Chicago"}
	case "RWI":
		return Location{"Rocky Mount", "America/New_York"}
	case "RWL":
		return Location{"Rawlins", "America/Denver"}
	case "RWN":
		return Location{"Rivne", "Europe/Kyiv"}
	case "RXE":
		return Location{"Rexburg", "America/Boise"}
	case "RXS":
		return Location{"Roxas City", "Asia/Manila"}
	case "RYB":
		return Location{"Rybinsk", "Europe/Moscow"}
	case "RYG":
		return Location{"Rygge", "Europe/Oslo"}
	case "RYK":
		return Location{"Rahim Yar Khan", "Asia/Karachi"}
	case "RYN":
		return Location{"Royan/Medis", "Europe/Paris"}
	case "RYO":
		return Location{"El Turbio", "America/Argentina/Rio_Gallegos"}
	case "RZA":
		return Location{"Santa Cruz", "America/Argentina/Rio_Gallegos"}
	case "RZE":
		return Location{"Rzeszow", "Europe/Warsaw"}
	case "RZN":
		return Location{"Ryazan", "Europe/Moscow"}
	case "RZP":
		return Location{"Taytay Airport", "Asia/Manila"}
	case "RZR":
		return Location{"", "Asia/Tehran"}
	case "RZV":
		return Location{"", "Europe/Istanbul"}
	case "RZZ":
		return Location{"Roanoke Rapids", "America/New_York"}
	case "SAA":
		return Location{"Saratoga", "America/Denver"}
	case "SAB":
		return Location{"Saba", "America/Kralendijk"}
	case "SAC":
		return Location{"Sacramento", "America/Los_Angeles"}
	case "SAD":
		return Location{"Safford", "America/Phoenix"}
	case "SAF":
		return Location{"Santa Fe", "America/Denver"}
	case "SAG":
		return Location{"Kakadi", "Asia/Kolkata"}
	case "SAH":
		return Location{"Sana'a", "Asia/Aden"}
	case "SAI":
		return Location{"Siem Reap", "Asia/Phnom_Penh"}
	case "SAK":
		return Location{"Saudarkrokur", "Atlantic/Reykjavik"}
	case "SAL":
		return Location{"Santa Clara", "America/El_Salvador"}
	case "SAN":
		return Location{"San Diego", "America/Los_Angeles"}
	case "SAO":
		return Location{"Sao Paulo", "America/Sao_Paulo"}
	case "SAP":
		return Location{"La Mesa", "America/Tegucigalpa"}
	case "SAQ":
		return Location{"Andros Island", "America/Nassau"}
	case "SAR":
		return Location{"Sparta", "America/Chicago"}
	case "SAS":
		return Location{"Salton City", "America/Los_Angeles"}
	case "SAT":
		return Location{"San Antonio", "America/Chicago"}
	case "SAU":
		return Location{"Sawu-Sawu Island", "Asia/Makassar"}
	case "SAV":
		return Location{"Savannah", "America/New_York"}
	case "SAW":
		return Location{"Istanbul", "Europe/Istanbul"}
	case "SAY":
		return Location{"Siena", "Europe/Rome"}
	case "SAZ":
		return Location{"Sasstown", "Africa/Monrovia"}
	case "SBA":
		return Location{"Santa Barbara", "America/Los_Angeles"}
	case "SBB":
		return Location{"Santa Barbara", "America/Caracas"}
	case "SBD":
		return Location{"San Bernardino", "America/Los_Angeles"}
	case "SBE":
		return Location{"", "Pacific/Port_Moresby"}
	case "SBG":
		return Location{"Sabang-We Island", "Asia/Jakarta"}
	case "SBH":
		return Location{"Gustavia", "America/St_Barthelemy"}
	case "SBI":
		return Location{"Koundara", "Africa/Conakry"}
	case "SBJ":
		return Location{"Sao Mateus", "America/Sao_Paulo"}
	case "SBK":
		return Location{"Saint-Brieuc/Armor", "Europe/Paris"}
	case "SBL":
		return Location{"Santa Ana del Yacuma", "America/La_Paz"}
	case "SBM":
		return Location{"Sheboygan", "America/Chicago"}
	case "SBN":
		return Location{"South Bend", "America/Indiana/Indianapolis"}
	case "SBP":
		return Location{"San Luis Obispo", "America/Los_Angeles"}
	case "SBQ":
		return Location{"Sibi", "Asia/Karachi"}
	case "SBR":
		return Location{"Saibai Island", "Australia/Brisbane"}
	case "SBS":
		return Location{"Steamboat Springs", "America/Denver"}
	case "SBT":
		return Location{"Sabetta", "Asia/Yekaterinburg"}
	case "SBU":
		return Location{"Springbok", "Africa/Johannesburg"}
	case "SBW":
		return Location{"Sibu", "Asia/Kuching"}
	case "SBX":
		return Location{"Shelby", "America/Denver"}
	case "SBY":
		return Location{"Salisbury", "America/New_York"}
	case "SBZ":
		return Location{"Sibiu", "Europe/Bucharest"}
	case "SCB":
		return Location{"Scribner", "America/Chicago"}
	case "SCC":
		return Location{"Deadhorse", "America/Anchorage"}
	case "SCD":
		return Location{"Sulaco", "America/Tegucigalpa"}
	case "SCE":
		return Location{"State College", "America/New_York"}
	case "SCF":
		return Location{"Scottsdale", "America/Phoenix"}
	case "SCG":
		return Location{"", "Australia/Brisbane"}
	case "SCH":
		return Location{"Schenectady", "America/New_York"}
	case "SCI":
		return Location{"", "America/Caracas"}
	case "SCK":
		return Location{"Stockton", "America/Los_Angeles"}
	case "SCL":
		return Location{"Santiago", "America/Santiago"}
	case "SCM":
		return Location{"Scammon Bay", "America/Nome"}
	case "SCN":
		return Location{"Saarbrucken", "Europe/Berlin"}
	case "SCO":
		return Location{"Aktau", "Asia/Aqtau"}
	case "SCP":
		return Location{"Mende", "Europe/Paris"}
	case "SCQ":
		return Location{"Santiago de Compostela", "Europe/Madrid"}
	case "SCR":
		return Location{"Sälen", "Europe/Stockholm"}
	case "SCS":
		return Location{"Shetland Islands", "Europe/London"}
	case "SCT":
		return Location{"Socotra Islands", "Asia/Aden"}
	case "SCU":
		return Location{"Santiago", "America/Havana"}
	case "SCV":
		return Location{"Suceava", "Europe/Bucharest"}
	case "SCW":
		return Location{"Syktyvkar", "Europe/Moscow"}
	case "SCY":
		return Location{"San Cristobal", "Pacific/Galapagos"}
	case "SCZ":
		return Location{"Santa Cruz/Graciosa Bay/Luova", "Pacific/Guadalcanal"}
	case "SDB":
		return Location{"Langebaanweg", "Africa/Johannesburg"}
	case "SDD":
		return Location{"Lubango", "Africa/Luanda"}
	case "SDE":
		return Location{"Santiago del Estero", "America/Argentina/Cordoba"}
	case "SDF":
		return Location{"Louisville", "America/Kentucky/Louisville"}
	case "SDG":
		return Location{"", "Asia/Tehran"}
	case "SDH":
		return Location{"Santa Rosa de Copan", "America/Tegucigalpa"}
	case "SDJ":
		return Location{"Sendai", "Asia/Tokyo"}
	case "SDK":
		return Location{"Sandakan", "Asia/Kuching"}
	case "SDL":
		return Location{"Sundsvall/ Harnosand", "Europe/Stockholm"}
	case "SDM":
		return Location{"San Diego", "America/Los_Angeles"}
	case "SDN":
		return Location{"Sandane", "Europe/Oslo"}
	case "SDP":
		return Location{"Sand Point", "America/Anchorage"}
	case "SDQ":
		return Location{"Santo Domingo", "America/Santo_Domingo"}
	case "SDR":
		return Location{"Santander", "Europe/Madrid"}
	case "SDS":
		return Location{"", "Asia/Tokyo"}
	case "SDT":
		return Location{"Saidu Sharif", "Asia/Karachi"}
	case "SDU":
		return Location{"Rio De Janeiro", "America/Sao_Paulo"}
	case "SDV":
		return Location{"Tel Aviv", "Asia/Jerusalem"}
	case "SDX":
		return Location{"Sedona", "America/Phoenix"}
	case "SDY":
		return Location{"Sidney", "America/Denver"}
	case "SEA":
		return Location{"Seattle", "America/Los_Angeles"}
	case "SEB":
		return Location{"Sabha", "Africa/Tripoli"}
	case "SEE":
		return Location{"San Diego/El Cajon", "America/Los_Angeles"}
	case "SEF":
		return Location{"Sebring", "America/New_York"}
	case "SEG":
		return Location{"Selinsgrove", "America/New_York"}
	case "SEH":
		return Location{"Senggeh-Papua Island", "Asia/Jayapura"}
	case "SEM":
		return Location{"Selma", "America/Chicago"}
	case "SEN":
		return Location{"Southend", "Europe/London"}
	case "SEO":
		return Location{"Seguela", "Africa/Abidjan"}
	case "SEP":
		return Location{"Stephenville", "America/Chicago"}
	case "SEQ":
		return Location{"Bengkalis-Sumatra Island", "Asia/Jakarta"}
	case "SER":
		return Location{"Seymour", "America/Indiana/Indianapolis"}
	case "SEU":
		return Location{"Seronera", "Africa/Dar_es_Salaam"}
	case "SEV":
		return Location{"Sievierodonetsk", "Europe/Kyiv"}
	case "SEW":
		return Location{"Siwa", "Africa/Cairo"}
	case "SEY":
		return Location{"Selibaby", "Africa/Nouakchott"}
	case "SEZ":
		return Location{"Mahe Island", "Indian/Mahe"}
	case "SFA":
		return Location{"Sfax", "Africa/Tunis"}
	case "SFB":
		return Location{"Orlando", "America/New_York"}
	case "SFC":
		return Location{"St-Francois", "America/Guadeloupe"}
	case "SFD":
		return Location{"Inglaterra", "America/Caracas"}
	case "SFE":
		return Location{"", "Asia/Manila"}
	case "SFF":
		return Location{"Spokane", "America/Los_Angeles"}
	case "SFG":
		return Location{"Grand Case", "America/Lower_Princes"}
	case "SFH":
		return Location{"", "America/Tijuana"}
	case "SFI":
		return Location{"Safi", "Africa/Casablanca"}
	case "SFJ":
		return Location{"Kangerlussuaq", "America/Nuuk"}
	case "SFK":
		return Location{"Soure", "America/Belem"}
	case "SFL":
		return Location{"Sao Filipe", "Atlantic/Cape_Verde"}
	case "SFM":
		return Location{"Sanford", "America/New_York"}
	case "SFN":
		return Location{"Santa Fe", "America/Argentina/Cordoba"}
	case "SFO":
		return Location{"San Francisco", "America/Los_Angeles"}
	case "SFQ":
		return Location{"Sanliurfa", "Europe/Istanbul"}
	case "SFS":
		return Location{"Olongapo City", "Asia/Manila"}
	case "SFT":
		return Location{"Skelleftea", "Europe/Stockholm"}
	case "SFZ":
		return Location{"Pawtucket", "America/New_York"}
	case "SGA":
		return Location{"Sheghnan", "Asia/Dushanbe"}
	case "SGC":
		return Location{"Surgut", "Asia/Yekaterinburg"}
	case "SGD":
		return Location{"Sonderborg", "Europe/Copenhagen"}
	case "SGE":
		return Location{"", "Europe/Berlin"}
	case "SGF":
		return Location{"Springfield", "America/Chicago"}
	case "SGG":
		return Location{"Simanggang", "Asia/Kuching"}
	case "SGH":
		return Location{"Springfield", "America/New_York"}
	case "SGI":
		return Location{"Sargodha", "Asia/Karachi"}
	case "SGN":
		return Location{"Ho Chi Minh City", "Asia/Ho_Chi_Minh"}
	case "SGO":
		return Location{"", "Australia/Brisbane"}
	case "SGP":
		return Location{"Shay Gap", "Australia/Perth"}
	case "SGQ":
		return Location{"Sanggata-Timor Island", "Asia/Makassar"}
	case "SGR":
		return Location{"Houston", "America/Chicago"}
	case "SGS":
		return Location{"", "Asia/Manila"}
	case "SGT":
		return Location{"Stuttgart", "America/Chicago"}
	case "SGU":
		return Location{"St George", "America/Denver"}
	case "SGV":
		return Location{"Sierra Grande", "America/Argentina/Salta"}
	case "SGX":
		return Location{"Songea", "Africa/Dar_es_Salaam"}
	case "SGY":
		return Location{"Skagway", "America/Juneau"}
	case "SGZ":
		return Location{"", "Asia/Bangkok"}
	case "SHA":
		return Location{"Shanghai", "Asia/Shanghai"}
	case "SHB":
		return Location{"Nakashibetsu", "Asia/Tokyo"}
	case "SHC":
		return Location{"Shire", "Africa/Addis_Ababa"}
	case "SHD":
		return Location{"Staunton/Waynesboro/Harrisonburg", "America/New_York"}
	case "SHE":
		return Location{"Shenyang", "Asia/Shanghai"}
	case "SHG":
		return Location{"Shungnak", "America/Anchorage"}
	case "SHH":
		return Location{"Shishmaref", "America/Nome"}
	case "SHI":
		return Location{"", "Asia/Tokyo"}
	case "SHJ":
		return Location{"Sharjah", "Asia/Dubai"}
	case "SHK":
		return Location{"Sehonghong", "Africa/Maseru"}
	case "SHL":
		return Location{"Shillong", "Asia/Kolkata"}
	case "SHM":
		return Location{"Shirahama", "Asia/Tokyo"}
	case "SHN":
		return Location{"Shelton", "America/Los_Angeles"}
	case "SHO":
		return Location{"", "Asia/Seoul"}
	case "SHP":
		return Location{"Qinhuangdao", "Asia/Shanghai"}
	case "SHQ":
		return Location{"", "Australia/Brisbane"}
	case "SHR":
		return Location{"Sheridan", "America/Denver"}
	case "SHS":
		return Location{"Shashi", "Asia/Shanghai"}
	case "SHT":
		return Location{"", "Australia/Melbourne"}
	case "SHU":
		return Location{"", "Australia/Darwin"}
	case "SHV":
		return Location{"Shreveport", "America/Chicago"}
	case "SHW":
		return Location{"", "Asia/Riyadh"}
	case "SHX":
		return Location{"Shageluk", "America/Anchorage"}
	case "SHY":
		return Location{"Shinyanga", "Africa/Dar_es_Salaam"}
	case "SHZ":
		return Location{"Seshutes", "Africa/Maseru"}
	case "SIA":
		return Location{"Xi'an", "Asia/Shanghai"}
	case "SIB":
		return Location{"Sibiti", "Africa/Brazzaville"}
	case "SID":
		return Location{"Espargos", "Atlantic/Cape_Verde"}
	case "SIE":
		return Location{"", "Europe/Lisbon"}
	case "SIF":
		return Location{"Simara", "Asia/Kathmandu"}
	case "SIG":
		return Location{"San Juan", "America/Puerto_Rico"}
	case "SIH":
		return Location{"Silgadi Doti", "Asia/Kathmandu"}
	case "SII":
		return Location{"Sidi Ifni", "Africa/Casablanca"}
	case "SIJ":
		return Location{"Siglufjordur", "Atlantic/Reykjavik"}
	case "SIK":
		return Location{"Sikeston", "America/Chicago"}
	case "SIL":
		return Location{"Sila Mission", "Pacific/Port_Moresby"}
	case "SIM":
		return Location{"Simbai", "Pacific/Port_Moresby"}
	case "SIN":
		return Location{"Singapore", "Asia/Singapore"}
	case "SIO":
		return Location{"", "Australia/Hobart"}
	case "SIP":
		return Location{"Simferopol", "Europe/Simferopol"}
	case "SIQ":
		return Location{"Pasirkuning-Singkep Island", "Asia/Jakarta"}
	case "SIR":
		return Location{"Sion", "Europe/Zurich"}
	case "SIS":
		return Location{"Sishen", "Africa/Johannesburg"}
	case "SIT":
		return Location{"Sitka", "America/Sitka"}
	case "SIU":
		return Location{"Siuna", "America/Managua"}
	case "SIV":
		return Location{"Sullivan", "America/Indiana/Indianapolis"}
	case "SIW":
		return Location{"Parapat-Sumatra Island", "Asia/Jakarta"}
	case "SIX":
		return Location{"Singleton", "Australia/Sydney"}
	case "SIY":
		return Location{"Montague", "America/Los_Angeles"}
	case "SJA":
		return Location{"San Juan de Marcona", "America/Lima"}
	case "SJB":
		return Location{"San Joaquin", "America/La_Paz"}
	case "SJC":
		return Location{"San Jose", "America/Los_Angeles"}
	case "SJD":
		return Location{"San Jose del Cabo", "America/Mazatlan"}
	case "SJE":
		return Location{"San Jose Del Guaviare", "America/Bogota"}
	case "SJI":
		return Location{"San Jose", "Asia/Manila"}
	case "SJJ":
		return Location{"Sarajevo", "Europe/Sarajevo"}
	case "SJK":
		return Location{"Sao Jose Dos Campos", "America/Sao_Paulo"}
	case "SJL":
		return Location{"Sao Gabriel Da Cachoeira", "America/Manaus"}
	case "SJN":
		return Location{"St Johns", "America/Phoenix"}
	case "SJO":
		return Location{"San Jose", "America/Costa_Rica"}
	case "SJP":
		return Location{"Sao Jose Do Rio Preto", "America/Sao_Paulo"}
	case "SJQ":
		return Location{"Sesheke", "Africa/Lusaka"}
	case "SJS":
		return Location{"San Jose de Chiquitos", "America/La_Paz"}
	case "SJT":
		return Location{"San Angelo", "America/Chicago"}
	case "SJU":
		return Location{"San Juan", "America/Puerto_Rico"}
	case "SJV":
		return Location{"San Javier", "America/La_Paz"}
	case "SJW":
		return Location{"Shijiazhuang", "Asia/Shanghai"}
	case "SJY":
		return Location{"Seinajoki / Ilmajoki", "Europe/Helsinki"}
	case "SJZ":
		return Location{"Velas", "Atlantic/Azores"}
	case "SKA":
		return Location{"Spokane", "America/Los_Angeles"}
	case "SKB":
		return Location{"Basseterre", "America/St_Kitts"}
	case "SKC":
		return Location{"Suki", "Pacific/Port_Moresby"}
	case "SKD":
		return Location{"Samarkand", "Asia/Samarkand"}
	case "SKE":
		return Location{"Geiteryggen", "Europe/Oslo"}
	case "SKF":
		return Location{"San Antonio", "America/Chicago"}
	case "SKG":
		return Location{"Thessaloniki", "Europe/Athens"}
	case "SKH":
		return Location{"Surkhet", "Asia/Kathmandu"}
	case "SKK":
		return Location{"Shaktoolik", "America/Anchorage"}
	case "SKL":
		return Location{"Broadford", "Europe/London"}
	case "SKN":
		return Location{"Hadsel", "Europe/Oslo"}
	case "SKO":
		return Location{"Sokoto", "Africa/Lagos"}
	case "SKP":
		return Location{"Skopje", "Europe/Skopje"}
	case "SKQ":
		return Location{"Sekakes", "Africa/Maseru"}
	case "SKS":
		return Location{"Vojens", "Europe/Copenhagen"}
	case "SKT":
		return Location{"Sialkot", "Asia/Karachi"}
	case "SKU":
		return Location{"Skiros Island", "Europe/Athens"}
	case "SKV":
		return Location{"", "Africa/Cairo"}
	case "SKW":
		return Location{"Skwentna", "America/Anchorage"}
	case "SKX":
		return Location{"Saransk", "Europe/Moscow"}
	case "SKZ":
		return Location{"Mirpur Khas", "Asia/Karachi"}
	case "SLA":
		return Location{"Salta", "America/Argentina/Salta"}
	case "SLB":
		return Location{"Storm Lake", "America/Chicago"}
	case "SLC":
		return Location{"Salt Lake City", "America/Denver"}
	case "SLD":
		return Location{"Sliac", "Europe/Bratislava"}
	case "SLE":
		return Location{"Salem", "America/Los_Angeles"}
	case "SLF":
		return Location{"", "Asia/Riyadh"}
	case "SLG":
		return Location{"Siloam Springs", "America/Chicago"}
	case "SLH":
		return Location{"Sola", "Pacific/Efate"}
	case "SLI":
		return Location{"Solwesi", "Africa/Lusaka"}
	case "SLJ":
		return Location{"Karijini National Park", "Australia/Perth"}
	case "SLK":
		return Location{"Saranac Lake", "America/New_York"}
	case "SLL":
		return Location{"Salalah", "Asia/Muscat"}
	case "SLM":
		return Location{"Salamanca", "Europe/Madrid"}
	case "SLN":
		return Location{"Salina", "America/Chicago"}
	case "SLO":
		return Location{"Salem", "America/Chicago"}
	case "SLP":
		return Location{"San Luis Potosi", "America/Mexico_City"}
	case "SLQ":
		return Location{"Sleetmute", "America/Anchorage"}
	case "SLR":
		return Location{"Sulphur Springs", "America/Chicago"}
	case "SLS":
		return Location{"Silistra", "Europe/Sofia"}
	case "SLT":
		return Location{"Salida", "America/Denver"}
	case "SLU":
		return Location{"Castries", "America/St_Lucia"}
	case "SLV":
		return Location{"", "Asia/Kolkata"}
	case "SLW":
		return Location{"Saltillo", "America/Monterrey"}
	case "SLX":
		return Location{"Salt Cay", "America/Grand_Turk"}
	case "SLY":
		return Location{"Salekhard", "Asia/Yekaterinburg"}
	case "SLZ":
		return Location{"Sao Luis", "America/Fortaleza"}
	case "SMA":
		return Location{"Vila do Porto", "Atlantic/Azores"}
	case "SMB":
		return Location{"Cerro Sombrero", "America/Punta_Arenas"}
	case "SMD":
		return Location{"Fort Wayne", "America/Indiana/Indianapolis"}
	case "SME":
		return Location{"Somerset", "America/New_York"}
	case "SMF":
		return Location{"Sacramento", "America/Los_Angeles"}
	case "SMG":
		return Location{"Santa Maria", "America/Lima"}
	case "SMI":
		return Location{"Samos Island", "Europe/Athens"}
	case "SMK":
		return Location{"St Michael", "America/Nome"}
	case "SML":
		return Location{"Stella Maris", "America/Nassau"}
	case "SMM":
		return Location{"Semporna", "Asia/Kuching"}
	case "SMN":
		return Location{"Salmon", "America/Boise"}
	case "SMO":
		return Location{"Santa Monica", "America/Los_Angeles"}
	case "SMQ":
		return Location{"Sampit-Borneo Island", "Asia/Pontianak"}
	case "SMR":
		return Location{"Santa Marta", "America/Bogota"}
	case "SMS":
		return Location{"", "Indian/Antananarivo"}
	case "SMU":
		return Location{"Sheep Mountain", "America/Anchorage"}
	case "SMV":
		return Location{"", "Europe/Zurich"}
	case "SMW":
		return Location{"Smara", "Africa/El_Aaiun"}
	case "SMX":
		return Location{"Santa Maria", "America/Los_Angeles"}
	case "SMY":
		return Location{"Simenti", "Africa/Dakar"}
	case "SMZ":
		return Location{"Stoelmanseiland", "America/Paramaribo"}
	case "SNA":
		return Location{"Santa Ana", "America/Los_Angeles"}
	case "SNB":
		return Location{"", "Australia/Darwin"}
	case "SNC":
		return Location{"Salinas", "America/Guayaquil"}
	case "SNE":
		return Location{"Preguica", "Atlantic/Cape_Verde"}
	case "SNF":
		return Location{"San Felipe", "America/Caracas"}
	case "SNG":
		return Location{"San Ignacio de Velasco", "America/La_Paz"}
	case "SNH":
		return Location{"", "Australia/Brisbane"}
	case "SNI":
		return Location{"Greenville", "Africa/Monrovia"}
	case "SNJ":
		return Location{"Pinar Del Rio", "America/Havana"}
	case "SNK":
		return Location{"Snyder", "America/Chicago"}
	case "SNL":
		return Location{"Shawnee", "America/Chicago"}
	case "SNM":
		return Location{"San Ignacio de Moxos", "America/La_Paz"}
	case "SNN":
		return Location{"Shannon", "Europe/Dublin"}
	case "SNO":
		return Location{"", "Asia/Bangkok"}
	case "SNP":
		return Location{"St Paul Island", "America/Nome"}
	case "SNR":
		return Location{"Saint-Nazaire/Montoir", "Europe/Paris"}
	case "SNS":
		return Location{"Salinas", "America/Los_Angeles"}
	case "SNU":
		return Location{"Santa Clara", "America/Havana"}
	case "SNV":
		return Location{"", "America/Caracas"}
	case "SNW":
		return Location{"Thandwe", "Asia/Yangon"}
	case "SNX":
		return Location{"Semnan", "Asia/Tehran"}
	case "SNY":
		return Location{"Sidney", "America/Denver"}
	case "SNZ":
		return Location{"Rio De Janeiro", "America/Sao_Paulo"}
	case "SOB":
		return Location{"Sarmellek", "Europe/Budapest"}
	case "SOC":
		return Location{"Sukarata(Solo)-Java Island", "Asia/Jakarta"}
	case "SOD":
		return Location{"Sorocaba", "America/Sao_Paulo"}
	case "SOE":
		return Location{"Souanke", "Africa/Brazzaville"}
	case "SOF":
		return Location{"Sofia", "Europe/Sofia"}
	case "SOG":
		return Location{"Sogndal", "Europe/Oslo"}
	case "SOJ":
		return Location{"Sorkjosen", "Europe/Oslo"}
	case "SOK":
		return Location{"Semonkong", "Africa/Maseru"}
	case "SOM":
		return Location{"", "America/Caracas"}
	case "SON":
		return Location{"Luganville", "Pacific/Efate"}
	case "SOO":
		return Location{"Soderhamn", "Europe/Stockholm"}
	case "SOP":
		return Location{"Pinehurst/Southern Pines", "America/New_York"}
	case "SOQ":
		return Location{"Sorong-Papua Island", "Asia/Jayapura"}
	case "SOT":
		return Location{"Sodankyla", "Europe/Helsinki"}
	case "SOU":
		return Location{"Southampton", "Europe/London"}
	case "SOV":
		return Location{"Seldovia", "America/Anchorage"}
	case "SOW":
		return Location{"Show Low", "America/Phoenix"}
	case "SOX":
		return Location{"Sogamoso", "America/Bogota"}
	case "SOY":
		return Location{"Stronsay", "Europe/London"}
	case "SOZ":
		return Location{"Solenzara", "Europe/Paris"}
	case "SPA":
		return Location{"Spartanburg", "America/New_York"}
	case "SPC":
		return Location{"Sta Cruz de la Palma", "Atlantic/Canary"}
	case "SPD":
		return Location{"Saidpur", "Asia/Dhaka"}
	case "SPE":
		return Location{"Sepulot", "Asia/Kuching"}
	case "SPF":
		return Location{"Spearfish", "America/Denver"}
	case "SPG":
		return Location{"St Petersburg", "America/New_York"}
	case "SPI":
		return Location{"Springfield", "America/Chicago"}
	case "SPJ":
		return Location{"Sparti", "Europe/Athens"}
	case "SPM":
		return Location{"Trier", "Europe/Berlin"}
	case "SPN":
		return Location{"Saipan Island", "Pacific/Saipan"}
	case "SPP":
		return Location{"Menongue", "Africa/Luanda"}
	case "SPS":
		return Location{"Wichita Falls", "America/Chicago"}
	case "SPU":
		return Location{"Split", "Europe/Zagreb"}
	case "SPW":
		return Location{"Spencer", "America/Chicago"}
	case "SPX":
		return Location{"Giza", "Africa/Cairo"}
	case "SPY":
		return Location{"", "Africa/Abidjan"}
	case "SPZ":
		return Location{"Springdale", "America/Chicago"}
	case "SQA":
		return Location{"Santa Ynez", "America/Los_Angeles"}
	case "SQC":
		return Location{"", "Australia/Perth"}
	case "SQD":
		return Location{"San Francisco de Yeso", "America/Lima"}
	case "SQG":
		return Location{"Sintang-Borneo Island", "Asia/Pontianak"}
	case "SQH":
		return Location{"Son-La", "Asia/Bangkok"}
	case "SQI":
		return Location{"Sterling/Rockfalls", "America/Chicago"}
	case "SQJ":
		return Location{"Sanming", "Asia/Shanghai"}
	case "SQL":
		return Location{"San Carlos", "America/Los_Angeles"}
	case "SQM":
		return Location{"Sao Miguel Do Araguaia", "America/Sao_Paulo"}
	case "SQN":
		return Location{"Sanana-Seram Island", "Asia/Jayapura"}
	case "SQO":
		return Location{"", "Europe/Stockholm"}
	case "SQQ":
		return Location{"Siauliai", "Europe/Vilnius"}
	case "SQR":
		return Location{"Soroako-Celebes Island", "Asia/Makassar"}
	case "SQU":
		return Location{"Plaza Saposoa", "America/Lima"}
	case "SQW":
		return Location{"Skive", "Europe/Copenhagen"}
	case "SQX":
		return Location{"Sao Miguel Do Oeste", "America/Sao_Paulo"}
	case "SQY":
		return Location{"Sao Lourenco Do Sul", "America/Sao_Paulo"}
	case "SQZ":
		return Location{"Scampton", "Europe/London"}
	case "SRA":
		return Location{"Santa Rosa", "America/Sao_Paulo"}
	case "SRB":
		return Location{"Santa Rosa", "America/La_Paz"}
	case "SRC":
		return Location{"Searcy", "America/Chicago"}
	case "SRD":
		return Location{"San Ramon / Mamore", "America/La_Paz"}
	case "SRE":
		return Location{"Sucre", "America/La_Paz"}
	case "SRG":
		return Location{"Semarang-Java Island", "Asia/Jakarta"}
	case "SRH":
		return Location{"Sarh", "Africa/Ndjamena"}
	case "SRJ":
		return Location{"San Borja", "America/La_Paz"}
	case "SRN":
		return Location{"", "Australia/Hobart"}
	case "SRP":
		return Location{"Leirvik", "Europe/Oslo"}
	case "SRQ":
		return Location{"Sarasota/Bradenton", "America/New_York"}
	case "SRT":
		return Location{"Soroti", "Africa/Kampala"}
	case "SRW":
		return Location{"Salisbury", "America/New_York"}
	case "SRX":
		return Location{"Sirt", "Africa/Tripoli"}
	case "SRY":
		return Location{"Sari", "Asia/Tehran"}
	case "SRZ":
		return Location{"Santa Cruz", "America/La_Paz"}
	case "SSA":
		return Location{"Salvador", "America/Bahia"}
	case "SSC":
		return Location{"Sumter", "America/New_York"}
	case "SSD":
		return Location{"San Felipe", "America/Santiago"}
	case "SSE":
		return Location{"Solapur", "Asia/Kolkata"}
	case "SSF":
		return Location{"San Antonio", "America/Chicago"}
	case "SSG":
		return Location{"Malabo", "Africa/Malabo"}
	case "SSH":
		return Location{"Sharm el-Sheikh", "Africa/Cairo"}
	case "SSI":
		return Location{"Brunswick", "America/New_York"}
	case "SSJ":
		return Location{"Alstahaug", "Europe/Oslo"}
	case "SSN":
		return Location{"", "Asia/Seoul"}
	case "SSO":
		return Location{"Sao Lourenco", "America/Sao_Paulo"}
	case "SSR":
		return Location{"Pentecost Island", "Pacific/Efate"}
	case "SST":
		return Location{"Santa Teresita", "America/Argentina/Buenos_Aires"}
	case "SSX":
		return Location{"Samsun", "Europe/Istanbul"}
	case "SSY":
		return Location{"Mbanza Congo", "Africa/Luanda"}
	case "SSZ":
		return Location{"Guaruja", "America/Sao_Paulo"}
	case "STA":
		return Location{"Skjern / Ringkobing", "Europe/Copenhagen"}
	case "STB":
		return Location{"", "America/Caracas"}
	case "STC":
		return Location{"St Cloud", "America/Chicago"}
	case "STD":
		return Location{"Santo Domingo", "America/Caracas"}
	case "STE":
		return Location{"Stevens Point", "America/Chicago"}
	case "STG":
		return Location{"St George", "America/Nome"}
	case "STH":
		return Location{"", "Australia/Brisbane"}
	case "STI":
		return Location{"Santiago", "America/Santo_Domingo"}
	case "STJ":
		return Location{"St Joseph", "America/Chicago"}
	case "STK":
		return Location{"Sterling", "America/Denver"}
	case "STL":
		return Location{"St Louis", "America/Chicago"}
	case "STM":
		return Location{"Santarem", "America/Santarem"}
	case "STN":
		return Location{"London", "Europe/London"}
	case "STP":
		return Location{"St Paul", "America/Chicago"}
	case "STQ":
		return Location{"St Marys", "America/New_York"}
	case "STR":
		return Location{"Stuttgart", "Europe/Berlin"}
	case "STS":
		return Location{"Santa Rosa", "America/Los_Angeles"}
	case "STT":
		return Location{"Charlotte Amalie", "America/St_Thomas"}
	case "STV":
		return Location{"", "Asia/Kolkata"}
	case "STW":
		return Location{"Stavropol", "Europe/Moscow"}
	case "STX":
		return Location{"Christiansted", "America/St_Thomas"}
	case "STY":
		return Location{"Salto", "America/Montevideo"}
	case "STZ":
		return Location{"Santa Terezinha", "America/Araguaina"}
	case "SUA":
		return Location{"Stuart", "America/New_York"}
	case "SUB":
		return Location{"Surabaya", "Asia/Jakarta"}
	case "SUD":
		return Location{"Stroud", "America/Chicago"}
	case "SUE":
		return Location{"Sturgeon Bay", "America/Chicago"}
	case "SUF":
		return Location{"Lamezia Terme", "Europe/Rome"}
	case "SUG":
		return Location{"Surigao City", "Asia/Manila"}
	case "SUH":
		return Location{"Sur", "Asia/Muscat"}
	case "SUI":
		return Location{"Sukhumi", "Asia/Tbilisi"}
	case "SUJ":
		return Location{"Satu Mare", "Europe/Bucharest"}
	case "SUL":
		return Location{"Sui", "Asia/Karachi"}
	case "SUM":
		return Location{"Sumter", "America/New_York"}
	case "SUN":
		return Location{"Hailey", "America/Boise"}
	case "SUO":
		return Location{"Sunriver", "America/Los_Angeles"}
	case "SUP":
		return Location{"Sumenep-Madura Island", "Asia/Jakarta"}
	case "SUQ":
		return Location{"Sucua", "America/Guayaquil"}
	case "SUR":
		return Location{"Summer Beaver", "America/Toronto"}
	case "SUS":
		return Location{"St Louis", "America/Chicago"}
	case "SUT":
		return Location{"Sumbawanga", "Africa/Dar_es_Salaam"}
	case "SUU":
		return Location{"Fairfield", "America/Los_Angeles"}
	case "SUV":
		return Location{"Nausori", "Pacific/Fiji"}
	case "SUW":
		return Location{"Superior", "America/Chicago"}
	case "SUX":
		return Location{"Sioux City", "America/Chicago"}
	case "SUY":
		return Location{"Suntar", "Asia/Yakutsk"}
	case "SVA":
		return Location{"Savoonga", "America/Nome"}
	case "SVB":
		return Location{"", "Indian/Antananarivo"}
	case "SVC":
		return Location{"Silver City", "America/Denver"}
	case "SVD":
		return Location{"Argyle", "America/St_Vincent"}
	case "SVE":
		return Location{"Susanville", "America/Los_Angeles"}
	case "SVF":
		return Location{"Save", "Africa/Porto-Novo"}
	case "SVG":
		return Location{"Stavanger", "Europe/Oslo"}
	case "SVH":
		return Location{"Statesville", "America/New_York"}
	case "SVI":
		return Location{"San Vicente Del Caguan", "America/Bogota"}
	case "SVJ":
		return Location{"Svolvaer", "Europe/Oslo"}
	case "SVL":
		return Location{"Savonlinna", "Europe/Helsinki"}
	case "SVN":
		return Location{"Savannah", "America/New_York"}
	case "SVO":
		return Location{"Moscow", "Europe/Moscow"}
	case "SVP":
		return Location{"Kuito", "Africa/Luanda"}
	case "SVQ":
		return Location{"Sevilla", "Europe/Madrid"}
	case "SVT":
		return Location{"Savuti", "Africa/Gaborone"}
	case "SVU":
		return Location{"Savusavu", "Pacific/Fiji"}
	case "SVW":
		return Location{"Sparrevohn", "America/Anchorage"}
	case "SVX":
		return Location{"Yekaterinburg", "Asia/Yekaterinburg"}
	case "SVZ":
		return Location{"", "America/Bogota"}
	case "SWA":
		return Location{"Shantou", "Asia/Shanghai"}
	case "SWC":
		return Location{"", "Australia/Melbourne"}
	case "SWD":
		return Location{"Seward", "America/Anchorage"}
	case "SWF":
		return Location{"Newburgh", "America/New_York"}
	case "SWH":
		return Location{"", "Australia/Melbourne"}
	case "SWJ":
		return Location{"Malekula Island", "Pacific/Efate"}
	case "SWN":
		return Location{"Sahiwal", "Asia/Karachi"}
	case "SWO":
		return Location{"Stillwater", "America/Chicago"}
	case "SWP":
		return Location{"Swakopmund", "Africa/Windhoek"}
	case "SWQ":
		return Location{"Sumbawa Island", "Asia/Makassar"}
	case "SWS":
		return Location{"Swansea", "Europe/London"}
	case "SWT":
		return Location{"Strezhevoy", "Asia/Tomsk"}
	case "SWU":
		return Location{"", "Asia/Seoul"}
	case "SWV":
		return Location{"Evensk", "Asia/Magadan"}
	case "SWW":
		return Location{"Sweetwater", "America/Chicago"}
	case "SWX":
		return Location{"Shakawe", "Africa/Gaborone"}
	case "SWY":
		return Location{"Sitiawan", "Asia/Kuala_Lumpur"}
	case "SXB":
		return Location{"Strasbourg", "Europe/Paris"}
	case "SXE":
		return Location{"West Sale", "Australia/Melbourne"}
	case "SXG":
		return Location{"Senanga", "Africa/Lusaka"}
	case "SXI":
		return Location{"", "Asia/Tehran"}
	case "SXJ":
		return Location{"Shanshan", "Asia/Shanghai"}
	case "SXK":
		return Location{"Saumlaki-Yamdena Island", "Asia/Jayapura"}
	case "SXL":
		return Location{"Sligo", "Europe/Dublin"}
	case "SXM":
		return Location{"Saint Martin", "America/Lower_Princes"}
	case "SXN":
		return Location{"", "Africa/Gaborone"}
	case "SXO":
		return Location{"Sao Felix Do Araguaia", "America/Cuiaba"}
	case "SXQ":
		return Location{"Soldotna", "America/Anchorage"}
	case "SXR":
		return Location{"Srinagar", "Asia/Kolkata"}
	case "SXS":
		return Location{"Sahabat", "Asia/Kuching"}
	case "SXT":
		return Location{"Taman Negara", "Asia/Kuala_Lumpur"}
	case "SXU":
		return Location{"Soddu", "Africa/Addis_Ababa"}
	case "SXV":
		return Location{"", "Asia/Kolkata"}
	case "SXX":
		return Location{"Sao Felix Do Xingu", "America/Belem"}
	case "SXZ":
		return Location{"Siirt", "Europe/Istanbul"}
	case "SYA":
		return Location{"Shemya", "America/Adak"}
	case "SYC":
		return Location{"Shiringayoc", "America/Lima"}
	case "SYD":
		return Location{"Sydney", "Australia/Sydney"}
	case "SYE":
		return Location{"Sadah", "Asia/Aden"}
	case "SYI":
		return Location{"Shelbyville", "America/Chicago"}
	case "SYJ":
		return Location{"", "Asia/Tehran"}
	case "SYK":
		return Location{"Stykkisholmur", "Atlantic/Reykjavik"}
	case "SYM":
		return Location{"Simao", "Asia/Shanghai"}
	case "SYN":
		return Location{"Stanton", "America/Chicago"}
	case "SYO":
		return Location{"Shonai", "Asia/Tokyo"}
	case "SYP":
		return Location{"Santiago", "America/Panama"}
	case "SYQ":
		return Location{"San Jose", "America/Costa_Rica"}
	case "SYR":
		return Location{"Syracuse", "America/New_York"}
	case "SYS":
		return Location{"Saskylakh", "Asia/Yakutsk"}
	case "SYT":
		return Location{"Saint-Yan", "Europe/Paris"}
	case "SYU":
		return Location{"Sue Islet", "Australia/Brisbane"}
	case "SYV":
		return Location{"Sylvester", "America/New_York"}
	case "SYW":
		return Location{"Sehwan Sharif", "Asia/Karachi"}
	case "SYX":
		return Location{"Sanya", "Asia/Shanghai"}
	case "SYY":
		return Location{"Stornoway", "Europe/London"}
	case "SYZ":
		return Location{"Shiraz", "Asia/Tehran"}
	case "SZA":
		return Location{"Soyo", "Africa/Luanda"}
	case "SZB":
		return Location{"Subang", "Asia/Kuala_Lumpur"}
	case "SZE":
		return Location{"Semera", "Africa/Addis_Ababa"}
	case "SZF":
		return Location{"Samsun", "Europe/Istanbul"}
	case "SZG":
		return Location{"Salzburg", "Europe/Berlin"}
	case "SZH":
		return Location{"Shuozhou", "Asia/Shanghai"}
	case "SZJ":
		return Location{"Isla de la Juventud", "America/Havana"}
	case "SZK":
		return Location{"Skukuza", "Africa/Johannesburg"}
	case "SZL":
		return Location{"Knob Noster", "America/Chicago"}
	case "SZM":
		return Location{"", "Africa/Windhoek"}
	case "SZN":
		return Location{"Santa Barbara", "America/Los_Angeles"}
	case "SZP":
		return Location{"Santa Paula", "America/Los_Angeles"}
	case "SZR":
		return Location{"Stara Zagora", "Europe/Sofia"}
	case "SZS":
		return Location{"Oban", "Pacific/Auckland"}
	case "SZT":
		return Location{"", "America/Mexico_City"}
	case "SZV":
		return Location{"Suzhou", "Asia/Shanghai"}
	case "SZW":
		return Location{"", "Europe/Berlin"}
	case "SZX":
		return Location{"Shenzhen", "Asia/Shanghai"}
	case "SZY":
		return Location{"Szymany", "Europe/Warsaw"}
	case "SZZ":
		return Location{"Goleniow", "Europe/Warsaw"}
	case "TAB":
		return Location{"Scarborough", "America/Port_of_Spain"}
	case "TAC":
		return Location{"Tacloban City", "Asia/Manila"}
	case "TAD":
		return Location{"Trinidad", "America/Denver"}
	case "TAE":
		return Location{"Daegu", "Asia/Seoul"}
	case "TAF":
		return Location{"", "Africa/Algiers"}
	case "TAG":
		return Location{"Panglao", "Asia/Manila"}
	case "TAH":
		return Location{"", "Pacific/Efate"}
	case "TAI":
		return Location{"Ta'izz", "Asia/Aden"}
	case "TAJ":
		return Location{"Aitape", "Pacific/Port_Moresby"}
	case "TAK":
		return Location{"Takamatsu", "Asia/Tokyo"}
	case "TAL":
		return Location{"Tanana", "America/Anchorage"}
	case "TAM":
		return Location{"Tampico", "America/Monterrey"}
	case "TAN":
		return Location{"", "Australia/Brisbane"}
	case "TAO":
		return Location{"Qingdao", "Asia/Shanghai"}
	case "TAP":
		return Location{"Tapachula", "America/Mexico_City"}
	case "TAQ":
		return Location{"Tarcoola", "Australia/Adelaide"}
	case "TAR":
		return Location{"Grottaglie", "Europe/Rome"}
	case "TAS":
		return Location{"Tashkent", "Asia/Tashkent"}
	case "TAT":
		return Location{"Poprad", "Europe/Bratislava"}
	case "TAU":
		return Location{"Tauramena", "America/Bogota"}
	case "TAW":
		return Location{"Tacuarembo", "America/Montevideo"}
	case "TAX":
		return Location{"Tikong-Taliabu Island", "Asia/Jayapura"}
	case "TAY":
		return Location{"Tartu", "Europe/Tallinn"}
	case "TAZ":
		return Location{"Dashoguz", "Asia/Ashgabat"}
	case "TBB":
		return Location{"Tuy Hoa", "Asia/Ho_Chi_Minh"}
	case "TBF":
		return Location{"", "Pacific/Tarawa"}
	case "TBG":
		return Location{"Tabubil", "Pacific/Port_Moresby"}
	case "TBH":
		return Location{"Romblon", "Asia/Manila"}
	case "TBI":
		return Location{"Cat Island", "America/Nassau"}
	case "TBJ":
		return Location{"Tabarka", "Africa/Tunis"}
	case "TBK":
		return Location{"", "Australia/Darwin"}
	case "TBL":
		return Location{"", "Australia/Perth"}
	case "TBN":
		return Location{"Fort Leonard Wood", "America/Chicago"}
	case "TBO":
		return Location{"Tabora", "Africa/Dar_es_Salaam"}
	case "TBP":
		return Location{"Tumbes", "America/Lima"}
	case "TBR":
		return Location{"Statesboro", "America/New_York"}
	case "TBS":
		return Location{"Tbilisi", "Asia/Tbilisi"}
	case "TBT":
		return Location{"Tabatinga", "America/Bogota"}
	case "TBU":
		return Location{"Nuku'alofa", "Pacific/Tongatapu"}
	case "TBW":
		return Location{"Tambov", "Europe/Moscow"}
	case "TBY":
		return Location{"Tshabong", "Africa/Gaborone"}
	case "TBZ":
		return Location{"Tabriz", "Asia/Tehran"}
	case "TCA":
		return Location{"Tennant Creek", "Australia/Darwin"}
	case "TCB":
		return Location{"Treasure Cay", "America/Nassau"}
	case "TCC":
		return Location{"Tucumcari", "America/Denver"}
	case "TCE":
		return Location{"Tulcea", "Europe/Bucharest"}
	case "TCG":
		return Location{"Tacheng", "Asia/Shanghai"}
	case "TCH":
		return Location{"Tchibanga", "Africa/Libreville"}
	case "TCL":
		return Location{"Tuscaloosa", "America/Chicago"}
	case "TCM":
		return Location{"Tacoma", "America/Los_Angeles"}
	case "TCN":
		return Location{"", "America/Mexico_City"}
	case "TCO":
		return Location{"Tumaco", "America/Bogota"}
	case "TCP":
		return Location{"Taba", "Africa/Cairo"}
	case "TCQ":
		return Location{"Tacna", "America/Lima"}
	case "TCR":
		return Location{"Thoothukkudi", "Asia/Kolkata"}
	case "TCS":
		return Location{"Truth Or Consequences", "America/Denver"}
	case "TCU":
		return Location{"Homeward", "Africa/Johannesburg"}
	case "TCW":
		return Location{"", "Australia/Sydney"}
	case "TCX":
		return Location{"Tabas", "Asia/Tehran"}
	case "TCZ":
		return Location{"Tengchong", "Asia/Shanghai"}
	case "TDA":
		return Location{"Trinidad", "America/Bogota"}
	case "TDD":
		return Location{"Trinidad", "America/La_Paz"}
	case "TDG":
		return Location{"", "Asia/Manila"}
	case "TDJ":
		return Location{"Tadjoura", "Africa/Djibouti"}
	case "TDK":
		return Location{"Taldy Kurgan", "Asia/Almaty"}
	case "TDL":
		return Location{"Tandil", "America/Argentina/Buenos_Aires"}
	case "TDN":
		return Location{"", "Australia/Perth"}
	case "TDO":
		return Location{"Toledo", "America/Los_Angeles"}
	case "TDP":
		return Location{"Corrientes", "America/Lima"}
	case "TDR":
		return Location{"", "Australia/Brisbane"}
	case "TDS":
		return Location{"Sasereme", "Pacific/Port_Moresby"}
	case "TDT":
		return Location{"Welverdiend", "Africa/Johannesburg"}
	case "TDV":
		return Location{"Tanandava", "Indian/Antananarivo"}
	case "TDW":
		return Location{"Amarillo", "America/Chicago"}
	case "TDX":
		return Location{"", "Asia/Bangkok"}
	case "TDZ":
		return Location{"Toledo", "America/New_York"}
	case "TEA":
		return Location{"Tela", "America/Tegucigalpa"}
	case "TEB":
		return Location{"Teterboro", "America/New_York"}
	case "TEC":
		return Location{"Telemaco Borba", "America/Sao_Paulo"}
	case "TED":
		return Location{"Thisted", "Europe/Copenhagen"}
	case "TEE":
		return Location{"Tebessi", "Africa/Algiers"}
	case "TEF":
		return Location{"", "Australia/Perth"}
	case "TEG":
		return Location{"Tenkodogo", "Africa/Ouagadougou"}
	case "TEI":
		return Location{"Tezu", "Asia/Kolkata"}
	case "TEL":
		return Location{"Telupid", "Asia/Kuching"}
	case "TEM":
		return Location{"", "Australia/Sydney"}
	case "TEN":
		return Location{"", "Asia/Shanghai"}
	case "TEQ":
		return Location{"Corlu", "Europe/Istanbul"}
	case "TER":
		return Location{"Lajes", "Atlantic/Azores"}
	case "TES":
		return Location{"Tessenei", "Africa/Asmara"}
	case "TET":
		return Location{"Tete", "Africa/Maputo"}
	case "TEU":
		return Location{"", "Pacific/Auckland"}
	case "TEV":
		return Location{"Teruel", "Europe/Madrid"}
	case "TEX":
		return Location{"Telluride", "America/Denver"}
	case "TEY":
		return Location{"Tingeyri", "Atlantic/Reykjavik"}
	case "TEZ":
		return Location{"", "Asia/Kolkata"}
	case "TFF":
		return Location{"Tefe", "America/Manaus"}
	case "TFI":
		return Location{"Tufi", "Pacific/Port_Moresby"}
	case "TFL":
		return Location{"Teofilo Otoni", "America/Sao_Paulo"}
	case "TFM":
		return Location{"Telefomin", "Pacific/Port_Moresby"}
	case "TFN":
		return Location{"Tenerife Island", "Atlantic/Canary"}
	case "TFS":
		return Location{"Tenerife Island", "Atlantic/Canary"}
	case "TFT":
		return Location{"Taftan", "Asia/Karachi"}
	case "TFU":
		return Location{"Chengdu", "Asia/Shanghai"}
	case "TGA":
		return Location{"", "Asia/Singapore"}
	case "TGC":
		return Location{"Tanjung Manis", "Asia/Kuching"}
	case "TGD":
		return Location{"Podgorica", "Europe/Podgorica"}
	case "TGG":
		return Location{"Kuala Terengganu", "Asia/Kuala_Lumpur"}
	case "TGH":
		return Location{"Tongoa Island", "Pacific/Efate"}
	case "TGI":
		return Location{"Tingo Maria", "America/Lima"}
	case "TGJ":
		return Location{"Tiga", "Pacific/Noumea"}
	case "TGK":
		return Location{"Taganrog", "Europe/Moscow"}
	case "TGM":
		return Location{"Targu Mures", "Europe/Bucharest"}
	case "TGN":
		return Location{"", "Australia/Melbourne"}
	case "TGO":
		return Location{"Tongliao", "Asia/Shanghai"}
	case "TGP":
		return Location{"Bor", "Asia/Krasnoyarsk"}
	case "TGQ":
		return Location{"Tangara Da Serra", "America/Cuiaba"}
	case "TGR":
		return Location{"Touggourt", "Africa/Algiers"}
	case "TGT":
		return Location{"Tanga", "Africa/Dar_es_Salaam"}
	case "TGU":
		return Location{"Tegucigalpa", "America/Tegucigalpa"}
	case "TGV":
		return Location{"Targovishte", "Europe/Sofia"}
	case "TGZ":
		return Location{"Tuxtla Gutierrez", "America/Mexico_City"}
	case "THA":
		return Location{"Tullahoma", "America/Chicago"}
	case "THB":
		return Location{"Thaba-Tseka", "Africa/Maseru"}
	case "THC":
		return Location{"Tchien", "Africa/Monrovia"}
	case "THD":
		return Location{"Thanh Hóa", "Asia/Bangkok"}
	case "THE":
		return Location{"Teresina", "America/Fortaleza"}
	case "THF":
		return Location{"Berlin", "Europe/Berlin"}
	case "THG":
		return Location{"Biloela", "Australia/Brisbane"}
	case "THH":
		return Location{"Taharoa", "Pacific/Auckland"}
	case "THI":
		return Location{"Tichitt", "Africa/Nouakchott"}
	case "THK":
		return Location{"Thakhek", "Asia/Bangkok"}
	case "THL":
		return Location{"Tachileik", "Asia/Yangon"}
	case "THM":
		return Location{"Thompson Falls", "America/Denver"}
	case "THN":
		return Location{"Trollhattan", "Europe/Stockholm"}
	case "THO":
		return Location{"Thorshofn", "Atlantic/Reykjavik"}
	case "THQ":
		return Location{"Tianshui", "Asia/Shanghai"}
	case "THR":
		return Location{"Tehran", "Asia/Tehran"}
	case "THS":
		return Location{"", "Asia/Bangkok"}
	case "THT":
		return Location{"Tamchakett", "Africa/Nouakchott"}
	case "THU":
		return Location{"Thule", "America/Thule"}
	case "THV":
		return Location{"York", "America/New_York"}
	case "THX":
		return Location{"Turukhansk", "Asia/Krasnoyarsk"}
	case "THY":
		return Location{"Thohoyandou", "Africa/Johannesburg"}
	case "THZ":
		return Location{"Tahoua", "Africa/Niamey"}
	case "TIA":
		return Location{"Tirana", "Europe/Tirane"}
	case "TIB":
		return Location{"Tibu", "America/Bogota"}
	case "TID":
		return Location{"Tiaret", "Africa/Algiers"}
	case "TIE":
		return Location{"Tippi", "Africa/Addis_Ababa"}
	case "TIF":
		return Location{"", "Asia/Riyadh"}
	case "TIH":
		return Location{"", "Pacific/Tahiti"}
	case "TII":
		return Location{"Tarin Kowt", "Asia/Kabul"}
	case "TIJ":
		return Location{"Tijuana", "America/Los_Angeles"}
	case "TIK":
		return Location{"Oklahoma City", "America/Chicago"}
	case "TIM":
		return Location{"Timika-Papua Island", "Asia/Jayapura"}
	case "TIN":
		return Location{"Tindouf", "Africa/Algiers"}
	case "TIO":
		return Location{"Tilin", "Asia/Yangon"}
	case "TIP":
		return Location{"Tripoli", "Africa/Tripoli"}
	case "TIQ":
		return Location{"Tinian Island", "Pacific/Saipan"}
	case "TIR":
		return Location{"Tirupati", "Asia/Kolkata"}
	case "TIU":
		return Location{"", "Pacific/Auckland"}
	case "TIV":
		return Location{"Tivat", "Europe/Podgorica"}
	case "TIW":
		return Location{"Tacoma", "America/Los_Angeles"}
	case "TIX":
		return Location{"Titusville", "America/New_York"}
	case "TIY":
		return Location{"Tidjikja", "Africa/Nouakchott"}
	case "TIZ":
		return Location{"Tari", "Pacific/Port_Moresby"}
	case "TJA":
		return Location{"Tarija", "America/La_Paz"}
	case "TJB":
		return Location{"Karinmunbesar Island", "Asia/Jakarta"}
	case "TJG":
		return Location{"Tanta-Tabalong-Borneo Island", "Asia/Makassar"}
	case "TJH":
		return Location{"Tajima", "Asia/Tokyo"}
	case "TJI":
		return Location{"Trujillo", "America/Tegucigalpa"}
	case "TJK":
		return Location{"Tokat", "Europe/Istanbul"}
	case "TJL":
		return Location{"Tres Lagoas", "America/Campo_Grande"}
	case "TJM":
		return Location{"Tyumen", "Asia/Yekaterinburg"}
	case "TJN":
		return Location{"Takume", "Pacific/Tahiti"}
	case "TJQ":
		return Location{"Tanjung Pandan-Belitung Island", "Asia/Jakarta"}
	case "TJS":
		return Location{"Tanjung Selor-Borneo Island", "Asia/Makassar"}
	case "TJU":
		return Location{"Kulyab", "Asia/Dushanbe"}
	case "TJV":
		return Location{"Thanjavur", "Asia/Kolkata"}
	case "TKA":
		return Location{"Talkeetna", "America/Anchorage"}
	case "TKC":
		return Location{"Tiko", "Africa/Douala"}
	case "TKD":
		return Location{"Sekondi-Takoradi", "Africa/Accra"}
	case "TKF":
		return Location{"Truckee", "America/Los_Angeles"}
	case "TKG":
		return Location{"Bandar Lampung-Sumatra Island", "Asia/Jakarta"}
	case "TKH":
		return Location{"", "Asia/Bangkok"}
	case "TKJ":
		return Location{"Tok", "America/Anchorage"}
	case "TKK":
		return Location{"Weno Island", "Pacific/Chuuk"}
	case "TKN":
		return Location{"Tokunoshima", "Asia/Tokyo"}
	case "TKO":
		return Location{"Tlokoeng", "Africa/Maseru"}
	case "TKP":
		return Location{"", "Pacific/Tahiti"}
	case "TKQ":
		return Location{"Kigoma", "Africa/Dar_es_Salaam"}
	case "TKR":
		return Location{"Thakurgaon", "Asia/Dhaka"}
	case "TKS":
		return Location{"Tokushima", "Asia/Tokyo"}
	case "TKT":
		return Location{"", "Asia/Bangkok"}
	case "TKU":
		return Location{"Turku", "Europe/Helsinki"}
	case "TKV":
		return Location{"Tatakoto", "Pacific/Tahiti"}
	case "TKW":
		return Location{"Tekin", "Pacific/Port_Moresby"}
	case "TKX":
		return Location{"", "Pacific/Tahiti"}
	case "TKY":
		return Location{"Turkey Creek", "Australia/Perth"}
	case "TKZ":
		return Location{"Tokoroa", "Pacific/Auckland"}
	case "TLA":
		return Location{"Teller", "America/Nome"}
	case "TLB":
		return Location{"Tarbela", "Asia/Karachi"}
	case "TLC":
		return Location{"Toluca", "America/Mexico_City"}
	case "TLD":
		return Location{"Tuli Lodge", "Africa/Gaborone"}
	case "TLE":
		return Location{"", "Indian/Antananarivo"}
	case "TLH":
		return Location{"Tallahassee", "America/New_York"}
	case "TLI":
		return Location{"Toli Toli-Celebes Island", "Asia/Makassar"}
	case "TLJ":
		return Location{"Takotna", "America/Anchorage"}
	case "TLK":
		return Location{"", "Asia/Yakutsk"}
	case "TLL":
		return Location{"Tallinn", "Europe/Tallinn"}
	case "TLM":
		return Location{"Tlemcen", "Africa/Algiers"}
	case "TLN":
		return Location{"Toulon/Hyeres/Le Palyvestre", "Europe/Paris"}
	case "TLQ":
		return Location{"Turpan", "Asia/Shanghai"}
	case "TLR":
		return Location{"Tulare", "America/Los_Angeles"}
	case "TLS":
		return Location{"Toulouse/Blagnac", "Europe/Paris"}
	case "TLU":
		return Location{"Tolu", "America/Bogota"}
	case "TLV":
		return Location{"Tel Aviv", "Asia/Jerusalem"}
	case "TLX":
		return Location{"Talca", "America/Santiago"}
	case "TLY":
		return Location{"Plastun", "Asia/Vladivostok"}
	case "TLZ":
		return Location{"Catalao", "America/Sao_Paulo"}
	case "TMA":
		return Location{"Tifton", "America/New_York"}
	case "TMB":
		return Location{"Miami", "America/New_York"}
	case "TMC":
		return Location{"Waikabubak-Sumba Island", "Asia/Makassar"}
	case "TMD":
		return Location{"Timbedra", "Africa/Nouakchott"}
	case "TME":
		return Location{"Tame", "America/Bogota"}
	case "TMF":
		return Location{"Thimarafushi", "Indian/Maldives"}
	case "TMG":
		return Location{"Tomanggong", "Asia/Kuching"}
	case "TMH":
		return Location{"Tanah Merah-Papua Island", "Asia/Jayapura"}
	case "TMI":
		return Location{"Tumling Tar", "Asia/Kathmandu"}
	case "TMJ":
		return Location{"Termez", "Asia/Samarkand"}
	case "TML":
		return Location{"Tamale", "Africa/Accra"}
	case "TMM":
		return Location{"", "Indian/Antananarivo"}
	case "TMN":
		return Location{"Tamana Island", "Pacific/Tarawa"}
	case "TMO":
		return Location{"", "America/Caracas"}
	case "TMP":
		return Location{"Tampere / Pirkkala", "Europe/Helsinki"}
	case "TMQ":
		return Location{"Tambao", "Africa/Ouagadougou"}
	case "TMR":
		return Location{"Tamanrasset", "Africa/Algiers"}
	case "TMS":
		return Location{"Sao Tome", "Africa/Sao_Tome"}
	case "TMT":
		return Location{"Oriximina", "America/Santarem"}
	case "TMU":
		return Location{"Nicoya", "America/Costa_Rica"}
	case "TMW":
		return Location{"Tamworth", "Australia/Sydney"}
	case "TMX":
		return Location{"Timimoun", "Africa/Algiers"}
	case "TMZ":
		return Location{"", "Pacific/Auckland"}
	case "TNA":
		return Location{"Jinan", "Asia/Shanghai"}
	case "TNB":
		return Location{"Tanah Grogot-Borneo Island", "Asia/Makassar"}
	case "TNC":
		return Location{"Tin City", "America/Nome"}
	case "TND":
		return Location{"Trinidad", "America/Havana"}
	case "TNE":
		return Location{"", "Asia/Tokyo"}
	case "TNF":
		return Location{"Toussus-le-Noble", "Europe/Paris"}
	case "TNG":
		return Location{"Tangier", "Africa/Casablanca"}
	case "TNH":
		return Location{"Tonghua", "Asia/Shanghai"}
	case "TNI":
		return Location{"", "Asia/Kolkata"}
	case "TNJ":
		return Location{"Tanjung Pinang-Bintan Island", "Asia/Jakarta"}
	case "TNL":
		return Location{"Ternopil", "Europe/Kyiv"}
	case "TNM":
		return Location{"Isla Rey Jorge", "America/Punta_Arenas"}
	case "TNN":
		return Location{"Tainan City", "Asia/Taipei"}
	case "TNO":
		return Location{"Santa Cruz", "America/Costa_Rica"}
	case "TNP":
		return Location{"Twentynine Palms", "America/Los_Angeles"}
	case "TNR":
		return Location{"Antananarivo", "Indian/Antananarivo"}
	case "TNT":
		return Location{"Miami", "America/New_York"}
	case "TNU":
		return Location{"Newton", "America/Chicago"}
	case "TNV":
		return Location{"Tabuaeran Island", "Pacific/Kiritimati"}
	case "TNX":
		return Location{"Stung Treng", "Asia/Phnom_Penh"}
	case "TNZ":
		return Location{"Tosontsengel", "Asia/Ulaanbaatar"}
	case "TOA":
		return Location{"Torrance", "America/Los_Angeles"}
	case "TOB":
		return Location{"Tobruk", "Africa/Tripoli"}
	case "TOC":
		return Location{"Toccoa", "America/New_York"}
	case "TOD":
		return Location{"Pulau Tioman", "Asia/Kuala_Lumpur"}
	case "TOE":
		return Location{"Tozeur", "Africa/Tunis"}
	case "TOF":
		return Location{"Tomsk", "Asia/Tomsk"}
	case "TOG":
		return Location{"Togiak Village", "America/Anchorage"}
	case "TOH":
		return Location{"Loh/Linua", "Pacific/Efate"}
	case "TOI":
		return Location{"Troy", "America/Chicago"}
	case "TOJ":
		return Location{"Madrid", "Europe/Madrid"}
	case "TOL":
		return Location{"Toledo", "America/New_York"}
	case "TOM":
		return Location{"Timbuktu", "Africa/Bamako"}
	case "TOO":
		return Location{"Coto Brus", "America/Costa_Rica"}
	case "TOP":
		return Location{"Topeka", "America/Chicago"}
	case "TOQ":
		return Location{"Tocopilla", "America/Santiago"}
	case "TOR":
		return Location{"Torrington", "America/Denver"}
	case "TOS":
		return Location{"Tromso", "Europe/Oslo"}
	case "TOT":
		return Location{"Totness", "America/Paramaribo"}
	case "TOU":
		return Location{"Touho", "Pacific/Noumea"}
	case "TOW":
		return Location{"Toledo", "America/Sao_Paulo"}
	case "TOX":
		return Location{"Tobolsk", "Asia/Yekaterinburg"}
	case "TOY":
		return Location{"Toyama", "Asia/Tokyo"}
	case "TPA":
		return Location{"Tampa", "America/New_York"}
	case "TPC":
		return Location{"Tarapoa", "America/Guayaquil"}
	case "TPE":
		return Location{"Taipei", "Asia/Taipei"}
	case "TPF":
		return Location{"Tampa", "America/New_York"}
	case "TPG":
		return Location{"Taiping", "Asia/Kuala_Lumpur"}
	case "TPH":
		return Location{"Tonopah", "America/Los_Angeles"}
	case "TPI":
		return Location{"Tapini", "Pacific/Port_Moresby"}
	case "TPJ":
		return Location{"Taplejung", "Asia/Kathmandu"}
	case "TPK":
		return Location{"Tapak Tuan-Sumatra Island", "Asia/Jakarta"}
	case "TPL":
		return Location{"Temple", "America/Chicago"}
	case "TPN":
		return Location{"Tiputini", "America/Guayaquil"}
	case "TPP":
		return Location{"Tarapoto", "America/Lima"}
	case "TPQ":
		return Location{"Tepic", "America/Mazatlan"}
	case "TPR":
		return Location{"Tom Price", "Australia/Perth"}
	case "TPS":
		return Location{"Trapani", "Europe/Rome"}
	case "TPU":
		return Location{"Tikapur", "Asia/Kathmandu"}
	case "TQD":
		return Location{"Al Habbaniyah", "Asia/Baghdad"}
	case "TQL":
		return Location{"Tarko-Sale", "Asia/Yekaterinburg"}
	case "TQN":
		return Location{"Taloqan", "Asia/Kabul"}
	case "TQO":
		return Location{"Tulum", "America/Cancun"}
	case "TQP":
		return Location{"", "Australia/Brisbane"}
	case "TQQ":
		return Location{"Waha-Tomea Island", "Asia/Makassar"}
	case "TQS":
		return Location{"Tres Esquinas", "America/Bogota"}
	case "TRA":
		return Location{"", "Asia/Tokyo"}
	case "TRB":
		return Location{"Turbo", "America/Bogota"}
	case "TRC":
		return Location{"Torreon", "America/Monterrey"}
	case "TRD":
		return Location{"Trondheim", "Europe/Oslo"}
	case "TRE":
		return Location{"Balemartine", "Europe/London"}
	case "TRF":
		return Location{"Torp", "Europe/Oslo"}
	case "TRG":
		return Location{"Tauranga", "Pacific/Auckland"}
	case "TRI":
		return Location{"Bristol/Johnson/Kingsport", "America/New_York"}
	case "TRK":
		return Location{"Tarakan Island", "Asia/Makassar"}
	case "TRL":
		return Location{"Terrell", "America/Chicago"}
	case "TRM":
		return Location{"Palm Springs", "America/Los_Angeles"}
	case "TRN":
		return Location{"Torino", "Europe/Rome"}
	case "TRO":
		return Location{"Taree", "Australia/Sydney"}
	case "TRQ":
		return Location{"Tarauaca", "America/Rio_Branco"}
	case "TRR":
		return Location{"Trincomalee", "Asia/Colombo"}
	case "TRS":
		return Location{"Trieste", "Europe/Rome"}
	case "TRU":
		return Location{"Trujillo", "America/Lima"}
	case "TRV":
		return Location{"Trivandrum", "Asia/Kolkata"}
	case "TRW":
		return Location{"Tarawa", "Pacific/Tarawa"}
	case "TRX":
		return Location{"Trenton", "America/Chicago"}
	case "TRY":
		return Location{"Tororo", "Africa/Kampala"}
	case "TRZ":
		return Location{"Tiruchirappally", "Asia/Kolkata"}
	case "TSA":
		return Location{"Taipei City", "Asia/Taipei"}
	case "TSB":
		return Location{"Tsumeb", "Africa/Windhoek"}
	case "TSC":
		return Location{"Taisha", "America/Guayaquil"}
	case "TSF":
		return Location{"Treviso", "Europe/Rome"}
	case "TSH":
		return Location{"Tshikapa", "Africa/Lubumbashi"}
	case "TSJ":
		return Location{"Tsushima", "Asia/Tokyo"}
	case "TSL":
		return Location{"", "America/Mexico_City"}
	case "TSM":
		return Location{"Taos", "America/Denver"}
	case "TSN":
		return Location{"Tianjin", "Asia/Shanghai"}
	case "TSP":
		return Location{"Tehachapi", "America/Los_Angeles"}
	case "TSR":
		return Location{"Timisoara", "Europe/Bucharest"}
	case "TST":
		return Location{"", "Asia/Bangkok"}
	case "TSU":
		return Location{"Tabiteuea South", "Pacific/Tarawa"}
	case "TSV":
		return Location{"Townsville", "Australia/Brisbane"}
	case "TSX":
		return Location{"Santan-Borneo Island", "Asia/Makassar"}
	case "TSY":
		return Location{"Tasikmalaya-Java Island", "Asia/Jakarta"}
	case "TTA":
		return Location{"Tan Tan", "Africa/Casablanca"}
	case "TTB":
		return Location{"Arbatax", "Europe/Rome"}
	case "TTC":
		return Location{"Taltal", "America/Santiago"}
	case "TTD":
		return Location{"Portland", "America/Los_Angeles"}
	case "TTE":
		return Location{"Sango-Ternate Island", "Asia/Jayapura"}
	case "TTG":
		return Location{"Tartagal", "America/Argentina/Salta"}
	case "TTH":
		return Location{"Thumrait", "Asia/Muscat"}
	case "TTI":
		return Location{"Tetiaroa", "Pacific/Tahiti"}
	case "TTJ":
		return Location{"Tottori", "Asia/Tokyo"}
	case "TTN":
		return Location{"Trenton", "America/New_York"}
	case "TTO":
		return Location{"Britton", "America/Chicago"}
	case "TTQ":
		return Location{"Roxana", "America/Costa_Rica"}
	case "TTR":
		return Location{"Tanah Toraja-Celebes Island", "Asia/Makassar"}
	case "TTS":
		return Location{"Tsaratanana", "Indian/Antananarivo"}
	case "TTT":
		return Location{"Taitung City", "Asia/Taipei"}
	case "TTU":
		return Location{"", "Africa/Casablanca"}
	case "TTX":
		return Location{"", "Australia/Perth"}
	case "TUA":
		return Location{"Tulcan", "America/Guayaquil"}
	case "TUB":
		return Location{"", "Pacific/Tahiti"}
	case "TUC":
		return Location{"San Miguel de Tucuman", "America/Argentina/Tucuman"}
	case "TUD":
		return Location{"Tambacounda", "Africa/Dakar"}
	case "TUF":
		return Location{"Tours/Val de Loire (Loire Valley)", "Europe/Paris"}
	case "TUG":
		return Location{"Tuguegarao City", "Asia/Manila"}
	case "TUI":
		return Location{"", "Asia/Riyadh"}
	case "TUJ":
		return Location{"Maji", "Africa/Addis_Ababa"}
	case "TUK":
		return Location{"Turbat", "Asia/Karachi"}
	case "TUL":
		return Location{"Tulsa", "America/Chicago"}
	case "TUM":
		return Location{"", "Australia/Sydney"}
	case "TUN":
		return Location{"Tunis", "Africa/Tunis"}
	case "TUO":
		return Location{"Taupo", "Pacific/Auckland"}
	case "TUP":
		return Location{"Tupelo", "America/Chicago"}
	case "TUQ":
		return Location{"Tougan", "Africa/Ouagadougou"}
	case "TUR":
		return Location{"Tucurui", "America/Belem"}
	case "TUS":
		return Location{"Tucson", "America/Phoenix"}
	case "TUU":
		return Location{"", "Asia/Riyadh"}
	case "TUV":
		return Location{"Tucupita", "America/Caracas"}
	case "TVA":
		return Location{"Morafenobe", "Indian/Antananarivo"}
	case "TVC":
		return Location{"Traverse City", "America/Detroit"}
	case "TVF":
		return Location{"Thief River Falls", "America/Chicago"}
	case "TVI":
		return Location{"Thomasville", "America/New_York"}
	case "TVL":
		return Location{"South Lake Tahoe", "America/Los_Angeles"}
	case "TVU":
		return Location{"Matei", "Pacific/Fiji"}
	case "TVY":
		return Location{"Dawei", "Asia/Yangon"}
	case "TWB":
		return Location{"", "Australia/Brisbane"}
	case "TWC":
		return Location{"Tumxuk", "Asia/Shanghai"}
	case "TWE":
		return Location{"Taylor", "America/Nome"}
	case "TWF":
		return Location{"Twin Falls", "America/Boise"}
	case "TWU":
		return Location{"Tawau", "Asia/Kuching"}
	case "TWZ":
		return Location{"Twitzel", "Pacific/Auckland"}
	case "TXA":
		return Location{"", "America/Mexico_City"}
	case "TXF":
		return Location{"Teixeira De Freitas", "America/Bahia"}
	case "TXG":
		return Location{"Taichung City", "Asia/Taipei"}
	case "TXK":
		return Location{"Texarkana", "America/Chicago"}
	case "TXL":
		return Location{"Berlin", "Europe/Berlin"}
	case "TXM":
		return Location{"Atinjoe-Papua Island", "Asia/Jayapura"}
	case "TXN":
		return Location{"Huangshan", "Asia/Shanghai"}
	case "TXU":
		return Location{"Tabou", "Africa/Abidjan"}
	case "TYB":
		return Location{"", "Australia/Sydney"}
	case "TYD":
		return Location{"Tynda", "Asia/Yakutsk"}
	case "TYF":
		return Location{"", "Europe/Stockholm"}
	case "TYG":
		return Location{"", "Australia/Brisbane"}
	case "TYL":
		return Location{"", "America/Lima"}
	case "TYM":
		return Location{"", "America/Nassau"}
	case "TYN":
		return Location{"Taiyuan", "Asia/Shanghai"}
	case "TYP":
		return Location{"Tobermorey", "Australia/Darwin"}
	case "TYR":
		return Location{"Tyler", "America/Chicago"}
	case "TYS":
		return Location{"Knoxville", "America/New_York"}
	case "TYT":
		return Location{"Treinta y Tres", "America/Montevideo"}
	case "TYZ":
		return Location{"Taylor", "America/Phoenix"}
	case "TZC":
		return Location{"Caro", "America/Detroit"}
	case "TZL":
		return Location{"Tuzla", "Europe/Sarajevo"}
	case "TZR":
		return Location{"Taszar", "Europe/Budapest"}
	case "TZX":
		return Location{"Trabzon", "Europe/Istanbul"}
	case "UAB":
		return Location{"Adana", "Europe/Istanbul"}
	case "UAH":
		return Location{"Ua Huka", "Pacific/Marquesas"}
	case "UAI":
		return Location{"Suai", "Asia/Dili"}
	case "UAK":
		return Location{"Narsarsuaq", "America/Nuuk"}
	case "UAL":
		return Location{"Luau", "Africa/Luanda"}
	case "UAM":
		return Location{"Andersen", "Pacific/Guam"}
	case "UAP":
		return Location{"Ua Pou", "Pacific/Marquesas"}
	case "UAQ":
		return Location{"San Juan", "America/Argentina/San_Juan"}
	case "UAR":
		return Location{"Bouarfa", "Africa/Casablanca"}
	case "UAS":
		return Location{"Samburu South", "Africa/Nairobi"}
	case "UAX":
		return Location{"Uaxactun", "America/Guatemala"}
	case "UBA":
		return Location{"Uberaba", "America/Sao_Paulo"}
	case "UBB":
		return Location{"Mabuiag Island", "Australia/Brisbane"}
	case "UBJ":
		return Location{"Ube", "Asia/Tokyo"}
	case "UBN":
		return Location{"Ulaanbaatar", "Asia/Ulaanbaatar"}
	case "UBP":
		return Location{"Ubon Ratchathani", "Asia/Bangkok"}
	case "UBR":
		return Location{"Ubrub-Papua Island", "Asia/Jayapura"}
	case "UBS":
		return Location{"Columbus", "America/Chicago"}
	case "UBT":
		return Location{"Ubatuba", "America/Sao_Paulo"}
	case "UBU":
		return Location{"", "Australia/Perth"}
	case "UCB":
		return Location{"Ulanqab", "Asia/Shanghai"}
	case "UCK":
		return Location{"Lutsk", "Europe/Kyiv"}
	case "UCN":
		return Location{"Buchanan", "Africa/Monrovia"}
	case "UCT":
		return Location{"Ukhta", "Europe/Moscow"}
	case "UCY":
		return Location{"Union City", "America/Chicago"}
	case "UCZ":
		return Location{"Uchiza", "America/Lima"}
	case "UDA":
		return Location{"", "Australia/Brisbane"}
	case "UDD":
		return Location{"Palm Springs", "America/Los_Angeles"}
	case "UDE":
		return Location{"Uden", "Europe/Amsterdam"}
	case "UDI":
		return Location{"Uberlandia", "America/Sao_Paulo"}
	case "UDJ":
		return Location{"Uzhhorod", "Europe/Bratislava"}
	case "UDN":
		return Location{"Udine", "Europe/Rome"}
	case "UDR":
		return Location{"Udaipur", "Asia/Kolkata"}
	case "UEE":
		return Location{"", "Australia/Hobart"}
	case "UEL":
		return Location{"Quelimane", "Africa/Maputo"}
	case "UEN":
		return Location{"Urengoy", "Asia/Yekaterinburg"}
	case "UEO":
		return Location{"", "Asia/Tokyo"}
	case "UES":
		return Location{"Waukesha", "America/Chicago"}
	case "UET":
		return Location{"Quetta", "Asia/Karachi"}
	case "UFA":
		return Location{"Ufa", "Asia/Yekaterinburg"}
	case "UGA":
		return Location{"Bulgan", "Asia/Ulaanbaatar"}
	case "UGC":
		return Location{"Urgench", "Asia/Samarkand"}
	case "UGN":
		return Location{"Chicago/Waukegan", "America/Chicago"}
	case "UGO":
		return Location{"Uige", "Africa/Luanda"}
	case "UGT":
		return Location{"Umnugobitour", "Asia/Ulaanbaatar"}
	case "UHE":
		return Location{"Uherske Hradiste", "Europe/Prague"}
	case "UIB":
		return Location{"Quibdo", "America/Bogota"}
	case "UIH":
		return Location{"Quy Nohn", "Asia/Ho_Chi_Minh"}
	case "UII":
		return Location{"Utila Island", "America/Tegucigalpa"}
	case "UIK":
		return Location{"Ust-Ilimsk", "Asia/Irkutsk"}
	case "UIL":
		return Location{"Quillayute", "America/Los_Angeles"}
	case "UIN":
		return Location{"Quincy", "America/Chicago"}
	case "UIO":
		return Location{"Quito", "America/Guayaquil"}
	case "UIP":
		return Location{"Quimper/Pluguffan", "Europe/Paris"}
	case "UIQ":
		return Location{"Quion Hill", "Pacific/Efate"}
	case "UIR":
		return Location{"", "Australia/Sydney"}
	case "UJE":
		return Location{"Ujae Atoll", "Pacific/Majuro"}
	case "UKA":
		return Location{"Ukunda", "Africa/Nairobi"}
	case "UKB":
		return Location{"Kobe", "Asia/Tokyo"}
	case "UKG":
		return Location{"Ust-Kuyga", "Asia/Vladivostok"}
	case "UKI":
		return Location{"Ukiah", "America/Los_Angeles"}
	case "UKK":
		return Location{"Ust Kamenogorsk", "Asia/Almaty"}
	case "UKS":
		return Location{"Sevastopol", "Europe/Simferopol"}
	case "UKT":
		return Location{"Quakertown", "America/New_York"}
	case "UKU":
		return Location{"Nuku", "Pacific/Port_Moresby"}
	case "UKX":
		return Location{"Ust-Kut", "Asia/Irkutsk"}
	case "ULA":
		return Location{"San Julian", "America/Argentina/Rio_Gallegos"}
	case "ULB":
		return Location{"Ambryn Island", "Pacific/Efate"}
	case "ULD":
		return Location{"Ulundi", "Africa/Johannesburg"}
	case "ULG":
		return Location{"", "Asia/Hovd"}
	case "ULH":
		return Location{"Al-'Ula", "Asia/Riyadh"}
	case "ULK":
		return Location{"Lensk", "Asia/Yakutsk"}
	case "ULM":
		return Location{"New Ulm", "America/Chicago"}
	case "ULN":
		return Location{"Ulan Bator", "Asia/Ulaanbaatar"}
	case "ULO":
		return Location{"", "Asia/Hovd"}
	case "ULP":
		return Location{"", "Australia/Brisbane"}
	case "ULQ":
		return Location{"Tulua", "America/Bogota"}
	case "ULU":
		return Location{"Gulu", "Africa/Kampala"}
	case "ULV":
		return Location{"Ulyanovsk", "Europe/Ulyanovsk"}
	case "ULX":
		return Location{"Ulusaba", "Africa/Johannesburg"}
	case "ULY":
		return Location{"Ulyanovsk", "Europe/Ulyanovsk"}
	case "UMA":
		return Location{"Maisi", "America/Havana"}
	case "UME":
		return Location{"Umea", "Europe/Stockholm"}
	case "UMI":
		return Location{"Quince Mil", "America/Lima"}
	case "UMM":
		return Location{"Summit", "America/Anchorage"}
	case "UMR":
		return Location{"Woomera", "Australia/Adelaide"}
	case "UMS":
		return Location{"Ust-Maya", "Asia/Khandyga"}
	case "UMT":
		return Location{"Umiat", "America/Anchorage"}
	case "UMU":
		return Location{"Umuarama", "America/Sao_Paulo"}
	case "UMY":
		return Location{"Sumy", "Europe/Kyiv"}
	case "UNA":
		return Location{"Una", "America/Bahia"}
	case "UND":
		return Location{"", "Asia/Kabul"}
	case "UNE":
		return Location{"Qacha's Nek", "Africa/Johannesburg"}
	case "UNG":
		return Location{"Kiunga", "Pacific/Port_Moresby"}
	case "UNI":
		return Location{"Union Island", "America/St_Vincent"}
	case "UNK":
		return Location{"Unalakleet", "America/Anchorage"}
	case "UNN":
		return Location{"", "Asia/Bangkok"}
	case "UNT":
		return Location{"Shetland Islands", "Europe/London"}
	case "UNU":
		return Location{"Juneau", "America/Chicago"}
	case "UOA":
		return Location{"Mururoa Atoll", "Pacific/Tahiti"}
	case "UOL":
		return Location{"Buol-Celebes Island", "Asia/Makassar"}
	case "UOS":
		return Location{"Sewanee", "America/Chicago"}
	case "UOX":
		return Location{"Oxford", "America/Chicago"}
	case "UPB":
		return Location{"Havana", "America/Havana"}
	case "UPG":
		return Location{"Ujung Pandang-Celebes Island", "Asia/Makassar"}
	case "UPL":
		return Location{"Upala", "America/Costa_Rica"}
	case "UPN":
		return Location{"", "America/Mexico_City"}
	case "UPP":
		return Location{"Hawi", "Pacific/Honolulu"}
	case "UPV":
		return Location{"Upavon", "Europe/London"}
	case "URA":
		return Location{"Uralsk", "Asia/Oral"}
	case "URC":
		return Location{"Urumqi", "Asia/Shanghai"}
	case "URD":
		return Location{"Ebermannstadt", "Europe/Berlin"}
	case "URE":
		return Location{"Kuressaare", "Europe/Tallinn"}
	case "URG":
		return Location{"Uruguaiana", "America/Sao_Paulo"}
	case "URJ":
		return Location{"Uray", "Asia/Yekaterinburg"}
	case "URM":
		return Location{"", "America/Caracas"}
	case "URO":
		return Location{"Rouen/Vallee de Seine", "Europe/Paris"}
	case "URR":
		return Location{"Urrao", "America/Bogota"}
	case "URS":
		return Location{"Kursk", "Europe/Moscow"}
	case "URT":
		return Location{"Surat Thani", "Asia/Bangkok"}
	case "URY":
		return Location{"", "Asia/Riyadh"}
	case "USA":
		return Location{"Concord", "America/New_York"}
	case "USH":
		return Location{"Ushuahia", "America/Argentina/Ushuaia"}
	case "USI":
		return Location{"Mabaruma", "America/Guyana"}
	case "USJ":
		return Location{"Usharal", "Asia/Almaty"}
	case "USK":
		return Location{"Usinsk", "Europe/Moscow"}
	case "USL":
		return Location{"", "Australia/Perth"}
	case "USM":
		return Location{"Na Thon (Ko Samui Island)", "Asia/Bangkok"}
	case "USN":
		return Location{"Ulsan", "Asia/Seoul"}
	case "USQ":
		return Location{"Usak", "Europe/Istanbul"}
	case "USR":
		return Location{"Ust-Nera", "Asia/Ust-Nera"}
	case "USS":
		return Location{"Sancti Spiritus", "America/Havana"}
	case "UST":
		return Location{"St Augustine", "America/New_York"}
	case "USU":
		return Location{"Coron", "Asia/Manila"}
	case "UTA":
		return Location{"", "Africa/Harare"}
	case "UTB":
		return Location{"", "Australia/Brisbane"}
	case "UTE":
		return Location{"Bultfontein", "Africa/Johannesburg"}
	case "UTG":
		return Location{"Quthing", "Africa/Maseru"}
	case "UTH":
		return Location{"Udon Thani", "Asia/Bangkok"}
	case "UTI":
		return Location{"Utti / Valkeala", "Europe/Helsinki"}
	case "UTM":
		return Location{"Tunica", "America/Chicago"}
	case "UTN":
		return Location{"Upington", "Africa/Johannesburg"}
	case "UTO":
		return Location{"Utopia Creek", "America/Anchorage"}
	case "UTP":
		return Location{"Rayong", "Asia/Bangkok"}
	case "UTR":
		return Location{"Uttaradit", "Asia/Bangkok"}
	case "UTS":
		return Location{"Ust-Tsylma", "Europe/Moscow"}
	case "UTT":
		return Location{"Mthatha", "Africa/Johannesburg"}
	case "UTW":
		return Location{"Queenstown", "Africa/Johannesburg"}
	case "UUA":
		return Location{"Bugulma", "Europe/Moscow"}
	case "UUD":
		return Location{"Ulan Ude", "Asia/Irkutsk"}
	case "UUK":
		return Location{"Kuparuk", "America/Anchorage"}
	case "UUN":
		return Location{"", "Asia/Choibalsan"}
	case "UUS":
		return Location{"Yuzhno-Sakhalinsk", "Asia/Sakhalin"}
	case "UVA":
		return Location{"Uvalde", "America/Chicago"}
	case "UVE":
		return Location{"Ouvea", "Pacific/Noumea"}
	case "UVF":
		return Location{"Vieux Fort", "America/St_Lucia"}
	case "UVL":
		return Location{"", "Africa/Cairo"}
	case "UYL":
		return Location{"Nyala", "Africa/Khartoum"}
	case "UYN":
		return Location{"Yulin", "Asia/Shanghai"}
	case "UYU":
		return Location{"Quijarro", "America/La_Paz"}
	case "UZC":
		return Location{"Uzice", "Europe/Belgrade"}
	case "UZU":
		return Location{"Curuzu Cuatia", "America/Argentina/Cordoba"}
	case "VAA":
		return Location{"Vaasa", "Europe/Helsinki"}
	case "VAC":
		return Location{"Cloppenburg", "Europe/Berlin"}
	case "VAD":
		return Location{"Valdosta", "America/New_York"}
	case "VAF":
		return Location{"Valence/Chabeuil", "Europe/Paris"}
	case "VAG":
		return Location{"Varginha", "America/Sao_Paulo"}
	case "VAH":
		return Location{"Vallegrande", "America/La_Paz"}
	case "VAI":
		return Location{"", "Pacific/Port_Moresby"}
	case "VAK":
		return Location{"Chevak", "America/Nome"}
	case "VAL":
		return Location{"Valenca", "America/Bahia"}
	case "VAM":
		return Location{"Maamigili", "Indian/Maldives"}
	case "VAN":
		return Location{"Van", "Europe/Istanbul"}
	case "VAO":
		return Location{"Suavanao", "Pacific/Guadalcanal"}
	case "VAP":
		return Location{"Vina Del Mar", "America/Santiago"}
	case "VAR":
		return Location{"Varna", "Europe/Sofia"}
	case "VAS":
		return Location{"Sivas", "Europe/Istanbul"}
	case "VAT":
		return Location{"Vatomandry", "Indian/Antananarivo"}
	case "VAU":
		return Location{"Vatukoula", "Pacific/Fiji"}
	case "VAV":
		return Location{"Vava'u Island", "Pacific/Tongatapu"}
	case "VAW":
		return Location{"Vardo", "Europe/Oslo"}
	case "VBA":
		return Location{"Aeng", "Asia/Yangon"}
	case "VBC":
		return Location{"Mandalay", "Asia/Yangon"}
	case "VBG":
		return Location{"Lompoc", "America/Los_Angeles"}
	case "VBP":
		return Location{"Bokpyinn", "Asia/Yangon"}
	case "VBS":
		return Location{"Brescia", "Europe/Rome"}
	case "VBV":
		return Location{"Vanua Balavu", "Pacific/Fiji"}
	case "VBY":
		return Location{"Visby", "Europe/Stockholm"}
	case "VCA":
		return Location{"Can Tho", "Asia/Ho_Chi_Minh"}
	case "VCD":
		return Location{"", "Australia/Darwin"}
	case "VCE":
		return Location{"Venezia", "Europe/Rome"}
	case "VCH":
		return Location{"Vichadero", "America/Montevideo"}
	case "VCL":
		return Location{"Dung Quat Bay", "Asia/Ho_Chi_Minh"}
	case "VCP":
		return Location{"Campinas", "America/Sao_Paulo"}
	case "VCR":
		return Location{"Carora", "America/Caracas"}
	case "VCS":
		return Location{"Con Ong", "Asia/Ho_Chi_Minh"}
	case "VCT":
		return Location{"Victoria", "America/Chicago"}
	case "VCV":
		return Location{"Victorville", "America/Los_Angeles"}
	case "VDA":
		return Location{"Eilat", "Asia/Jerusalem"}
	case "VDB":
		return Location{"", "Europe/Oslo"}
	case "VDC":
		return Location{"Vitória da Conquista", "America/Bahia"}
	case "VDE":
		return Location{"El Hierro Island", "Atlantic/Canary"}
	case "VDH":
		return Location{"Dong Hoi", "Asia/Bangkok"}
	case "VDI":
		return Location{"Vidalia", "America/New_York"}
	case "VDM":
		return Location{"Viedma / Carmen de Patagones", "America/Argentina/Salta"}
	case "VDO":
		return Location{"Vân Đồn", "Asia/Bangkok"}
	case "VDP":
		return Location{"", "America/Caracas"}
	case "VDR":
		return Location{"Villa Dolores", "America/Argentina/Cordoba"}
	case "VDS":
		return Location{"Vadso", "Europe/Oslo"}
	case "VDY":
		return Location{"", "Asia/Kolkata"}
	case "VDZ":
		return Location{"Valdez", "America/Anchorage"}
	case "VEE":
		return Location{"Venetie", "America/Anchorage"}
	case "VEL":
		return Location{"Vernal", "America/Denver"}
	case "VER":
		return Location{"Veracruz", "America/Mexico_City"}
	case "VEV":
		return Location{"Barakoma", "Pacific/Guadalcanal"}
	case "VEY":
		return Location{"Vestmannaeyjar", "Atlantic/Reykjavik"}
	case "VFA":
		return Location{"Victoria Falls", "Africa/Harare"}
	case "VGA":
		return Location{"", "Asia/Kolkata"}
	case "VGD":
		return Location{"Vologda", "Europe/Moscow"}
	case "VGO":
		return Location{"Vigo", "Europe/Madrid"}
	case "VGT":
		return Location{"Las Vegas", "America/Los_Angeles"}
	case "VGZ":
		return Location{"Villa Garzon", "America/Bogota"}
	case "VHC":
		return Location{"Saurimo", "Africa/Luanda"}
	case "VHM":
		return Location{"", "Europe/Stockholm"}
	case "VHN":
		return Location{"Van Horn", "America/Chicago"}
	case "VHV":
		return Location{"Verkhnevilyuisk", "Asia/Yakutsk"}
	case "VHY":
		return Location{"Vichy/Charmeil", "Europe/Paris"}
	case "VHZ":
		return Location{"Vahitahi", "Pacific/Tahiti"}
	case "VIA":
		return Location{"Videira", "America/Sao_Paulo"}
	case "VIC":
		return Location{"Vicenza", "Europe/Rome"}
	case "VID":
		return Location{"Vidin", "Europe/Sofia"}
	case "VIE":
		return Location{"Vienna", "Europe/Vienna"}
	case "VIG":
		return Location{"El Vigia", "America/Caracas"}
	case "VIH":
		return Location{"Rolla/Vichy", "America/Chicago"}
	case "VII":
		return Location{"Vinh", "Asia/Bangkok"}
	case "VIJ":
		return Location{"Spanish Town", "America/Tortola"}
	case "VIL":
		return Location{"Dakhla", "Africa/El_Aaiun"}
	case "VIN":
		return Location{"Vinnitsa", "Europe/Kyiv"}
	case "VIP":
		return Location{"Payerne", "Europe/Zurich"}
	case "VIQ":
		return Location{"Viqueque", "Asia/Dili"}
	case "VIR":
		return Location{"Durban", "Africa/Johannesburg"}
	case "VIS":
		return Location{"Visalia", "America/Los_Angeles"}
	case "VIT":
		return Location{"Alava", "Europe/Madrid"}
	case "VIX":
		return Location{"Vitoria", "America/Sao_Paulo"}
	case "VIY":
		return Location{"Villacoublay/Velizy", "Europe/Paris"}
	case "VJB":
		return Location{"Xai-Xai", "Africa/Maputo"}
	case "VJI":
		return Location{"Abingdon", "America/New_York"}
	case "VKG":
		return Location{"Rach Gia", "Asia/Ho_Chi_Minh"}
	case "VKO":
		return Location{"Moscow", "Europe/Moscow"}
	case "VKS":
		return Location{"Vicksburg", "America/Chicago"}
	case "VKT":
		return Location{"Vorkuta", "Europe/Moscow"}
	case "VKV":
		return Location{"Arkhangelsk", "Europe/Moscow"}
	case "VKX":
		return Location{"Friendly", "America/New_York"}
	case "VLA":
		return Location{"Vandalia", "America/Chicago"}
	case "VLC":
		return Location{"Valencia", "Europe/Madrid"}
	case "VLD":
		return Location{"Valdosta", "America/New_York"}
	case "VLG":
		return Location{"Villa Gesell", "America/Argentina/Buenos_Aires"}
	case "VLI":
		return Location{"Port Vila", "Pacific/Efate"}
	case "VLK":
		return Location{"", "Europe/Moscow"}
	case "VLL":
		return Location{"Valladolid", "Europe/Madrid"}
	case "VLM":
		return Location{"Villamontes", "America/La_Paz"}
	case "VLN":
		return Location{"Valencia", "America/Caracas"}
	case "VLP":
		return Location{"Vila Rica", "America/Cuiaba"}
	case "VLR":
		return Location{"Vallenar", "America/Santiago"}
	case "VLS":
		return Location{"Valesdir", "Pacific/Efate"}
	case "VLU":
		return Location{"Velikiye Luki", "Europe/Moscow"}
	case "VLV":
		return Location{"Valera", "America/Caracas"}
	case "VME":
		return Location{"Villa Reynolds", "America/Argentina/San_Luis"}
	case "VMU":
		return Location{"Baimuru", "Pacific/Port_Moresby"}
	case "VNA":
		return Location{"Saravane", "Asia/Vientiane"}
	case "VNC":
		return Location{"Venice", "America/New_York"}
	case "VND":
		return Location{"Vangaindrano", "Indian/Antananarivo"}
	case "VNE":
		return Location{"Vannes/Meucon", "Europe/Paris"}
	case "VNO":
		return Location{"Vilnius", "Europe/Vilnius"}
	case "VNR":
		return Location{"", "Australia/Brisbane"}
	case "VNS":
		return Location{"Varanasi", "Asia/Kolkata"}
	case "VNT":
		return Location{"Ventspils", "Europe/Riga"}
	case "VNX":
		return Location{"Vilanculo", "Africa/Maputo"}
	case "VNY":
		return Location{"Van Nuys", "America/Los_Angeles"}
	case "VOD":
		return Location{"Vodochoky", "Europe/Prague"}
	case "VOG":
		return Location{"Volgograd", "Europe/Volgograd"}
	case "VOH":
		return Location{"", "Indian/Antananarivo"}
	case "VOI":
		return Location{"Voinjama", "Africa/Monrovia"}
	case "VOK":
		return Location{"Camp Douglas", "America/Chicago"}
	case "VOL":
		return Location{"Nea Anchialos", "Europe/Athens"}
	case "VOT":
		return Location{"Votuporanga", "America/Sao_Paulo"}
	case "VOZ":
		return Location{"Voronezh", "Europe/Moscow"}
	case "VPE":
		return Location{"Ngiva", "Africa/Luanda"}
	case "VPN":
		return Location{"Vopnafjordur", "Atlantic/Reykjavik"}
	case "VPS":
		return Location{"Valparaiso", "America/Chicago"}
	case "VPY":
		return Location{"Chimoio", "Africa/Maputo"}
	case "VPZ":
		return Location{"Valparaiso", "America/Chicago"}
	case "VQQ":
		return Location{"Jacksonville", "America/New_York"}
	case "VQS":
		return Location{"Vieques Island", "America/Puerto_Rico"}
	case "VRA":
		return Location{"Varadero", "America/Havana"}
	case "VRB":
		return Location{"Vero Beach", "America/New_York"}
	case "VRC":
		return Location{"Virac", "Asia/Manila"}
	case "VRE":
		return Location{"Vredendal", "Africa/Johannesburg"}
	case "VRI":
		return Location{"", "Europe/Moscow"}
	case "VRK":
		return Location{"Varkaus / Joroinen", "Europe/Helsinki"}
	case "VRL":
		return Location{"", "Europe/Lisbon"}
	case "VRN":
		return Location{"Verona", "Europe/Rome"}
	case "VRO":
		return Location{"Matanzas", "America/Havana"}
	case "VRU":
		return Location{"Vyrburg", "Africa/Johannesburg"}
	case "VRZ":
		return Location{"Lavras", "America/Sao_Paulo"}
	case "VSA":
		return Location{"Villahermosa", "America/Mexico_City"}
	case "VSE":
		return Location{"", "Europe/Lisbon"}
	case "VSF":
		return Location{"Springfield", "America/New_York"}
	case "VSG":
		return Location{"Luhansk", "Europe/Kyiv"}
	case "VSK":
		return Location{"Kennewick", "America/Los_Angeles"}
	case "VST":
		return Location{"Stockholm / Vasteras", "Europe/Stockholm"}
	case "VSV":
		return Location{"Shravasti", "Asia/Kolkata"}
	case "VTB":
		return Location{"Vitebsk", "Europe/Minsk"}
	case "VTE":
		return Location{"Vientiane", "Asia/Vientiane"}
	case "VTF":
		return Location{"Vatulele", "Pacific/Fiji"}
	case "VTG":
		return Location{"Vung Tau", "Asia/Ho_Chi_Minh"}
	case "VTL":
		return Location{"Luxeuil", "Europe/Paris"}
	case "VTM":
		return Location{"Beersheba", "Asia/Jerusalem"}
	case "VTN":
		return Location{"Valentine", "America/Chicago"}
	case "VTU":
		return Location{"Las Tunas", "America/Havana"}
	case "VTZ":
		return Location{"Visakhapatnam", "Asia/Kolkata"}
	case "VUP":
		return Location{"Valledupar", "America/Bogota"}
	case "VUS":
		return Location{"Velikiy Ustyug", "Europe/Moscow"}
	case "VVB":
		return Location{"Mahanoro", "Indian/Antananarivo"}
	case "VVC":
		return Location{"Villavicencio", "America/Bogota"}
	case "VVI":
		return Location{"Santa Cruz", "America/La_Paz"}
	case "VVK":
		return Location{"Vastervik", "Europe/Stockholm"}
	case "VVO":
		return Location{"Vladivostok", "Asia/Vladivostok"}
	case "VVZ":
		return Location{"Illizi", "Africa/Algiers"}
	case "VXC":
		return Location{"Lichinga", "Africa/Maputo"}
	case "VXE":
		return Location{"Sao Pedro", "Atlantic/Cape_Verde"}
	case "VXO":
		return Location{"Vaxjo", "Europe/Stockholm"}
	case "VYD":
		return Location{"Vryheid", "Africa/Johannesburg"}
	case "VYI":
		return Location{"Vilyuisk", "Asia/Yakutsk"}
	case "VYS":
		return Location{"Peru", "America/Chicago"}
	case "WAA":
		return Location{"Wales", "America/Nome"}
	case "WAC":
		return Location{"Waca", "Africa/Addis_Ababa"}
	case "WAE":
		return Location{"", "Asia/Riyadh"}
	case "WAF":
		return Location{"Waana", "Asia/Karachi"}
	case "WAG":
		return Location{"Wanganui", "Pacific/Auckland"}
	case "WAH":
		return Location{"Wahpeton", "America/Chicago"}
	case "WAI":
		return Location{"Antsohihy", "Indian/Antananarivo"}
	case "WAK":
		return Location{"Ankazoabo", "Indian/Antananarivo"}
	case "WAL":
		return Location{"Wallops Island", "America/New_York"}
	case "WAM":
		return Location{"Ambatondrazaka", "Indian/Antananarivo"}
	case "WAO":
		return Location{"Wabo", "Pacific/Port_Moresby"}
	case "WAP":
		return Location{"Alto Palena", "America/Santiago"}
	case "WAQ":
		return Location{"Antsalova", "Indian/Antananarivo"}
	case "WAR":
		return Location{"Waris-Papua Island", "Asia/Jayapura"}
	case "WAT":
		return Location{"Waterford", "Europe/Dublin"}
	case "WAV":
		return Location{"", "Australia/Darwin"}
	case "WAW":
		return Location{"Warsaw", "Europe/Warsaw"}
	case "WAX":
		return Location{"Zuwara", "Africa/Tripoli"}
	case "WAY":
		return Location{"Waynesburg", "America/New_York"}
	case "WAZ":
		return Location{"", "Australia/Brisbane"}
	case "WBA":
		return Location{"Seram Island", "Asia/Jayapura"}
	case "WBD":
		return Location{"Befandriana", "Indian/Antananarivo"}
	case "WBG":
		return Location{"", "Europe/Berlin"}
	case "WBM":
		return Location{"", "Pacific/Port_Moresby"}
	case "WBO":
		return Location{"Beroroha", "Indian/Antananarivo"}
	case "WBQ":
		return Location{"Beaver", "America/Anchorage"}
	case "WBR":
		return Location{"Big Rapids", "America/Detroit"}
	case "WBU":
		return Location{"Boulder", "America/Denver"}
	case "WBW":
		return Location{"Wilkes-Barre", "America/New_York"}
	case "WCA":
		return Location{"Castro", "America/Santiago"}
	case "WCH":
		return Location{"Chaiten", "America/Santiago"}
	case "WCR":
		return Location{"Chandalar Lake", "America/Anchorage"}
	case "WDG":
		return Location{"Enid", "America/Chicago"}
	case "WDH":
		return Location{"Windhoek", "Africa/Windhoek"}
	case "WDI":
		return Location{"", "Australia/Brisbane"}
	case "WDR":
		return Location{"Winder", "America/New_York"}
	case "WEF":
		return Location{"Weifang", "Asia/Shanghai"}
	case "WEH":
		return Location{"Weihai", "Asia/Shanghai"}
	case "WEI":
		return Location{"Weipa", "Australia/Brisbane"}
	case "WEL":
		return Location{"Welkom", "Africa/Johannesburg"}
	case "WET":
		return Location{"Wagethe-Papua Island", "Asia/Jayapura"}
	case "WEW":
		return Location{"", "Australia/Sydney"}
	case "WEX":
		return Location{"Wexford", "Europe/Dublin"}
	case "WFI":
		return Location{"", "Indian/Antananarivo"}
	case "WFK":
		return Location{"Frenchville", "America/New_York"}
	case "WGA":
		return Location{"Wagga Wagga", "Australia/Sydney"}
	case "WGB":
		return Location{"Bahawalnagar", "Asia/Karachi"}
	case "WGC":
		return Location{"Warrangal", "Asia/Kolkata"}
	case "WGE":
		return Location{"", "Australia/Sydney"}
	case "WGO":
		return Location{"Winchester", "America/New_York"}
	case "WGP":
		return Location{"Waingapu-Sumba Island", "Asia/Makassar"}
	case "WGT":
		return Location{"", "Australia/Melbourne"}
	case "WHA":
		return Location{"Wuhu", "Asia/Shanghai"}
	case "WHF":
		return Location{"Wadi Halfa", "Africa/Khartoum"}
	case "WHK":
		return Location{"", "Pacific/Auckland"}
	case "WHO":
		return Location{"", "Pacific/Auckland"}
	case "WHP":
		return Location{"Los Angeles", "America/Los_Angeles"}
	case "WHS":
		return Location{"Whalsay", "Europe/London"}
	case "WHT":
		return Location{"Wharton", "America/Chicago"}
	case "WHU":
		return Location{"Wuhu", "Asia/Shanghai"}
	case "WIC":
		return Location{"Wick", "Europe/London"}
	case "WIE":
		return Location{"Wiesbaden", "Europe/Berlin"}
	case "WIK":
		return Location{"", "Pacific/Auckland"}
	case "WIL":
		return Location{"Nairobi", "Africa/Nairobi"}
	case "WIN":
		return Location{"", "Australia/Brisbane"}
	case "WIO":
		return Location{"", "Australia/Sydney"}
	case "WIR":
		return Location{"Wairoa", "Pacific/Auckland"}
	case "WIT":
		return Location{"", "Australia/Perth"}
	case "WIX":
		return Location{"", "America/Mexico_City"}
	case "WJF":
		return Location{"Lancaster", "America/Los_Angeles"}
	case "WJR":
		return Location{"Wajir", "Africa/Nairobi"}
	case "WJU":
		return Location{"Wonju", "Asia/Seoul"}
	case "WKA":
		return Location{"", "Pacific/Auckland"}
	case "WKB":
		return Location{"", "Australia/Melbourne"}
	case "WKF":
		return Location{"Pretoria", "Africa/Johannesburg"}
	case "WKI":
		return Location{"Hwange", "Africa/Harare"}
	case "WKJ":
		return Location{"Wakkanai", "Asia/Tokyo"}
	case "WKR":
		return Location{"", "America/Nassau"}
	case "WLA":
		return Location{"Wallal", "Australia/Perth"}
	case "WLC":
		return Location{"", "Australia/Sydney"}
	case "WLD":
		return Location{"Winfield/Arkansas City", "America/Chicago"}
	case "WLE":
		return Location{"", "Australia/Brisbane"}
	case "WLG":
		return Location{"Wellington", "Pacific/Auckland"}
	case "WLH":
		return Location{"Walaha", "Pacific/Efate"}
	case "WLK":
		return Location{"Selawik", "America/Anchorage"}
	case "WLL":
		return Location{"", "Australia/Darwin"}
	case "WLO":
		return Location{"", "Australia/Darwin"}
	case "WLP":
		return Location{"", "Australia/Perth"}
	case "WLS":
		return Location{"Wallis Island", "Pacific/Wallis"}
	case "WLW":
		return Location{"Willows", "America/Los_Angeles"}
	case "WMA":
		return Location{"Mandritsara", "Indian/Antananarivo"}
	case "WMB":
		return Location{"", "Australia/Melbourne"}
	case "WMC":
		return Location{"Winnemucca", "America/Los_Angeles"}
	case "WMD":
		return Location{"Mandabe", "Indian/Antananarivo"}
	case "WME":
		return Location{"", "Australia/Perth"}
	case "WMH":
		return Location{"Mountain Home", "America/Chicago"}
	case "WMI":
		return Location{"Warsaw", "Europe/Warsaw"}
	case "WML":
		return Location{"Malaimbandy", "Indian/Antananarivo"}
	case "WMN":
		return Location{"", "Indian/Antananarivo"}
	case "WMO":
		return Location{"White Mountain", "America/Nome"}
	case "WMP":
		return Location{"Mampikony", "Indian/Antananarivo"}
	case "WMR":
		return Location{"Mananara Nord", "Indian/Antananarivo"}
	case "WMT":
		return Location{"Zunyi", "Asia/Shanghai"}
	case "WMX":
		return Location{"Wamena-Papua Island", "Asia/Jayapura"}
	case "WNA":
		return Location{"Napakiak", "America/Anchorage"}
	case "WND":
		return Location{"", "Australia/Perth"}
	case "WNN":
		return Location{"Wunnumin Lake", "America/Winnipeg"}
	case "WNP":
		return Location{"Naga", "Asia/Manila"}
	case "WNR":
		return Location{"", "Australia/Brisbane"}
	case "WNS":
		return Location{"Nawabash", "Asia/Karachi"}
	case "WNZ":
		return Location{"Wenzhou", "Asia/Shanghai"}
	case "WOA":
		return Location{"Wonenara", "Pacific/Port_Moresby"}
	case "WOE":
		return Location{"Bergen Op Zoom", "Europe/Amsterdam"}
	case "WOL":
		return Location{"", "Australia/Sydney"}
	case "WON":
		return Location{"", "Australia/Brisbane"}
	case "WOT":
		return Location{"Wang-an", "Asia/Taipei"}
	case "WOW":
		return Location{"Willow", "America/Anchorage"}
	case "WPA":
		return Location{"Puerto Aysen", "America/Santiago"}
	case "WPB":
		return Location{"Port Berge", "Indian/Antananarivo"}
	case "WPC":
		return Location{"Pincher Creek", "America/Edmonton"}
	case "WPK":
		return Location{"", "Australia/Brisbane"}
	case "WPR":
		return Location{"Porvenir", "America/Punta_Arenas"}
	case "WPU":
		return Location{"Puerto Williams", "America/Argentina/Ushuaia"}
	case "WRB":
		return Location{"Warner Robins", "America/New_York"}
	case "WRE":
		return Location{"", "Pacific/Auckland"}
	case "WRG":
		return Location{"Wrangell", "America/Sitka"}
	case "WRI":
		return Location{"Wrightstown", "America/New_York"}
	case "WRL":
		return Location{"Worland", "America/Denver"}
	case "WRO":
		return Location{"Wroclaw", "Europe/Warsaw"}
	case "WRT":
		return Location{"Warton", "Europe/London"}
	case "WRW":
		return Location{"", "Australia/Perth"}
	case "WRY":
		return Location{"Westray", "Europe/London"}
	case "WRZ":
		return Location{"Weerawila", "Asia/Colombo"}
	case "WSD":
		return Location{"White Sands", "America/Denver"}
	case "WSF":
		return Location{"Cape Sarichef", "America/Nome"}
	case "WSG":
		return Location{"Washington", "America/New_York"}
	case "WSH":
		return Location{"Shirley", "America/New_York"}
	case "WSK":
		return Location{"Wushan", "Asia/Shanghai"}
	case "WSN":
		return Location{"South Naknek", "America/Anchorage"}
	case "WSO":
		return Location{"Washabo", "America/Guyana"}
	case "WSP":
		return Location{"Waspam", "America/Managua"}
	case "WSR":
		return Location{"Wasior-Papua Island", "Asia/Jayapura"}
	case "WST":
		return Location{"Westerly", "America/New_York"}
	case "WSU":
		return Location{"Wasu", "Pacific/Port_Moresby"}
	case "WSY":
		return Location{"", "Australia/Brisbane"}
	case "WSZ":
		return Location{"", "Pacific/Auckland"}
	case "WTA":
		return Location{"Tambohorano", "Indian/Antananarivo"}
	case "WTB":
		return Location{"Wellcamp", "Australia/Brisbane"}
	case "WTD":
		return Location{"West End", "America/Nassau"}
	case "WTK":
		return Location{"Noatak", "America/Nome"}
	case "WTN":
		return Location{"Waddington", "Europe/London"}
	case "WTP":
		return Location{"Fatima Mission", "Pacific/Port_Moresby"}
	case "WTS":
		return Location{"Tsiroanomandidy", "Indian/Antananarivo"}
	case "WTZ":
		return Location{"", "Pacific/Auckland"}
	case "WUA":
		return Location{"Wuhai", "Asia/Shanghai"}
	case "WUD":
		return Location{"", "Australia/Adelaide"}
	case "WUG":
		return Location{"Wau", "Pacific/Port_Moresby"}
	case "WUH":
		return Location{"Wuhan", "Asia/Shanghai"}
	case "WUI":
		return Location{"", "Australia/Perth"}
	case "WUN":
		return Location{"", "Australia/Perth"}
	case "WUS":
		return Location{"Wuyishan", "Asia/Shanghai"}
	case "WUU":
		return Location{"Wau", "Africa/Juba"}
	case "WUX":
		return Location{"Wuxi", "Asia/Shanghai"}
	case "WUZ":
		return Location{"Wuzhou", "Asia/Shanghai"}
	case "WVB":
		return Location{"Walvis Bay", "Africa/Windhoek"}
	case "WVI":
		return Location{"Watsonville", "America/Los_Angeles"}
	case "WVK":
		return Location{"", "Indian/Antananarivo"}
	case "WVL":
		return Location{"Waterville", "America/New_York"}
	case "WVN":
		return Location{"Wilhelmshaven", "Europe/Berlin"}
	case "WWD":
		return Location{"Wildwood", "America/New_York"}
	case "WWI":
		return Location{"Woodie Woodie", "Australia/Perth"}
	case "WWK":
		return Location{"Wewak", "Pacific/Port_Moresby"}
	case "WWR":
		return Location{"Woodward", "America/Chicago"}
	case "WWT":
		return Location{"Newtok", "America/Nome"}
	case "WWY":
		return Location{"West Wyalong", "Australia/Sydney"}
	case "WXF":
		return Location{"Wethersfield", "Europe/London"}
	case "WXN":
		return Location{"Wanxian", "Asia/Shanghai"}
	case "WYA":
		return Location{"Whyalla", "Australia/Adelaide"}
	case "WYE":
		return Location{"Yengema", "Africa/Freetown"}
	case "WYN":
		return Location{"", "Australia/Perth"}
	case "WYS":
		return Location{"West Yellowstone", "America/Denver"}
	case "XAB":
		return Location{"Abbeville (Buigny/Saint-Maclou)", "Europe/Paris"}
	case "XAC":
		return Location{"Arcachon/La Teste-de-Buch", "Europe/Paris"}
	case "XAI":
		return Location{"Xinyang", "Asia/Shanghai"}
	case "XAN":
		return Location{"Evreux", "Europe/Paris"}
	case "XAP":
		return Location{"Chapeco", "America/Sao_Paulo"}
	case "XAR":
		return Location{"Aribinda", "Africa/Ouagadougou"}
	case "XAS":
		return Location{"Ales/Deaux", "Europe/Paris"}
	case "XAU":
		return Location{"Saul", "America/Cayenne"}
	case "XAV":
		return Location{"Angers", "Europe/Paris"}
	case "XBE":
		return Location{"Bearskin Lake", "America/Winnipeg"}
	case "XBG":
		return Location{"Bogande", "Africa/Ouagadougou"}
	case "XBJ":
		return Location{"Birjand", "Asia/Tehran"}
	case "XBK":
		return Location{"Bourg/Ceyzeriat", "Europe/Paris"}
	case "XBO":
		return Location{"Boulsa", "Africa/Ouagadougou"}
	case "XBQ":
		return Location{"Blois/Le Breuil", "Europe/Paris"}
	case "XBV":
		return Location{"Beaune/Challanges", "Europe/Paris"}
	case "XCB":
		return Location{"Cambrai/Epinoy", "Europe/Paris"}
	case "XCD":
		return Location{"Chalon/Champforgueil", "Europe/Paris"}
	case "XCH":
		return Location{"Christmas Island", "Indian/Christmas"}
	case "XCM":
		return Location{"Chatham-Kent", "America/Toronto"}
	case "XCO":
		return Location{"", "Australia/Melbourne"}
	case "XCP":
		return Location{"Compiegne", "Europe/Paris"}
	case "XCR":
		return Location{"Chalons/Vatry", "Europe/Paris"}
	case "XCX":
		return Location{"Biarritz", "Europe/Paris"}
	case "XCZ":
		return Location{"Charleville-Mezieres", "Europe/Paris"}
	case "XDA":
		return Location{"Perigueux", "Europe/Paris"}
	case "XDE":
		return Location{"Diebougou", "Africa/Ouagadougou"}
	case "XDJ":
		return Location{"Djibo", "Africa/Ouagadougou"}
	case "XDK":
		return Location{"Les Moeres", "Europe/Brussels"}
	case "XEN":
		return Location{"", "Asia/Shanghai"}
	case "XFN":
		return Location{"Xiangfan", "Asia/Shanghai"}
	case "XFW":
		return Location{"Hamburg", "Europe/Berlin"}
	case "XGA":
		return Location{"Gaoua", "Africa/Ouagadougou"}
	case "XGG":
		return Location{"Gorom-Gorom", "Africa/Ouagadougou"}
	case "XGN":
		return Location{"Xangongo", "Africa/Luanda"}
	case "XGR":
		return Location{"Kangiqsualujjuaq", "America/Toronto"}
	case "XGT":
		return Location{"Cahors", "Europe/Paris"}
	case "XIC":
		return Location{"Xichang", "Asia/Shanghai"}
	case "XIJ":
		return Location{"Ahmed Al Jaber AB", "Asia/Kuwait"}
	case "XIL":
		return Location{"Xilinhot", "Asia/Shanghai"}
	case "XIN":
		return Location{"Xingning", "Asia/Shanghai"}
	case "XIY":
		return Location{"Xianyang", "Asia/Shanghai"}
	case "XJM":
		return Location{"Mangla", "Asia/Karachi"}
	case "XKA":
		return Location{"Kantchari", "Africa/Ouagadougou"}
	case "XKH":
		return Location{"Xieng Khouang", "Asia/Vientiane"}
	case "XKS":
		return Location{"Kasabonika", "America/Winnipeg"}
	case "XKY":
		return Location{"Kaya", "Africa/Ouagadougou"}
	case "XLN":
		return Location{"Laon-Chambry", "Europe/Paris"}
	case "XLR":
		return Location{"Libourne/Artigues-de-Lussac", "Europe/Paris"}
	case "XLS":
		return Location{"Saint Louis", "Africa/Dakar"}
	case "XLU":
		return Location{"Leo", "Africa/Ouagadougou"}
	case "XLW":
		return Location{"Lemwerder", "Europe/Berlin"}
	case "XMC":
		return Location{"", "Australia/Melbourne"}
	case "XMD":
		return Location{"Madison", "America/Chicago"}
	case "XME":
		return Location{"Maubeuge/Elesmes", "Europe/Paris"}
	case "XMF":
		return Location{"Montbeliard/Courcelles", "Europe/Paris"}
	case "XMG":
		return Location{"Mahendranagar", "Asia/Kathmandu"}
	case "XMH":
		return Location{"", "Pacific/Tahiti"}
	case "XMI":
		return Location{"Masasi", "Africa/Dar_es_Salaam"}
	case "XMJ":
		return Location{"Mont-de-Marsan", "Europe/Paris"}
	case "XMK":
		return Location{"Annecy", "Europe/Paris"}
	case "XML":
		return Location{"", "Australia/Adelaide"}
	case "XMN":
		return Location{"Xiamen", "Asia/Shanghai"}
	case "XMS":
		return Location{"Macas", "America/Guayaquil"}
	case "XMU":
		return Location{"Moulins/Montbeugny", "Europe/Paris"}
	case "XMW":
		return Location{"Montauban", "Europe/Paris"}
	case "XMY":
		return Location{"Yam Island", "Australia/Brisbane"}
	case "XNA":
		return Location{"Fayetteville/Springdale/", "America/Chicago"}
	case "XNN":
		return Location{"Xining", "Asia/Shanghai"}
	case "XNU":
		return Location{"Nouna", "Africa/Ouagadougou"}
	case "XOG":
		return Location{"Orange/Caritat", "Europe/Paris"}
	case "XPA":
		return Location{"Pama", "Africa/Ouagadougou"}
	case "XPL":
		return Location{"Comayagua", "America/Tegucigalpa"}
	case "XPR":
		return Location{"Pine Ridge", "America/Denver"}
	case "XQP":
		return Location{"Quepos", "America/Costa_Rica"}
	case "XRH":
		return Location{"Richmond", "Australia/Sydney"}
	case "XRY":
		return Location{"Jerez de la Forntera", "Europe/Madrid"}
	case "XSB":
		return Location{"Sir Bani Yas", "Asia/Dubai"}
	case "XSC":
		return Location{"", "America/Grand_Turk"}
	case "XSD":
		return Location{"Tonopah", "America/Los_Angeles"}
	case "XSE":
		return Location{"Sebba", "Africa/Ouagadougou"}
	case "XSI":
		return Location{"South Indian Lake", "America/Winnipeg"}
	case "XSJ":
		return Location{"Peronne/Saint-Quentin", "Europe/Paris"}
	case "XSL":
		return Location{"Rochefort", "Europe/Paris"}
	case "XSP":
		return Location{"Seletar", "Asia/Kuala_Lumpur"}
	case "XSU":
		return Location{"Saumur/Saint-Florent", "Europe/Paris"}
	case "XTB":
		return Location{"Rochefort", "Europe/Paris"}
	case "XTG":
		return Location{"", "Australia/Brisbane"}
	case "XTL":
		return Location{"Tadoule Lake", "America/Winnipeg"}
	case "XTO":
		return Location{"", "Australia/Brisbane"}
	case "XTR":
		return Location{"", "Australia/Brisbane"}
	case "XUZ":
		return Location{"Xuzhou", "Asia/Shanghai"}
	case "XVF":
		return Location{"Villefranche/Tarare", "Europe/Paris"}
	case "XVN":
		return Location{"Verdun/Le Rozelier", "Europe/Paris"}
	case "XVO":
		return Location{"Vesoul/Frotey", "Europe/Paris"}
	case "XVS":
		return Location{"Valenciennes/Denain", "Europe/Paris"}
	case "XWA":
		return Location{"Williston", "America/Chicago"}
	case "XWP":
		return Location{"Hassleholm", "Europe/Stockholm"}
	case "XXN":
		return Location{"Riyadh", "Asia/Riyadh"}
	case "XYA":
		return Location{"Yandina", "Pacific/Guadalcanal"}
	case "XYE":
		return Location{"Ye", "Asia/Yangon"}
	case "XYR":
		return Location{"Yellow River Mission", "Pacific/Port_Moresby"}
	case "XZA":
		return Location{"Zabre", "Africa/Ouagadougou"}
	case "YAA":
		return Location{"Anahim Lake", "America/Vancouver"}
	case "YAB":
		return Location{"", "America/Rankin_Inlet"}
	case "YAC":
		return Location{"Cat Lake", "America/Winnipeg"}
	case "YAG":
		return Location{"Fort Frances", "America/Winnipeg"}
	case "YAH":
		return Location{"La Grande-4", "America/Toronto"}
	case "YAI":
		return Location{"Chillan", "America/Santiago"}
	case "YAK":
		return Location{"Yakutat", "America/Yakutat"}
	case "YAL":
		return Location{"Alert Bay", "America/Vancouver"}
	case "YAM":
		return Location{"Sault Ste Marie", "America/Detroit"}
	case "YAN":
		return Location{"Yangambi", "Africa/Lubumbashi"}
	case "YAO":
		return Location{"Yaounde", "Africa/Douala"}
	case "YAP":
		return Location{"Yap Island", "Pacific/Chuuk"}
	case "YAR":
		return Location{"La Grande-3", "America/Toronto"}
	case "YAS":
		return Location{"Yasawa Island", "Pacific/Fiji"}
	case "YAT":
		return Location{"Attawapiskat", "America/Toronto"}
	case "YAU":
		return Location{"Kattiniq", "America/Toronto"}
	case "YAV":
		return Location{"Winnipeg", "America/Winnipeg"}
	case "YAW":
		return Location{"Halifax", "America/Halifax"}
	case "YAX":
		return Location{"Angling Lake", "America/Winnipeg"}
	case "YAY":
		return Location{"St. Anthony", "America/St_Johns"}
	case "YAZ":
		return Location{"Tofino", "America/Vancouver"}
	case "YBA":
		return Location{"Banff", "America/Edmonton"}
	case "YBB":
		return Location{"Kugaaruk", "America/Cambridge_Bay"}
	case "YBC":
		return Location{"Baie-Comeau", "America/Toronto"}
	case "YBE":
		return Location{"Uranium City", "America/Swift_Current"}
	case "YBG":
		return Location{"Bagotville", "America/Toronto"}
	case "YBI":
		return Location{"Black Tickle", "America/St_Johns"}
	case "YBK":
		return Location{"Baker Lake", "America/Rankin_Inlet"}
	case "YBL":
		return Location{"Campbell River", "America/Vancouver"}
	case "YBO":
		return Location{"Bob Quinn Lake", "America/Vancouver"}
	case "YBP":
		return Location{"Yibin", "Asia/Shanghai"}
	case "YBR":
		return Location{"Brandon", "America/Winnipeg"}
	case "YBT":
		return Location{"Brochet", "America/Winnipeg"}
	case "YBV":
		return Location{"Berens River", "America/Winnipeg"}
	case "YBW":
		return Location{"Calgary", "America/Edmonton"}
	case "YBX":
		return Location{"Lourdes-De-Blanc-Sablon", "America/Blanc-Sablon"}
	case "YBY":
		return Location{"Bonnyville", "America/Edmonton"}
	case "YCB":
		return Location{"Cambridge Bay", "America/Cambridge_Bay"}
	case "YCC":
		return Location{"Cornwall", "America/Toronto"}
	case "YCD":
		return Location{"Nanaimo", "America/Vancouver"}
	case "YCE":
		return Location{"Centralia", "America/Toronto"}
	case "YCG":
		return Location{"Castlegar", "America/Vancouver"}
	case "YCH":
		return Location{"Miramichi", "America/Moncton"}
	case "YCK":
		return Location{"Colville Lake", "America/Inuvik"}
	case "YCL":
		return Location{"Charlo", "America/Moncton"}
	case "YCM":
		return Location{"St Catharines", "America/Toronto"}
	case "YCN":
		return Location{"Cochrane", "America/Toronto"}
	case "YCO":
		return Location{"Kugluktuk", "America/Cambridge_Bay"}
	case "YCQ":
		return Location{"Chetwynd", "America/Dawson_Creek"}
	case "YCR":
		return Location{"Cross Lake", "America/Winnipeg"}
	case "YCS":
		return Location{"Chesterfield Inlet", "America/Rankin_Inlet"}
	case "YCT":
		return Location{"Coronation", "America/Edmonton"}
	case "YCU":
		return Location{"Yuncheng", "Asia/Shanghai"}
	case "YCW":
		return Location{"Chilliwack", "America/Vancouver"}
	case "YCY":
		return Location{"Clyde River", "America/Iqaluit"}
	case "YCZ":
		return Location{"Fairmont Hot Springs", "America/Edmonton"}
	case "YDA":
		return Location{"Dawson City", "America/Dawson"}
	case "YDB":
		return Location{"Burwash", "America/Dawson"}
	case "YDC":
		return Location{"Town of Princeton", "America/Vancouver"}
	case "YDF":
		return Location{"Deer Lake", "America/St_Johns"}
	case "YDG":
		return Location{"Digby", "America/Halifax"}
	case "YDL":
		return Location{"Dease Lake", "America/Vancouver"}
	case "YDM":
		return Location{"Ross River", "America/Whitehorse"}
	case "YDN":
		return Location{"Dauphin", "America/Winnipeg"}
	case "YDO":
		return Location{"Dolbeau-St-Felicien", "America/Toronto"}
	case "YDP":
		return Location{"Nain", "America/Goose_Bay"}
	case "YDQ":
		return Location{"Dawson Creek", "America/Dawson_Creek"}
	case "YDT":
		return Location{"Burlington", "America/Toronto"}
	case "YDV":
		return Location{"Bloodvein River", "America/Winnipeg"}
	case "YEB":
		return Location{"Bar River", "America/Toronto"}
	case "YEC":
		return Location{"", "Asia/Seoul"}
	case "YEE":
		return Location{"Midland", "America/Toronto"}
	case "YEG":
		return Location{"Edmonton", "America/Edmonton"}
	case "YEI":
		return Location{"Bursa", "Europe/Istanbul"}
	case "YEK":
		return Location{"Arviat", "America/Rankin_Inlet"}
	case "YEL":
		return Location{"Elliot Lake", "America/Toronto"}
	case "YEM":
		return Location{"Manitowaning", "America/Toronto"}
	case "YEN":
		return Location{"Estevan", "America/Regina"}
	case "YEO":
		return Location{"Yeovil", "Europe/London"}
	case "YER":
		return Location{"Fort Severn", "America/Toronto"}
	case "YES":
		return Location{"", "Asia/Tehran"}
	case "YET":
		return Location{"Edson", "America/Edmonton"}
	case "YEU":
		return Location{"Eureka", "America/Atikokan"}
	case "YEV":
		return Location{"Inuvik", "America/Inuvik"}
	case "YEY":
		return Location{"Amos", "America/Toronto"}
	case "YFA":
		return Location{"Fort Albany", "America/Toronto"}
	case "YFB":
		return Location{"Iqaluit", "America/Iqaluit"}
	case "YFC":
		return Location{"Fredericton", "America/Moncton"}
	case "YFD":
		return Location{"Brantford", "America/Toronto"}
	case "YFE":
		return Location{"Forestville", "America/Toronto"}
	case "YFH":
		return Location{"Fort Hope", "America/Toronto"}
	case "YFI":
		return Location{"Suncor Energy Site", "America/Edmonton"}
	case "YFJ":
		return Location{"Wekweeti", "America/Edmonton"}
	case "YFO":
		return Location{"Flin Flon", "America/Winnipeg"}
	case "YFR":
		return Location{"Fort Resolution", "America/Edmonton"}
	case "YFS":
		return Location{"Fort Simpson", "America/Inuvik"}
	case "YFX":
		return Location{"St. Lewis", "America/St_Johns"}
	case "YGB":
		return Location{"Texada", "America/Vancouver"}
	case "YGD":
		return Location{"Goderich", "America/Toronto"}
	case "YGE":
		return Location{"Golden", "America/Edmonton"}
	case "YGH":
		return Location{"Fort Good Hope", "America/Inuvik"}
	case "YGJ":
		return Location{"Yonago", "Asia/Tokyo"}
	case "YGK":
		return Location{"Kingston", "America/Toronto"}
	case "YGL":
		return Location{"La Grande Riviere", "America/Toronto"}
	case "YGM":
		return Location{"Gimli", "America/Winnipeg"}
	case "YGO":
		return Location{"Gods Lake Narrows", "America/Winnipeg"}
	case "YGP":
		return Location{"Gaspe", "America/Toronto"}
	case "YGQ":
		return Location{"Geraldton", "America/Toronto"}
	case "YGR":
		return Location{"Iles-de-la-Madeleine", "America/Halifax"}
	case "YGT":
		return Location{"Igloolik", "America/Iqaluit"}
	case "YGV":
		return Location{"Havre St-Pierre", "America/Toronto"}
	case "YGW":
		return Location{"Kuujjuarapik", "America/Iqaluit"}
	case "YGX":
		return Location{"Gillam", "America/Winnipeg"}
	case "YGZ":
		return Location{"Grise Fiord", "America/Iqaluit"}
	case "YHA":
		return Location{"Port Hope Simpson", "America/St_Johns"}
	case "YHB":
		return Location{"Hudson Bay", "America/Regina"}
	case "YHD":
		return Location{"Dryden", "America/Winnipeg"}
	case "YHE":
		return Location{"Hope", "America/Vancouver"}
	case "YHF":
		return Location{"Hearst", "America/Toronto"}
	case "YHI":
		return Location{"Ulukhaktok", "America/Edmonton"}
	case "YHJ":
		return Location{"Nanchang", "Asia/Shanghai"}
	case "YHK":
		return Location{"Gjoa Haven", "America/Cambridge_Bay"}
	case "YHM":
		return Location{"Hamilton", "America/Toronto"}
	case "YHN":
		return Location{"Hornepayne", "America/Toronto"}
	case "YHO":
		return Location{"Hopedale", "America/Goose_Bay"}
	case "YHP":
		return Location{"Poplar Hill", "America/Winnipeg"}
	case "YHR":
		return Location{"Chevery", "America/Blanc-Sablon"}
	case "YHT":
		return Location{"Haines Junction", "America/Whitehorse"}
	case "YHU":
		return Location{"Montreal", "America/Toronto"}
	case "YHY":
		return Location{"Hay River", "America/Edmonton"}
	case "YHZ":
		return Location{"Halifax", "America/Halifax"}
	case "YIA":
		return Location{"Yogyakarta", "Asia/Jakarta"}
	case "YIB":
		return Location{"Atikokan", "America/Atikokan"}
	case "YIE":
		return Location{"Arxan", "Asia/Shanghai"}
	case "YIF":
		return Location{"St-Augustin", "America/Blanc-Sablon"}
	case "YIH":
		return Location{"Yichang", "Asia/Shanghai"}
	case "YIK":
		return Location{"Ivujivik", "America/Iqaluit"}
	case "YIN":
		return Location{"Yining", "Asia/Shanghai"}
	case "YIO":
		return Location{"Pond Inlet", "America/Iqaluit"}
	case "YIP":
		return Location{"Detroit", "America/Detroit"}
	case "YIV":
		return Location{"Island Lake", "America/Winnipeg"}
	case "YIW":
		return Location{"Yiwu", "Asia/Shanghai"}
	case "YJA":
		return Location{"Jasper", "America/Edmonton"}
	case "YJF":
		return Location{"Fort Liard", "America/Inuvik"}
	case "YJM":
		return Location{"Fort St. James", "America/Vancouver"}
	case "YJN":
		return Location{"St Jean", "America/Toronto"}
	case "YJS":
		return Location{"Samjiyon", "Asia/Pyongyang"}
	case "YJT":
		return Location{"Stephenville", "America/St_Johns"}
	case "YKA":
		return Location{"Kamloops", "America/Vancouver"}
	case "YKC":
		return Location{"Collins Bay", "America/Regina"}
	case "YKF":
		return Location{"Kitchener", "America/Toronto"}
	case "YKG":
		return Location{"Kangirsuk", "America/Toronto"}
	case "YKJ":
		return Location{"Key Lake", "America/Regina"}
	case "YKL":
		return Location{"Schefferville", "America/Toronto"}
	case "YKM":
		return Location{"Yakima", "America/Los_Angeles"}
	case "YKN":
		return Location{"Yankton", "America/Chicago"}
	case "YKO":
		return Location{"Yuksekova", "Europe/Istanbul"}
	case "YKQ":
		return Location{"Waskaganish", "America/Iqaluit"}
	case "YKS":
		return Location{"Yakutsk", "Asia/Yakutsk"}
	case "YKU":
		return Location{"Chisasibi", "America/Toronto"}
	case "YKX":
		return Location{"Kirkland Lake", "America/Toronto"}
	case "YKY":
		return Location{"Kindersley", "America/Swift_Current"}
	case "YKZ":
		return Location{"Toronto", "America/Toronto"}
	case "YLB":
		return Location{"Lac La Biche", "America/Edmonton"}
	case "YLC":
		return Location{"Kimmirut", "America/Iqaluit"}
	case "YLD":
		return Location{"Chapleau", "America/Toronto"}
	case "YLE":
		return Location{"Whati", "America/Edmonton"}
	case "YLG":
		return Location{"Yalgoo", "Australia/Perth"}
	case "YLH":
		return Location{"Lansdowne House", "America/Toronto"}
	case "YLI":
		return Location{"", "Europe/Helsinki"}
	case "YLJ":
		return Location{"Meadow Lake", "America/Swift_Current"}
	case "YLK":
		return Location{"Barrie-Orillia", "America/Toronto"}
	case "YLL":
		return Location{"Lloydminster", "America/Edmonton"}
	case "YLQ":
		return Location{"La Tuque", "America/Toronto"}
	case "YLR":
		return Location{"Leaf Rapids", "America/Winnipeg"}
	case "YLT":
		return Location{"Alert", "America/Iqaluit"}
	case "YLV":
		return Location{"Yevlakh", "Asia/Baku"}
	case "YLW":
		return Location{"Kelowna", "America/Vancouver"}
	case "YLX":
		return Location{"Yulin", "Asia/Shanghai"}
	case "YLY":
		return Location{"Langley", "America/Vancouver"}
	case "YMA":
		return Location{"Mayo", "America/Whitehorse"}
	case "YME":
		return Location{"Matane", "America/Toronto"}
	case "YMG":
		return Location{"Manitouwadge", "America/Toronto"}
	case "YMH":
		return Location{"Mary's Harbour", "America/St_Johns"}
	case "YMJ":
		return Location{"Moose Jaw", "America/Regina"}
	case "YMK":
		return Location{"Mys Kamennyi", "Asia/Yekaterinburg"}
	case "YML":
		return Location{"Charlevoix", "America/Toronto"}
	case "YMM":
		return Location{"Fort McMurray", "America/Edmonton"}
	case "YMN":
		return Location{"Makkovik", "America/Goose_Bay"}
	case "YMO":
		return Location{"Moosonee", "America/Toronto"}
	case "YMS":
		return Location{"Yurimaguas", "America/Lima"}
	case "YMT":
		return Location{"Chibougamau", "America/Toronto"}
	case "YMW":
		return Location{"Maniwaki", "America/Toronto"}
	case "YMX":
		return Location{"Montreal", "America/Toronto"}
	case "YNA":
		return Location{"Natashquan", "America/Blanc-Sablon"}
	case "YNB":
		return Location{"", "Asia/Riyadh"}
	case "YNC":
		return Location{"Wemindji", "America/Iqaluit"}
	case "YND":
		return Location{"Gatineau", "America/Toronto"}
	case "YNE":
		return Location{"Norway House", "America/Winnipeg"}
	case "YNG":
		return Location{"Youngstown/Warren", "America/New_York"}
	case "YNH":
		return Location{"Hudson's Hope", "America/Dawson_Creek"}
	case "YNJ":
		return Location{"Yanji", "Asia/Shanghai"}
	case "YNL":
		return Location{"Points North Landing", "America/Regina"}
	case "YNM":
		return Location{"Matagami", "America/Toronto"}
	case "YNN":
		return Location{"Nejanilini Lake", "America/Winnipeg"}
	case "YNO":
		return Location{"North Spirit Lake", "America/Winnipeg"}
	case "YNS":
		return Location{"Nemiscau", "America/Toronto"}
	case "YNT":
		return Location{"Yantai", "Asia/Shanghai"}
	case "YNX":
		return Location{"Snap Lake", "America/Edmonton"}
	case "YNY":
		return Location{"Sokcho / Gangneung", "Asia/Seoul"}
	case "YNZ":
		return Location{"Yancheng", "Asia/Shanghai"}
	case "YOA":
		return Location{"Ekati", "America/Edmonton"}
	case "YOC":
		return Location{"Old Crow", "America/Dawson"}
	case "YOD":
		return Location{"Cold Lake", "America/Edmonton"}
	case "YOG":
		return Location{"Ogoki Post", "America/Toronto"}
	case "YOH":
		return Location{"Oxford House", "America/Winnipeg"}
	case "YOJ":
		return Location{"High Level", "America/Edmonton"}
	case "YOL":
		return Location{"Yola", "Africa/Lagos"}
	case "YOO":
		return Location{"Oshawa", "America/Toronto"}
	case "YOP":
		return Location{"Rainbow Lake", "America/Edmonton"}
	case "YOS":
		return Location{"Owen Sound", "America/Toronto"}
	case "YOT":
		return Location{"Yotvata", "Asia/Jerusalem"}
	case "YOW":
		return Location{"Ottawa", "America/Toronto"}
	case "YPA":
		return Location{"Prince Albert", "America/Regina"}
	case "YPC":
		return Location{"Paulatuk", "America/Inuvik"}
	case "YPD":
		return Location{"Parry Sound", "America/Toronto"}
	case "YPE":
		return Location{"Peace River", "America/Edmonton"}
	case "YPG":
		return Location{"Portage", "America/Winnipeg"}
	case "YPH":
		return Location{"Inukjuak", "America/Toronto"}
	case "YPJ":
		return Location{"Aupaluk", "America/Toronto"}
	case "YPL":
		return Location{"Pickle Lake", "America/Atikokan"}
	case "YPM":
		return Location{"Pikangikum", "America/Winnipeg"}
	case "YPN":
		return Location{"Port-Menier", "America/Toronto"}
	case "YPO":
		return Location{"Peawanuck", "America/Toronto"}
	case "YPQ":
		return Location{"Peterborough", "America/Toronto"}
	case "YPR":
		return Location{"Prince Rupert", "America/Vancouver"}
	case "YPS":
		return Location{"Port Hawkesbury", "America/Glace_Bay"}
	case "YPW":
		return Location{"Powell River", "America/Vancouver"}
	case "YPX":
		return Location{"Puvirnituq", "America/Iqaluit"}
	case "YPY":
		return Location{"Fort Chipewyan", "America/Edmonton"}
	case "YPZ":
		return Location{"Burns Lake", "America/Vancouver"}
	case "YQA":
		return Location{"Muskoka", "America/Toronto"}
	case "YQB":
		return Location{"Quebec", "America/Toronto"}
	case "YQC":
		return Location{"Quaqtaq", "America/Iqaluit"}
	case "YQD":
		return Location{"The Pas", "America/Winnipeg"}
	case "YQF":
		return Location{"Red Deer", "America/Edmonton"}
	case "YQG":
		return Location{"Windsor", "America/Toronto"}
	case "YQH":
		return Location{"Watson Lake", "America/Whitehorse"}
	case "YQI":
		return Location{"Yarmouth", "America/Halifax"}
	case "YQK":
		return Location{"Kenora", "America/Winnipeg"}
	case "YQL":
		return Location{"Lethbridge", "America/Edmonton"}
	case "YQM":
		return Location{"Moncton", "America/Moncton"}
	case "YQN":
		return Location{"Nakina", "America/Toronto"}
	case "YQQ":
		return Location{"Comox", "America/Vancouver"}
	case "YQR":
		return Location{"Regina", "America/Regina"}
	case "YQS":
		return Location{"St Thomas", "America/Toronto"}
	case "YQT":
		return Location{"Thunder Bay", "America/Toronto"}
	case "YQU":
		return Location{"Grande Prairie", "America/Edmonton"}
	case "YQV":
		return Location{"Yorkton", "America/Regina"}
	case "YQW":
		return Location{"North Battleford", "America/Swift_Current"}
	case "YQX":
		return Location{"Gander", "America/St_Johns"}
	case "YQY":
		return Location{"Sydney", "America/Glace_Bay"}
	case "YQZ":
		return Location{"Quesnel", "America/Vancouver"}
	case "YRA":
		return Location{"Gameti", "America/Edmonton"}
	case "YRB":
		return Location{"Resolute Bay", "America/Resolute"}
	case "YRF":
		return Location{"Cartwright", "America/Goose_Bay"}
	case "YRG":
		return Location{"Rigolet", "America/Goose_Bay"}
	case "YRI":
		return Location{"Riviere-du-Loup", "America/Toronto"}
	case "YRJ":
		return Location{"Roberval", "America/Toronto"}
	case "YRL":
		return Location{"Red Lake", "America/Winnipeg"}
	case "YRM":
		return Location{"Rocky Mountain House", "America/Edmonton"}
	case "YRO":
		return Location{"Ottawa", "America/Toronto"}
	case "YRP":
		return Location{"Carp", "America/Toronto"}
	case "YRQ":
		return Location{"Trois-Rivieres", "America/Toronto"}
	case "YRS":
		return Location{"Red Sucker Lake", "America/Winnipeg"}
	case "YRT":
		return Location{"Rankin Inlet", "America/Rankin_Inlet"}
	case "YRV":
		return Location{"Revelstoke", "America/Vancouver"}
	case "YSA":
		return Location{"Sable Island", "America/Halifax"}
	case "YSB":
		return Location{"Sudbury", "America/Toronto"}
	case "YSC":
		return Location{"Sherbrooke", "America/Toronto"}
	case "YSE":
		return Location{"Squamish", "America/Vancouver"}
	case "YSF":
		return Location{"Stony Rapids", "America/Regina"}
	case "YSG":
		return Location{"Lutselk'e", "America/Edmonton"}
	case "YSH":
		return Location{"Smiths Falls", "America/Toronto"}
	case "YSJ":
		return Location{"Saint John", "America/Moncton"}
	case "YSK":
		return Location{"Sanikiluaq", "America/Iqaluit"}
	case "YSL":
		return Location{"St Leonard", "America/Moncton"}
	case "YSM":
		return Location{"Fort Smith", "America/Edmonton"}
	case "YSN":
		return Location{"Salmon Arm", "America/Vancouver"}
	case "YSO":
		return Location{"Postville", "America/Goose_Bay"}
	case "YSP":
		return Location{"Marathon", "America/Toronto"}
	case "YSQ":
		return Location{"Songyuan", "Asia/Shanghai"}
	case "YST":
		return Location{"St. Theresa Point", "America/Winnipeg"}
	case "YSU":
		return Location{"Summerside", "America/Halifax"}
	case "YSY":
		return Location{"Sachs Harbour", "America/Inuvik"}
	case "YTA":
		return Location{"Pembroke", "America/Toronto"}
	case "YTD":
		return Location{"Thicket Portage", "America/Winnipeg"}
	case "YTE":
		return Location{"Cape Dorset", "America/Iqaluit"}
	case "YTF":
		return Location{"Alma", "America/Toronto"}
	case "YTH":
		return Location{"Thompson", "America/Winnipeg"}
	case "YTL":
		return Location{"Big Trout Lake", "America/Winnipeg"}
	case "YTM":
		return Location{"Riviere Rouge", "America/Toronto"}
	case "YTQ":
		return Location{"Tasiujaq", "America/Toronto"}
	case "YTR":
		return Location{"Trenton", "America/Toronto"}
	case "YTS":
		return Location{"Timmins", "America/Toronto"}
	case "YTW":
		return Location{"Yutian", "Asia/Shanghai"}
	case "YTZ":
		return Location{"Toronto", "America/Toronto"}
	case "YUB":
		return Location{"Tuktoyaktuk", "America/Inuvik"}
	case "YUD":
		return Location{"Umiujaq", "America/Iqaluit"}
	case "YUE":
		return Location{"", "Australia/Darwin"}
	case "YUL":
		return Location{"Montreal", "America/Toronto"}
	case "YUM":
		return Location{"Yuma", "America/Phoenix"}
	case "YUS":
		return Location{"Yushu", "Asia/Shanghai"}
	case "YUT":
		return Location{"Repulse Bay", "America/Rankin_Inlet"}
	case "YUX":
		return Location{"Hall Beach", "America/Iqaluit"}
	case "YUY":
		return Location{"Rouyn-Noranda", "America/Toronto"}
	case "YVA":
		return Location{"Moroni", "Indian/Comoro"}
	case "YVB":
		return Location{"Bonaventure", "America/Toronto"}
	case "YVC":
		return Location{"La Ronge", "America/Regina"}
	case "YVE":
		return Location{"Vernon", "America/Vancouver"}
	case "YVG":
		return Location{"Vermilion", "America/Edmonton"}
	case "YVM":
		return Location{"Qikiqtarjuaq", "America/Iqaluit"}
	case "YVO":
		return Location{"Val-d'Or", "America/Toronto"}
	case "YVP":
		return Location{"Kuujjuaq", "America/Toronto"}
	case "YVQ":
		return Location{"Norman Wells", "America/Inuvik"}
	case "YVR":
		return Location{"Vancouver", "America/Vancouver"}
	case "YVT":
		return Location{"Buffalo Narrows", "America/Swift_Current"}
	case "YVV":
		return Location{"Wiarton", "America/Toronto"}
	case "YVZ":
		return Location{"Deer Lake", "America/Winnipeg"}
	case "YWA":
		return Location{"Petawawa", "America/Toronto"}
	case "YWB":
		return Location{"Kangiqsujuaq", "America/Toronto"}
	case "YWG":
		return Location{"Winnipeg", "America/Winnipeg"}
	case "YWJ":
		return Location{"Deline", "America/Inuvik"}
	case "YWK":
		return Location{"Wabush", "America/Goose_Bay"}
	case "YWL":
		return Location{"Williams Lake", "America/Vancouver"}
	case "YWM":
		return Location{"Williams Harbour", "America/St_Johns"}
	case "YWP":
		return Location{"Webequie", "America/Toronto"}
	case "YWY":
		return Location{"Wrigley", "America/Inuvik"}
	case "YXC":
		return Location{"Cranbrook", "America/Edmonton"}
	case "YXD":
		return Location{"Edmonton", "America/Edmonton"}
	case "YXE":
		return Location{"Saskatoon", "America/Regina"}
	case "YXH":
		return Location{"Medicine Hat", "America/Edmonton"}
	case "YXJ":
		return Location{"Fort St.John", "America/Dawson_Creek"}
	case "YXK":
		return Location{"Rimouski", "America/Toronto"}
	case "YXL":
		return Location{"Sioux Lookout", "America/Winnipeg"}
	case "YXN":
		return Location{"Whale Cove", "America/Rankin_Inlet"}
	case "YXP":
		return Location{"Pangnirtung", "America/Iqaluit"}
	case "YXQ":
		return Location{"Beaver Creek", "America/Dawson"}
	case "YXR":
		return Location{"Earlton", "America/Toronto"}
	case "YXS":
		return Location{"Prince George", "America/Vancouver"}
	case "YXT":
		return Location{"Terrace", "America/Vancouver"}
	case "YXU":
		return Location{"London", "America/Toronto"}
	case "YXX":
		return Location{"Abbotsford", "America/Los_Angeles"}
	case "YXY":
		return Location{"Whitehorse", "America/Whitehorse"}
	case "YXZ":
		return Location{"Wawa", "America/Toronto"}
	case "YYA":
		return Location{"Yueyang", "Asia/Shanghai"}
	case "YYB":
		return Location{"North Bay", "America/Toronto"}
	case "YYC":
		return Location{"Calgary", "America/Edmonton"}
	case "YYD":
		return Location{"Smithers", "America/Vancouver"}
	case "YYE":
		return Location{"Fort Nelson", "America/Fort_Nelson"}
	case "YYF":
		return Location{"Penticton", "America/Vancouver"}
	case "YYG":
		return Location{"Charlottetown", "America/Halifax"}
	case "YYH":
		return Location{"Taloyoak", "America/Cambridge_Bay"}
	case "YYJ":
		return Location{"Victoria", "America/Vancouver"}
	case "YYL":
		return Location{"Lynn Lake", "America/Winnipeg"}
	case "YYM":
		return Location{"Cowley", "America/Edmonton"}
	case "YYN":
		return Location{"Swift Current", "America/Swift_Current"}
	case "YYQ":
		return Location{"Churchill", "America/Winnipeg"}
	case "YYR":
		return Location{"Goose Bay", "America/Goose_Bay"}
	case "YYT":
		return Location{"St. John's", "America/St_Johns"}
	case "YYU":
		return Location{"Kapuskasing", "America/Toronto"}
	case "YYW":
		return Location{"Armstrong", "America/Toronto"}
	case "YYY":
		return Location{"Mont-Joli", "America/Toronto"}
	case "YYZ":
		return Location{"Toronto", "America/Toronto"}
	case "YZD":
		return Location{"Toronto", "America/Toronto"}
	case "YZE":
		return Location{"Gore Bay", "America/Toronto"}
	case "YZF":
		return Location{"Yellowknife", "America/Edmonton"}
	case "YZG":
		return Location{"Salluit", "America/Toronto"}
	case "YZH":
		return Location{"Slave Lake", "America/Edmonton"}
	case "YZP":
		return Location{"Sandspit", "America/Vancouver"}
	case "YZR":
		return Location{"Sarnia", "America/Toronto"}
	case "YZS":
		return Location{"Coral Harbour", "America/Atikokan"}
	case "YZT":
		return Location{"Port Hardy", "America/Vancouver"}
	case "YZU":
		return Location{"Whitecourt", "America/Edmonton"}
	case "YZV":
		return Location{"Sept-Iles", "America/Toronto"}
	case "YZW":
		return Location{"Teslin", "America/Whitehorse"}
	case "YZX":
		return Location{"Greenwood", "America/Halifax"}
	case "YZY":
		return Location{"Mackenzie", "America/Vancouver"}
	case "YZZ":
		return Location{"Trail", "America/Vancouver"}
	case "ZAC":
		return Location{"York Landing", "America/Winnipeg"}
	case "ZAD":
		return Location{"Zadar", "Europe/Zagreb"}
	case "ZAG":
		return Location{"Zagreb", "Europe/Zagreb"}
	case "ZAH":
		return Location{"Zahedan", "Asia/Tehran"}
	case "ZAJ":
		return Location{"Zaranj", "Asia/Kabul"}
	case "ZAL":
		return Location{"Valdivia", "America/Santiago"}
	case "ZAM":
		return Location{"Zamboanga City", "Asia/Manila"}
	case "ZAO":
		return Location{"Cahors/Lalbenque", "Europe/Paris"}
	case "ZAR":
		return Location{"Zaria", "Africa/Lagos"}
	case "ZAT":
		return Location{"Zhaotong", "Asia/Shanghai"}
	case "ZAZ":
		return Location{"Zaragoza", "Europe/Madrid"}
	case "ZBD":
		return Location{"Boundary Bay", "America/Vancouver"}
	case "ZBE":
		return Location{"Zabreh", "Europe/Prague"}
	case "ZBF":
		return Location{"Bathurst", "America/Moncton"}
	case "ZBM":
		return Location{"Bromont", "America/Toronto"}
	case "ZBO":
		return Location{"", "Australia/Brisbane"}
	case "ZBR":
		return Location{"Chabahar", "Asia/Tehran"}
	case "ZBY":
		return Location{"Sayaboury", "Asia/Vientiane"}
	case "ZCA":
		return Location{"Neheim-Husten", "Europe/Berlin"}
	case "ZCL":
		return Location{"Zacatecas", "America/Mexico_City"}
	case "ZCN":
		return Location{"", "Europe/Berlin"}
	case "ZCO":
		return Location{"Temuco", "America/Santiago"}
	case "ZEC":
		return Location{"Secunda", "Africa/Johannesburg"}
	case "ZEE":
		return Location{"Kelsey", "America/Winnipeg"}
	case "ZEL":
		return Location{"Bella Bella", "America/Vancouver"}
	case "ZEM":
		return Location{"Eastmain River", "America/Iqaluit"}
	case "ZER":
		return Location{"", "Asia/Kolkata"}
	case "ZFA":
		return Location{"Faro", "America/Whitehorse"}
	case "ZFD":
		return Location{"Fond-Du-Lac", "America/Swift_Current"}
	case "ZFG":
		return Location{"Pukatawagan", "America/Winnipeg"}
	case "ZFL":
		return Location{"Zhaosu", "Asia/Shanghai"}
	case "ZFM":
		return Location{"Fort Mcpherson", "America/Inuvik"}
	case "ZFN":
		return Location{"Tulita", "America/Inuvik"}
	case "ZGF":
		return Location{"Grand Forks", "America/Los_Angeles"}
	case "ZGI":
		return Location{"Gods River", "America/Winnipeg"}
	case "ZGL":
		return Location{"", "Australia/Brisbane"}
	case "ZGM":
		return Location{"Ngoma", "Africa/Lusaka"}
	case "ZGR":
		return Location{"Little Grand Rapids", "America/Winnipeg"}
	case "ZGU":
		return Location{"Gaua Island", "Pacific/Efate"}
	case "ZHA":
		return Location{"Zhanjiang", "Asia/Shanghai"}
	case "ZHI":
		return Location{"", "Europe/Zurich"}
	case "ZHM":
		return Location{"Shamshernagar", "Asia/Dhaka"}
	case "ZHP":
		return Location{"High Prairie", "America/Edmonton"}
	case "ZHY":
		return Location{"Zhongwei", "Asia/Shanghai"}
	case "ZHZ":
		return Location{"Oppin", "Europe/Berlin"}
	case "ZIA":
		return Location{"Zhukovsky", "Europe/Moscow"}
	case "ZIC":
		return Location{"Victoria", "America/Santiago"}
	case "ZIG":
		return Location{"Ziguinchor", "Africa/Dakar"}
	case "ZIH":
		return Location{"Ixtapa", "America/Mexico_City"}
	case "ZIN":
		return Location{"", "Europe/Zurich"}
	case "ZIX":
		return Location{"Zhigansk", "Asia/Yakutsk"}
	case "ZJG":
		return Location{"Jenpeg", "America/Winnipeg"}
	case "ZJI":
		return Location{"", "Europe/Zurich"}
	case "ZJN":
		return Location{"Swan River", "America/Winnipeg"}
	case "ZKB":
		return Location{"Kasaba Bay", "Africa/Lusaka"}
	case "ZKE":
		return Location{"Kashechewan", "America/Toronto"}
	case "ZKG":
		return Location{"Kegaska", "America/Blanc-Sablon"}
	case "ZKM":
		return Location{"Sette Cama", "Africa/Libreville"}
	case "ZKP":
		return Location{"Zyryanka", "Asia/Srednekolymsk"}
	case "ZLO":
		return Location{"Manzanillo", "America/Mexico_City"}
	case "ZLR":
		return Location{"Linares", "America/Santiago"}
	case "ZLT":
		return Location{"La Tabatiere", "America/Blanc-Sablon"}
	case "ZLX":
		return Location{"Zalingei", "Africa/Khartoum"}
	case "ZMG":
		return Location{"Magdeburg", "Europe/Berlin"}
	case "ZMH":
		return Location{"108 Mile", "America/Vancouver"}
	case "ZMM":
		return Location{"", "America/Mexico_City"}
	case "ZMT":
		return Location{"Masset", "America/Vancouver"}
	case "ZND":
		return Location{"Zinder", "Africa/Niamey"}
	case "ZNE":
		return Location{"Newman", "Australia/Perth"}
	case "ZNF":
		return Location{"", "Europe/Berlin"}
	case "ZNG":
		return Location{"Poplar River", "America/Winnipeg"}
	case "ZNV":
		return Location{"Koblenz", "Europe/Berlin"}
	case "ZNZ":
		return Location{"Kiembi Samaki", "Africa/Dar_es_Salaam"}
	case "ZOJ":
		return Location{"Marl", "Europe/Berlin"}
	case "ZOS":
		return Location{"Osorno", "America/Santiago"}
	case "ZOW":
		return Location{"Klausheide", "Europe/Berlin"}
	case "ZPB":
		return Location{"Sachigo Lake", "America/Winnipeg"}
	case "ZPC":
		return Location{"Pucon", "America/Santiago"}
	case "ZPH":
		return Location{"Zephyrhills", "America/New_York"}
	case "ZPO":
		return Location{"Pinehouse Lake", "America/Regina"}
	case "ZPQ":
		return Location{"", "Europe/Berlin"}
	case "ZQC":
		return Location{"Speyer", "Europe/Berlin"}
	case "ZQF":
		return Location{"Trier", "Europe/Berlin"}
	case "ZQL":
		return Location{"Donaueschingen", "Europe/Berlin"}
	case "ZQN":
		return Location{"Queenstown", "Pacific/Auckland"}
	case "ZQW":
		return Location{"Zweibrucken", "Europe/Berlin"}
	case "ZRE":
		return Location{"Zrenjanin", "Europe/Belgrade"}
	case "ZRH":
		return Location{"Zurich", "Europe/Zurich"}
	case "ZRI":
		return Location{"Serui-Japen Island", "Asia/Jayapura"}
	case "ZRJ":
		return Location{"Round Lake", "America/Winnipeg"}
	case "ZRM":
		return Location{"Sarmi-Papua Island", "Asia/Jayapura"}
	case "ZSA":
		return Location{"San Salvador", "America/Nassau"}
	case "ZSE":
		return Location{"St Pierre", "Indian/Reunion"}
	case "ZSJ":
		return Location{"Sandy Lake", "America/Winnipeg"}
	case "ZSP":
		return Location{"St. Paul", "America/Edmonton"}
	case "ZSS":
		return Location{"Sassandra", "Africa/Abidjan"}
	case "ZST":
		return Location{"Stewart", "America/Vancouver"}
	case "ZTA":
		return Location{"", "Pacific/Tahiti"}
	case "ZTB":
		return Location{"Tete-a-la-Baleine", "America/Blanc-Sablon"}
	case "ZTH":
		return Location{"Zakynthos Island", "Europe/Athens"}
	case "ZTM":
		return Location{"Shamattawa", "America/Winnipeg"}
	case "ZTR":
		return Location{"", "Europe/Kyiv"}
	case "ZTU":
		return Location{"Zaqatala", "Asia/Baku"}
	case "ZUC":
		return Location{"Ignace", "America/Winnipeg"}
	case "ZUD":
		return Location{"Ancud", "America/Santiago"}
	case "ZUH":
		return Location{"Zhuhai", "Asia/Shanghai"}
	case "ZUL":
		return Location{"Zilfi", "Asia/Riyadh"}
	case "ZUM":
		return Location{"Churchill Falls", "America/Goose_Bay"}
	case "ZVA":
		return Location{"", "Indian/Antananarivo"}
	case "ZVK":
		return Location{"", "Asia/Bangkok"}
	case "ZWA":
		return Location{"", "Indian/Antananarivo"}
	case "ZWH":
		return Location{"Lac Brochet", "America/Winnipeg"}
	case "ZWK":
		return Location{"Suwalki", "Europe/Warsaw"}
	case "ZWL":
		return Location{"Wollaston Lake", "America/Regina"}
	case "ZXB":
		return Location{"", "Europe/Oslo"}
	case "ZYI":
		return Location{"Zunyi", "Asia/Shanghai"}
	case "ZYL":
		return Location{"Sylhet", "Asia/Dhaka"}
	case "ZZE":
		return Location{"Zangilan", "Asia/Baku"}
	case "ZZO":
		return Location{"Tymovskoye", "Asia/Sakhalin"}
	case "ZZU":
		return Location{"Mzuzu", "Africa/Blantyre"}
	case "ZZV":
		return Location{"Zanesville", "America/New_York"}
	}
	return Location{"Not supported IATA Code", "Not supported IATA Code"}
}
