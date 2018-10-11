# PAYU BIZ UTIL

Prostředník mezi klientem a PayU.

Vycházíme z https://documentation.payubiz.in/embedded-form/

Prostředník reaguje na GET request `/payu` a odešle POST request na PayU.

## Povinné parametry
- `txnid` : string (uuid)
- `amount` : number (INR)
- `firstname` : string
- `email` : string
- `phone` : string



## Prostředník reaguje i na další dva GET requesty např.

- `/payu/success` - podle toho na klientovi zjistím jestli je vše v pořádku a na server udělám `placeNewOrder` s daným `uuid`
- `/payu/failed` - vypíšu chybu

## Flow
1. Já v klientu vygeneruju `uuid` a posílám povinné informace na prostředníka pomocí GET
2. Server vygeneruje ověřovací `hash` a posílá a vyrenderuje HTML, kde je <form method=POST> a tím posílám request na PAYU (při onload)
3. Tady už mám jejich webovku a vyberu si platbu dle libosti. Takže moc nevím co se tu děje nebo ne

4. Celé to končí buď na `surl` nebo `furl` nebo tím, že uživatel zavře **webview**, ve kterém to celé renderujeme

## Já myslím, že to je easy a na začátek použitelné





# TEMP větev na POST /upload
Udělal jsem dočasné špatné řešení, ale abych vůbec byl schopný udělat order.
Jelikož DirectApp umí upload pouze pomocí FORMDATA nebo STRING (base64).
Udělal jsem zase prostředníka (později by šlo vyřešit např. Lambdou přímo na bucketu).

Očekává POST na `/upload` kde mu přijde `file` a `url`.
Jedině co dělá ta PUT request a nahraje file na S3.
Je to sice dlouhé a 2x tak náročné, ale zatím nic lepšího nejde.
