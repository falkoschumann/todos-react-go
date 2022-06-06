async function getJson(endpoint: string, parameter?: any) {
  let url = endpoint;
  if (parameter) {
    endpoint += "?" + new URLSearchParams(parameter);
  }
  const response = await fetch(url);
  if (!response.ok) {
    const msg = await response.text();
    throw new Error(`${endpoint} - ${msg}`);
  }

  return response.json();
}

async function postJson(endpoint: string, payload: any) {
  const response = await fetch(endpoint, {
    method: "POST",
    body: JSON.stringify(payload),
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
  });
  if (!response.ok) {
    const msg = await response.text();
    throw new Error(`${endpoint} - ${msg}`);
  }

  return response.json();
}

const APIUtils = {
  getJson,
  postJson,
};

export default APIUtils;
