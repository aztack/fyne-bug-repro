const cp = require('child_process');
const fs = require('fs');
const path = require('path')
const projectRoot = path.resolve(__dirname, '..');
const ver = require(path.resolve(projectRoot, 'package.json')).version;
const targetPlatform = process.argv[2] || 'darwin';
const _appName = 'Bug Reproduction';
const appName = targetPlatform === 'darwin' ? `${_appName}.app` : `${_appName}.exe`;

process.cwd(projectRoot)
if (!fs.existsSync(appName)) {
    console.error(`${appName} does not exists, please build it first.`)
    process.exit(-1);
}
let cmd = []
const targetZip = `./dist/${_appName}-${ver}-${targetPlatform}.zip`;
try {
    if (targetPlatform === 'darwin') {
        cmd.push('zip -r', `"${targetZip}"`, `"${appName}"`, '-x "*.DS_Store"')
    } else if (targetPlatform === 'win32') {
        cmd.push('zip', `"${targetZip}"`, `"${appName}"`, 'config.json')
    }
    cmd = cmd.join(' ');
    console.log(cmd);
    cp.execSync(cmd);
} catch (e) {
    console.error(e)
}