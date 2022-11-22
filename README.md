# Rename Orc

DFIR-ORC when collecting artifacts will rename the files to a unique name by adding a large prefix and a large suffix to the original file name. This mechanism is used to avoid collisions when multiple files with the same name are collected. However it can be a problem when the artefact is meant to be process that is expecting the original filename.

This tool is a simple tool that will rename the files back to their original name. If there a collision, it will add a suffix like _(1) to the file name.

## Usage

```
orc-rename -dir <directory>
```
