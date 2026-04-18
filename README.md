# KBBI CLI

CLI tool sederhana untuk mencari definisi kata dalam Kamus Besar Bahasa Indonesia (KBBI) secara offline dan cepat menggunakan Go.

## Fitur
- **Offline:** Tidak memerlukan koneksi internet (Database V6 sudah di-embed).
- **Cepat:** Hasil pencarian instan menggunakan Go.
- **Stand-alone:** Hanya satu file binary, tidak perlu install database terpisah.

## Instalasi

### Binary (Recommended)
Unduh binary terbaru dari halaman [Releases](https://github.com/haidaralfarizy/kbbi-cli/releases), lalu pindahkan ke path Anda:

```bash
# Contoh untuk Linux/macOS
chmod +x kbbi
sudo mv kbbi /usr/local/bin/
```

### Source (Go 1.16+)
```bash
git clone https://github.com/haidaralfarizy/kbbi-cli.git
cd kbbi-cli
go build -o kbbi
```

## Penggunaan
```bash
kbbi <kata>
```

Contoh:
```bash
kbbi merdeka
```

## Credits
Data KBBI berasal dari repositori [HirziDevs/alita-sambung-kata](https://github.com/HirziDevs/alita-sambung-kata) yang diekstrak dari database KBBI resmi.

## License
MIT
