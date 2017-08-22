const Nightmare = require('nightmare')
const nightmare = Nightmare({ show: true })

function test_equipments() {
  nightmare
    .goto('http://localhost:8080/')
    .wait(500)
    .click('input#add_equipment')
    .wait(500)
    .type('input#equipmentName', 'HoBFilter')
    .wait(500)
    .click('button#outlet')
    .wait(500)
    .click('div#react-tabs-1 > div.container:nth-child(1) > div:nth-child(2) > div:nth-child(2) > div.dropdown.open.btn-group:nth-child(2) > ul.dropdown-menu:nth-child(2) > li:nth-child(1) > a:nth-child(1)')
    .wait(500)
    .click('input#createEquipment')
    .wait(500)
    .click('input#HoBFilter')
    .wait(500)
    .click('input#HoBFilter')
    .wait(500)
    .click('input#HoBFilter')
    .wait(500)
    .click('input#equipment-0')
    .wait(2500)
    .end()
      .then(function (result) {
        console.log(result)
      })
      .catch(function (error) {
        console.error('Error:', error)
      })
}
module.exports = test_equipments;
