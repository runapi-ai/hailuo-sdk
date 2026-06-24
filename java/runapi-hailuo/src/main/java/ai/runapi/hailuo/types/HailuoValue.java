package ai.runapi.hailuo.types;

import ai.runapi.core.types.RunApiValue;

abstract class HailuoValue extends RunApiValue {
  HailuoValue(String value) {
    super(value);
  }
}
