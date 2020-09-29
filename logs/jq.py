import base64

import hashlib

import hmac

import time

import urllib.request

##钉钉sha256签名

timestamp = round(time.time() * 1000)

secret = 'this is secret'
print('secret ',secret )
secret_enc = bytes(secret,encoding='UTF-8')
print('secret_enc ',secret_enc )
string_to_sign = '{}\n{}'.format(timestamp, secret)
print('string_to_sign ',string_to_sign  )
string_to_sign_enc = bytes(string_to_sign,encoding='utf-8')
print('string_to_sign_enc  ',string_to_sign_enc   )
hmac_code = hmac.new(secret_enc, string_to_sign_enc, digestmod=hashlib.sha256).digest()
print('hmac_code',hmac_code)
sign = urllib.request.quote(base64.b64encode(hmac_code))

print(timestamp)

print(sign)
