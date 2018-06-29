import React from 'react'
import { Area, Tooltip, YAxis, XAxis, AreaChart } from 'recharts'
import {fetchTCUsage} from '../redux/actions/tcs'
import {connect} from 'react-redux'

class chart extends React.Component {
  componentDidMount () {
    this.props.fetchTCUsage(this.props.sensor_id)
    var timer = window.setInterval(() => { this.props.fetchTCUsage(this.props.sensor_id) }, 10 * 1000)
    this.setState({timer: timer})
  }

  componentWillUnmount () {
    if (this.state && this.state.timer) {
      window.clearInterval(this.state.timer)
    }
  }

  render () {
    if (this.props.usage === undefined) {
      return (<div />)
    }
    if (this.props.config === undefined) {
      return (<div />)
    }
    var min = this.props.config.chart_min
    var max = this.props.config.chart_max
    return (
      <div className='container'>
        <span className='h6'>{this.props.config.name} - Temperature</span>
        <AreaChart width={this.props.width} height={this.props.height} data={this.props.usage.current}>
          <defs>
            <linearGradient id='gradient' x1='0' y1='0' x2='0' y2='1'>
              <stop offset='5%' stopColor='#00C851' stopOpacity={0.8} />
              <stop offset='95%' stopColor='#007E33' stopOpacity={0} />
            </linearGradient>
          </defs>
          <YAxis domain={[min, max]} dataKey='temperature' />
          <XAxis dataKey='time' />
          <Tooltip />
          <Area type='linear' dataKey='temperature' stroke='#007E33' isAnimationActive={false} fillOpacity={1} fill='url(#gradient)' />
        </AreaChart>
      </div>
    )
  }
}

const mapStateToProps = (state, ownProps) => {
  return {
    config: state.tcs.find((el) => { return el.id === ownProps.sensor_id }),
    usage: state.tc_usage[ownProps.sensor_id]
  }
}

const mapDispatchToProps = (dispatch) => {
  return {
    fetchTCUsage: (id) => dispatch(fetchTCUsage(id))
  }
}

const ReadingsChart = connect(mapStateToProps, mapDispatchToProps)(chart)
export default ReadingsChart
