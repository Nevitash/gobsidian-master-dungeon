# Gobsidian Master Dungeon _(gobsidian-master-dungeon)_

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

CLI tool that merges Obsidian vault notes into a single formatted Markdown file.

Gobsidian Master Dungeon consolidates notes, dossiers, and dungeon materials from an Obsidian.md knowledge vault into unified documents, supporting path filters, heading prefixes, and custom Go templates.

## Table of Contents

- [Background](#background)
- [Install](#install)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contributing](#contributing)
- [License](#license)

## Background

Game masters and authors maintaining extensive campaign vaults in Obsidian often need a clean, single-document compilation for offline reading, printing, or sending to external tabletop tools. Gobsidian Master Dungeon automates vault crawling, link resolution, and templated export into a standalone Markdown file.

## Install

Requires Go 1.25 or newer:

```bash
git clone https://github.com/Nevitash/gobsidian-master-dungeon.git
cd gobsidian-master-dungeon
go build -o gobsidian-master-dungeon.exe
```

## Usage

Basic usage to merge an entire vault:

```bash
./gobsidian-master-dungeon -v "/path/to/obsidian/vault" -o "./merged-vault.md"
```

With path filtering and custom templating:

```bash
./gobsidian-master-dungeon -v "./vault" -o "./dungeon-handout.md" --include-file "*.md" --prefix-headings
```

## Maintainers

[@Nevitash](https://github.com/Nevitash)

## Contributing

Pull requests are welcome. For major architectural changes or flags, please open an issue first.

## License

[MIT](LICENSE)
