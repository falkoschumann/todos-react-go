import { render, screen } from '@testing-library/react';
import { Router } from 'react-router-dom';
import { createMemoryHistory } from 'history';

import { Filter } from './utils';
import Footer from './Footer';

describe('Footer', () => {
  it('renders 1 todo with singular counter.', () => {
    const history = createMemoryHistory();

    render(
      <Router location={history.location} navigator={history}>
        <Footer activeCount={1} completedCount={1} filter={Filter.All} />
      </Router>
    );

    const countElement = screen.getByTestId('todo-count');
    expect(countElement.textContent).toEqual('1 item left');
  });

  it('renders multiple todos with plural counter.', () => {
    const history = createMemoryHistory();

    render(
      <Router location={history.location} navigator={history}>
        <Footer activeCount={2} completedCount={0} filter={Filter.All} />
      </Router>
    );

    const countElement = screen.getByTestId('todo-count');
    expect(countElement.textContent).toEqual('2 items left');
  });

  it('renders all as current page.', () => {
    const history = createMemoryHistory();

    render(
      <Router location={history.location} navigator={history}>
        <Footer activeCount={1} completedCount={1} filter={Filter.All} />
      </Router>
    );

    const linkElement: HTMLAnchorElement = screen.getByRole('link', { current: 'page' });
    expect(linkElement.text).toEqual('All');
  });

  it('renders active as current page .', () => {
    const history = createMemoryHistory();

    render(
      <Router location={history.location} navigator={history}>
        <Footer activeCount={1} completedCount={1} filter={Filter.Active} />
      </Router>
    );

    const linkElement: HTMLAnchorElement = screen.getByRole('link', { current: 'page' });
    expect(linkElement.text).toEqual('Active');
  });

  it('renders completed as current page.', () => {
    const history = createMemoryHistory();

    render(
      <Router location={history.location} navigator={history}>
        <Footer activeCount={1} completedCount={1} filter={Filter.Completed} />
      </Router>
    );

    const linkElement: HTMLAnchorElement = screen.getByRole('link', { current: 'page' });
    expect(linkElement.text).toEqual('Completed');
  });

  it('display clear completed button.', () => {
    const history = createMemoryHistory();

    render(
      <Router location={history.location} navigator={history}>
        <Footer activeCount={1} completedCount={1} filter={Filter.All} />
      </Router>
    );

    const buttonElement = screen.getByText(/clear completed/i);
    expect(buttonElement).toBeInTheDocument();
  });

  it('do not display clear completed button.', () => {
    const history = createMemoryHistory();

    render(
      <Router location={history.location} navigator={history}>
        <Footer activeCount={2} completedCount={0} filter={Filter.All} />
      </Router>
    );

    const buttonElement = screen.queryByText(/clear completed/i);
    expect(buttonElement).not.toBeInTheDocument();
  });
});
