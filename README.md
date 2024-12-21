Connect to docker cli

    docker exec -it yuth-optician-postgresql bash

Connect to mongodb

    psql -U <username>

เข้าใช้งานด้วยคำสั่ง $psql -U <username> -d <database> -W
- \l ดู database ทั้งหมด
- \c <database> เข้าใช้งาน database name ที่ต้องการ
- \dt ดู table ทั้งหมดใน database ปัจจุบัน
- \dt+ ดู size และ description ของ table เพิ่มเติม
- \d <table name> ดูโครงสร้างของ table
- \dv ดู view ทั้งหมด
- \df ดู function ทั้งหมด
- \dx ดู extension ที่ติดตั้งทั้งหมด
- \du ดู user ทั้งหมด
- \e ทำการเปิด editor เพื่อแก้ไข command
- \s <file name> บันทึก comand ลง file
- \! clear สำหรับ clear หน้าจอ
- \h => Help
- \? ดูว่ามี command อะไรให้ใช้งานบ้าง
- \q ออก เลิกใช้งาน …

Database

    db.items.insertMany([
    {
        "name": "Sword of Valor",
        "description": "One-Handed Sword",
        "damage": 120,
        "level_required": 30,
        "price": 1500
    },
    {
        "name": "Elven Bow",
        "description": "Bow",
        "damage": 90,
        "level_required": 25,
        "price": 1200
    },
    {
        "name": "Staff of Fire",
        "description": "Staff",
        "damage": 105,
        "level_required": 28,
        "price": 1400
    },
    {
        "name": "Dwarven Hammer",
        "description": "Two-Handed Hammer",
        "damage": 160,
        "level_required": 35,
        "price": 1800
    },
    {
        "name": "Assassin's Dagger",
        "description": "Dagger",
        "damage": 80,
        "level_required": 22,
        "price": 1000
    }
    ])