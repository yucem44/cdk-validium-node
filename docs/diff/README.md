# Diff documentation

The CDK Validium node is a fork from the [zkevm-node](https://github.com/0xPolygonHermez/zkevm-node). This documentation will go through the main changes introduced by the fork.

A part from the explanation provided in this document, there is a web UI that shows the actual (relevant diff). In order to generate the UI, run:

1. `npm install -g diff2html-cli` (only once, to install the software)
2. 
```bash
diff -ruN \
-I ".*github.com\/0x.*" \
-x ".git" \
-x ".github" \
-x ".gitignore" \
-x ".vscode" \
-x "ci" \
-x "environments" \
-x "*.md" \
-x "*.html" \
-x "*.html" \
-x "*.json" \
-x "*.toml" \
-x "*.abi" \
-x "*.bin" \
-x "*.pb.go" \
-x "smartcontracts" \
-x "go.sum" \
-x "mock*.go" \
<path/to/zkevm-node> <path/to/cdk-validium-node> | \
diff2html -i stdin -s side -t "zkEVM node vs CDK validium node</br><h2>v0.3.0<h2/>" \
-F diff-zkevm-vs-cdk
```

**NOTE:** before running the command, replace `<path/to/zkevm-node>` and `<path/to/cdk-validium-node>` with the paths where you've cloned the repos on your machine. Also, remember to checkout the relevant branches/tags on each repo