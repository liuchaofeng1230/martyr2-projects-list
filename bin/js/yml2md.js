#!/usr/bin/env node

var
    fs     = require('fs'),
    yaml   = require('js-yaml'),
    pastry = require('pastry'),
    resultArr = [],
    doc, projects, pkg;

try {
    pkg = JSON.parse(fs.readFileSync('package.json'));

    resultArr.push('# ' + pkg.name);
    resultArr.push(pkg.description);

    doc = yaml.safeLoad(fs.readFileSync('doc/yml/projects.yml', 'utf8'));
    projects = doc.map;

    pastry.each(projects, function (category) {
        resultArr.push('## ' + category.category);
        pastry.each(category.contents, function (sort) {
            resultArr.push('### ' + sort.sort);
            pastry.each(sort.links, function (link) {
                resultArr.push('#### ' + link.name);
                resultArr.push(link.description);
            });
        });
    });

    fs.writeFile('doc/md/projects.md', resultArr.join('\r\n\r\n'), function (err) {
        if(err) {
            console.log(err);
        }
    });
} catch (e) {
    console.log(e);
}

