import React from 'react';
import { connect } from 'react-redux';
import { PaymentMethod } from '../../api';
import NewPaymentMethod from './components/NewPaymentMethod';
import PaymentMethodComponent from './components/PaymentMethod';
import { removePaymentMethod, addPaymentMethod } from '../../store/User';
import './EditPaymentMethods.scss';

interface Props {
  stripe?: any;
  elements?: any;

  paymentMethods: PaymentMethod[];
  addPaymentMethod: (pm: PaymentMethod) => void;
  removePaymentMethod: (pm: PaymentMethod) => void;
}

interface State {
  saving: boolean;
  error?: string;
}

class EditPaymentMethods extends React.Component<Props, State> {
  readonly state: State = { saving: false };
  
  setError(error?: string) {
    this.setState({ error, saving: false })
  }

  render():JSX.Element {
    const { paymentMethods } = this.props;
    const { error, saving } = this.state;

    return(
      <div className='EditPaymentMethods'>
        { this.state.error ? <p className='error'>{error}</p> : null }

        { this.renderPaymentMethods() }

        <NewPaymentMethod
          saving={saving}
          key={paymentMethods.length}
          onError={this.setError.bind(this)}
          onSuccess={this.props.addPaymentMethod}
          onSubmit={() => this.setState({ saving: true })}  />
      </div>
    );
  }

  renderPaymentMethods(): JSX.Element {
    if(this.props.paymentMethods.length === 0) return null;
    
    return(
      <div className='existing'>
        <h3>Existing Payment Methods</h3>
        { this.props.paymentMethods.map((pm: PaymentMethod) => {
          return <PaymentMethodComponent
                    key={pm.id}
                    paymentMethod={pm}
                    onError={this.setError.bind(this)} 
                    onDelete={this.props.removePaymentMethod} />
        })}
      </div>
    )
  }
}

function mapStateToProps(state: any): any {
  return({
    paymentMethods: state.user.user.paymentMethods,
  });
}

function mapDispatchToProps(dispatch: Function): any {
  return({
    addPaymentMethod: (pm: PaymentMethod) => dispatch(addPaymentMethod(pm)),
    removePaymentMethod: (pm: PaymentMethod) => dispatch(removePaymentMethod(pm)),
  });
}

export default connect(mapStateToProps, mapDispatchToProps)(EditPaymentMethods);