package {{packageName}}.{{moduleName}}.entity;
{% if hasJoinColumn  %}
import com.baomidou.mybatisplus.annotations.TableField;{% endif %}
import com.baomidou.mybatisplus.annotations.TableId;
import com.baomidou.mybatisplus.annotations.TableName;
import lombok.Data;
import java.io.Serializable;{% if hasBigDecimal  %}
import java.math.BigDecimal;{% endif %}{% if hasDate  %}import java.util.Date;{% endif %}{% if hasTime  %}import java.sql.Time;{% endif %}



/**
 * {{ table.TableComment }}
 *
 * @author {{ authorName }}
 * @email {{ emailAddress }}
 * @date {{ genTime }}
 */
@TableName("{{ table.TableName }}")
@Data
public class {{ className }}Entity implements Serializable {
	private static final long serialVersionUID = 1L;
    {% for column in listColumns %}{% if !column.IsJoinColumn%}
	/**
	 * {{ column.ColumnComment }}
	 */
    {% if column.ColumnKey == "PRI" %}
    @TableId
    {% endif %}
    private {{ column.JavaType }} {{ column.FieldName}};
    {% endif %}
    {% endfor %}
    {% for column in listColumns %}
    {% if column.IsJoinColumn %}
	/**
	 * {{ column.ColumnComment }}
	 */
    @TableField(exist = false)
    private {{ column.JavaType }} {{ column.FieldName}};
    {% endif %}
    {% endfor %}
}
