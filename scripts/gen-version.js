const fs = require('fs');
const path = require('path')
const projectRoot = path.resolve(__dirname, '..');
const ver = require(path.resolve(projectRoot, 'package.json')).version;

process.cwd(projectRoot)
fs.writeFileSync(path.resolve(projectRoot, "version.go"), `
// Auto-Generated, DO NOT MODIFY
package main

const Version = "${ver}"
`);