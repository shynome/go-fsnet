# Changelog

## [0.0.2] - 2023-03-24

### Fix

- `fsnet.File.IsDir()` 现在从 `fsnet.File.Mode()` 中返回是否为文件夹, 避免`ModeTpye`检查不一致
- `fsnet.ModeType` 现在不再错误的包含 `fs.ModeDir`
