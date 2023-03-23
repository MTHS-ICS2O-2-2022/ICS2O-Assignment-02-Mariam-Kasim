// Copyright (c) 2020 Mr. Coxall All rights reserved
//
// Created by: Mariam Kasim
// Created on: Mar 2023
// This file contains the JS functions for index.html


'use strict'
/**
* This function calculates area of a square
*/
function calculate () {
 // input
 const sidea = parseInt(document.getElementById('side-a').value)
 const sideb = parseInt(document.getElementById('side-b').value)


 // process
 const area = sidea * sideb


 // output
 document.getElementById('area').innerHTML = 'Area is: ' + area + ' cm²'
}


