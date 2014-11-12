#!/usr/bin/env node

var
    fs   = require('fs'),
    yaml = require('js-yaml'),
    doc, jsonStr;

try {
    doc = yaml.safeLoad(fs.readFileSync('doc/yml/projects.yml', 'utf8'));
    jsonStr = JSON.stringify(doc, null, 4);
    fs.writeFile('doc/json/projects.json', jsonStr, function (err) {
        if(err) {
            console.log(err);
        }
    });
} catch (e) {
    console.log(e);
}

