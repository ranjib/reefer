import React from 'react'
import Common from './common.jsx'
import Display from './display.jsx'
import TemperatureChart from './temperature_chart.jsx'
import EquipmentsChart from './equipment_chart.jsx'

export default class Dashboard extends Common {
  constructor (props) {
    super(props)
    this.state = {
      info: {}
    }
    this.refresh = this.refresh.bind(this)
    this.showDisplay = this.showDisplay.bind(this)
    this.showCharts = this.showCharts.bind(this)
  }

  showCharts () {
    var charts = []
    if (this.props.capabilities.temperature) {
      charts.push(<TemperatureChart key={'chart-1'} />)
    }
    if (this.props.capabilities.equipments) {
      charts.push(<EquipmentsChart key={'chart-2'} />)
    }
    return charts
  }

  showDisplay () {
    if (!this.state.info.display) {
      return
    }
    return <Display />
  }

  componentWillMount () {
    this.refresh()
    setInterval(this.refresh, 180 * 1000)
  }

  refresh () {
    this.ajaxGet({
      url: '/api/info',
      success: function (data) {
        this.setState({
          info: data
        })
      }.bind(this)
    })
  }

  render () {
    return (
      <div className='container'>
        {super.render()}
        <div className='row'>
          <div className='col-sm-2'>Time</div>
          <div className='col-sm-6'>{this.state.info.current_time}</div>
        </div>
        <div className='row'>
          <div className='col-sm-2'>IP</div>
          <div className='col-sm-3'>{this.state.info.ip}</div>
        </div>
        <div className='row'>
          <div className='col-sm-2'>Up Since</div>
          <div className='col-sm-3'>{this.state.info.uptime}</div>
        </div>
        <div className='row'>
          <div className='col-sm-2'>Version</div>
          <div className='col-sm-3'>{this.state.info.version}</div>
        </div>
        <div className='row'>
          <div className='col-sm-2'>CPU Temperature</div>
          <div className='col-sm-3'>{this.state.info.cpu_temperature}</div>
        </div>
        <div className='row'>
          {this.showDisplay()}
        </div>
        <div className='row'>
          {this.showCharts()}
        </div>
      </div>
    )
  }
}
