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

**Možná i tyhle, to se domluvíme viz níže, jestli je nutné je posílat nebo je mít jen na serveru**
- `surl` : string (redirect v případě **OK**)
- `furl` : string (redirect v případě **FAIL**)

## K diskuzi
- `signature` : string (generovat na klientovi `signature` a ověřovat ho poté na serveru tyhle parametry, jestli přišli v pořádku)

## Prostředník reaguje i na další dva GET requesty např.

- `/payu/success` - podle toho zjistím jestli v pořádku a na server udělám `placeNewOrder` s daným `uuid`
- `/payu/failed` - vypíšu chybu

## Flow
1. Já v klientu vygeneruju `uuid` a posílám povinné informace na prostředníka
2. Ten vygeneruje ověřovací `hash` a posílá POST request na PAYU
3. Tady už mám jejich webovku a vyberu si platbu dle libosti. Takže moc nevím co se tu děje nebo ne

4. Celé to končí buď na `surl` nebo `furl` nebo tím, že uživatel zavře **webview**, ve kterém to celé renderujeme

## Já myslím, že to je easy a na začátek použitelné
