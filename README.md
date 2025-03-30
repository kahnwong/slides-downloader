# slides-downloader

```bash
fd txt | xargs -I {} aria2c --input-file {} --auto-file-renaming=false --continue=true --check-integrity=true
```

## Parameters

## sched

```bash
PARALLELISM=1
RANDOM_DELAY_SECOND=16
```
