## 📘 Wiki: Java Install on Debian

> 🌐 Select language / Выберите язык:  
> 🇬🇧 [English](#english) | 🇷🇺 [Русский](#русский)

---

## 🇬🇧 English

### ⚡ Quick Install (x86_64 Debian-based systems)

```bash
sudo curl -fsSL http://files.zabiyaka.net/java-install-on-debian/latest/no-gui/linux/amd64/java-install-on-debian -o /usr/local/bin/java-install-on-debian && sudo chmod +x /usr/local/bin/java-install-on-debian && sudo java-install-on-debian
```

### ✅ After install: Set default Java version

```bash
sudo update-alternatives --config java
sudo update-alternatives --config javac
```

### 📌 Features

- Install Java SE 7–23
- No GUI required
- Automatically detects and adds new Java to `/usr/lib/jvm`
- Works on servers and terminals

### 🔧 Requirements

- Debian/Ubuntu (or derivative)
- `bash`, `curl`, `tar`
- Internet connection

---

## 🇷🇺 Русский

### ⚡ Быстрая установка (x86_64 на Debian и Ubuntu)

```bash
sudo curl -fsSL http://files.zabiyaka.net/java-install-on-debian/latest/no-gui/linux/amd64/java-install-on-debian -o /usr/local/bin/java-install-on-debian && sudo chmod +x /usr/local/bin/java-install-on-debian && sudo java-install-on-debian
```

### ✅ После установки: выбрать версию Java

```bash
sudo update-alternatives --config java
sudo update-alternatives --config javac
```

### 📌 Возможности

- Установка Java SE от 7 до 23
- Без графики, подходит для серверов
- Установленные версии попадают в `/usr/lib/jvm`
- Поддерживает переключение версий

### 🔧 Требования

- Debian / Ubuntu
- Утилиты: `bash`, `curl`, `tar`
- Доступ в интернет
