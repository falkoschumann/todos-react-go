import { render, screen } from '@testing-library/react';
import { createMemoryHistory } from 'history';
import { Router } from 'react-router-dom';

import App from './App';

describe('App', () => {
  test('renders successfully.', () => {
    const history = createMemoryHistory();
    history.push('/');

    render(
      <Router location={history.location} navigator={history}>
        <App />
      </Router>
    );

    const linkElement = screen.getByText(/todos/i);
    expect(linkElement).toBeInTheDocument();
  });
});
