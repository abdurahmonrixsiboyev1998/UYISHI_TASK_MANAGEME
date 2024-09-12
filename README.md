# Uy vazifasi: WebSocket xabarnomalari bilan vazifalarni boshqarish tizimi

## Topshiriq Tavsifi:
Real vaqtda yangilanishlarni ta'minlovchi, WebSocket xabarnomalari bilan integratsiyalangan vazifalarni boshqarish tizimini yaratish.


## Talablar:
1. **Vazifalarni boshqarish::**
    - Foydalanuvchilar quyidagi imkoniyatlarga ega bo'lishlari kerak:
        - Vazifa yaratish.
        - Vazifani yangilash.
        - Vazifani o'chirish.
        - Vazifalar ro'yxatini ko'rish.
    - Har bir vazifa quyidagi maydonlarga ega bo'lishi kerak:
        - `title`: Vazifa qisqa tavsifi.
        - `description`: Vazifaning batafsil tavsifi.
        - `status`: Vazifa holati (masalan, "`Kutmoqda`", "`Bajarilmoqda`", "`Bajarildi`").
        - `due_date`: Vazifaning bajarilishi kerak bo'lgan oxirgi muddat (ixtiyoriy).

2. **WebSocket xabarnomalari**
    - WebSocket orqali real vaqt rejimida barcha ulangan mijozlarga quyidagi yangilanishlar haqida xabar yuborilishi kerak:
        - Vazifa yaratilganida.
        - Vazifa yangilanganida.
        - Vazifa o'chirilganida.
