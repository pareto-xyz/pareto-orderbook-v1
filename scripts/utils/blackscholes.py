import math
import numpy as np
from scipy.stats import norm


def get_probability_factors(spot, strike, sigma, tau, rate):
  d1 = (math.log(spot/strike)+(rate+tau*sigma**2/2.0))/(sigma*math.sqrt(tau))
  d2 = d1-sigma*math.sqrt(tau)
  return d1, d2


def get_call_price(spot, strike, sigma, tau, rate):
  d1, d2 = get_probability_factors(spot, strike, sigma, tau, rate)
  price = spot*norm.cdf(d1)-strike*math.exp(-rate*tau)*norm.cdf(d2)
  return price


def get_put_price(spot, strike, sigma, tau, rate):
  d1, d2 = get_probability_factors(spot, strike, sigma, tau, rate)
  price = strike*math.exp(-rate*tau)*norm.cdf(-d2)-spot*norm.cdf(-d1)
  return price


def get_vega(spot, strike, sigma, tau, rate):
  d1, _ = get_probability_factors(spot, strike, sigma, tau, rate)
  vega = spot*math.sqrt(tau)*norm.pdf(d1)
  return vega


def brute_force_iv(spot, strike, rate, tau, market, is_call=True):
  candidates = np.arange(0.0001,4,0.001)
  diffs = np.zeros_like(candidates)

  for i in range(len(candidates)):
    candidate = candidates[i]
    if is_call:
      diff = get_call_price(spot, strike, candidate, tau, rate) - market
    else:
      diff = get_put_price(spot, strike, candidate, tau, rate) - market
    diffs[i] = abs(diff)
  
  best_i = np.argmin(diffs)
  return candidates[best_i], diffs[best_i]
