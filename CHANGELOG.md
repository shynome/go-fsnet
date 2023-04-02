# Changelog

## [1.0.0] - 2023-04-02

### Improve

- 非 wasm 目标现在可以按照预期访问网络, 不再出现文件不存在的错误了. 这令你可以使用一致的代码生成 cgi 和 wagi 程序
- 添加 fsnet wasm 测试, 现在可以自信的发布 1.0.0 版本了

## [0.0.2] - 2023-03-24

### Fix

- `fsnet.File.IsDir()` 现在从 `fsnet.File.Mode()` 中返回是否为文件夹, 避免`ModeTpye`检查不一致
- `fsnet.ModeType` 现在不再错误的包含 `fs.ModeDir`
