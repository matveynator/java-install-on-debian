
# 📘 Java Install on Debian

> 🌐 Select language / Выберите язык:  
> 🇬🇧 [English](#english) | 🇷🇺 [Русский](#русский)

---

## English

### ⚡ Quick Install (x86_64 Debian-based systems)

```bash
sudo curl -fsSL http://files.zabiyaka.net/java-install-on-debian/latest/no-gui/linux/amd64/java-install-on-debian -o /usr/local/bin/java-install-on-debian && sudo chmod +x /usr/local/bin/java-install-on-debian && sudo java-install-on-debian
```

### 📂 Manual Run

If the script is already installed, you can run it anytime with:

```bash
sudo java-install-on-debian
```

### ✅ After install: Set default Java version

```bash
sudo update-alternatives --config java
sudo update-alternatives --config javac
```

### 📌 Features

- Install Java SE 7–23
- No GUI required
- Works on servers and terminals
- Adds Java to `/usr/lib/jvm`
- Simple version selection via menu

### 🔧 Requirements

- Debian/Ubuntu or derivative
- `bash`, `curl`, `tar`
- Internet connection

---

## Русский

### ⚡ Быстрая установка (x86_64 на Debian и Ubuntu)

```bash
sudo curl -fsSL http://files.zabiyaka.net/java-install-on-debian/latest/no-gui/linux/amd64/java-install-on-debian -o /usr/local/bin/java-install-on-debian && sudo chmod +x /usr/local/bin/java-install-on-debian && sudo java-install-on-debian
```

### 📂 Ручной запуск

Если скрипт уже установлен, запустить его можно в любое время так:

```bash
sudo java-install-on-debian
```

### ✅ После установки: выбрать версию Java

```bash
sudo update-alternatives --config java
sudo update-alternatives --config javac
```

### 📌 Возможности

- Установка Java SE от 7 до 23
- Без графического интерфейса
- Подходит для серверов и терминалов
- Java добавляется в `/usr/lib/jvm`
- Простой выбор версии через меню

### 🔧 Требования

- Debian / Ubuntu
- Установлены `bash`, `curl`, `tar`
- Интернет-соединение
