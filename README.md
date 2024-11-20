# chat
Chat Ilovasi Backend
Umumiy Ma'lumot

Bu loyiha Chat Ilovasi uchun backend qismi bo‘lib, foydalanuvchilarni autentifikatsiya qilish, guruh va xabarlarni boshqarish, hamda fayllarni saqlash funksiyalarini amalga oshiradi. Loyihada Go dasturlash tili va PostgreSQL ma'lumotlar bazasidan foydalanilgan. Kod modulli arxitekturaga asoslangan bo‘lib, kengaytirish va texnik xizmat ko‘rsatishni osonlashtiradi.

# Xususiyatlari

    Autentifikatsiya:
        Foydalanuvchilarni ro‘yxatdan o‘tkazish va tizimga kirishini ta'minlash.
        JWT token asosida autentifikatsiya qilish.

    Foydalanuvchilarni boshqarish:
        Foydalanuvchi profillari bilan CRUD operatsiyalarini bajarish.
        Rolga asoslangan kirish huquqlarini boshqarish.

    Guruhlarni boshqarish:
        Guruhlarni yaratish va boshqarish.
        Foydalanuvchilarni guruhlarga qo‘shish va olib tashlash.

    Xabarlar:
        Guruhlarda xabar yuborish va qabul qilish.
        Xabar tarixini saqlash.

    Fayllarni boshqarish:
        Xabarlarga biriktirilgan fayllarni saqlash va yuklab olish.

# Texnologiyalar

    Dasturlash tili: Go
    Ma'lumotlar bazasi: PostgreSQL
    Arxitektura: Modulli va interfeysga asoslangan tuzilma

# O‘rnatish

    Loyihani yuklab oling:

git clone https://github.com/Mubinabd/chat.git
cd chat

# Konfiguratsiyani sozlang:

    config.yaml faylini tahrirlang va o‘z ma’lumotlar bazasi va sozlamalaringizni kiriting.


# Ilovani ishga tushirish:

    docker compose up

# Muallif

    Mubina Bahodirova
    