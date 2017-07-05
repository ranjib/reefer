import React from 'react'
import $ from 'jquery'

export default class Admin extends React.Component {
  constructor (props) {
    super(props)
    this.powerOff = this.powerOff.bind(this)
    this.reboot = this.reboot.bind(this)
  }

  powerOff () {
    $.ajax({
      url: '/api/admin/poweroff',
      type: 'POST',
      success: function (data) {
      },
      error: function (xhr, status, err) {
        console.log(err.toString())
      }
    })
  }
  reboot () {
    $.ajax({
      url: '/api/admin/reboot',
      type: 'POST',
      success: function (data) {
      },
      error: function (xhr, status, err) {
        console.log(err.toString())
      }
    })
  }

  render () {
    return (
      <div className='container'>
        <input value='Poweroff' onClick={this.powerOff} type='button' className='btn btn-outline-danger' />
        <input value='Reboot' onClick={this.reboot} type='button' className='btn btn-outline-danger' />
      </div>
    )
  }
}
